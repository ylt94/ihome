package handler

import (
	"context"
	"errors"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/ylt94/ihome/services/registerAndLogin/model"
	"github.com/ylt94/ihome/services/registerAndLogin/utils"

	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
)

type Login struct{}

func (e *Login) Login(ctx context.Context, req *login.Request, rsp *login.Response) error {
	log.Info("Received RegisterAndLogin.Call request")
	rsp.Status = login.LoginStatus_Fail
	user := &model.User{}
	err := user.GetUserByPhone(req.Phone)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("用户不存在")
	}

	md5Pwd := utils.EncryptionByMD5(req.Pwd)
	if md5Pwd != user.PasswordHash {
		return errors.New("密码错误!")
	}
	rsp.Token = utils.CreateTokenByJWT(req.Phone)
	rsp.Status = login.LoginStatus_Success

	return nil
}

func (e *Login) Auth(ctx context.Context, req *login.AuthRequest, rsp *login.AuthResponse) error {
	token := req.Token
	phone := utils.DecryptionTokenByJWT(token)
	user := &model.User{}
	user.GetUserByPhone(phone)

	rsp.Status = login.AuthStatus_AuthFail
	if user.Id > 0 {
		rsp.Status = login.AuthStatus_AuthSuccess
		rsp.User.Id = int32(user.Id)
		rsp.User.Name = user.Name
		rsp.User.Mobile = user.Mobile
		rsp.User.AvatarUrl = user.AvatarUrl
		return nil
	}

	return errors.New("身份验证失败!")
}