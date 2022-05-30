package handler

import (
	"context"
	"errors"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/ylt94/ihome/services/registerAndLogin/model"
	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
	"github.com/ylt94/ihome/services/registerAndLogin/utils"
)

type Login struct{}

func (e *Login) Login(ctx context.Context, req *login.LoginRequest, rsp *login.LoginResponse) error {
	log.Info("Received RegisterAndLogin.Call request")
	rsp.LoginStatus = login.LoginStatus_LoginFail
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
	rsp.LoginStatus = login.LoginStatus_LoginSuccess

	return nil
}

func (e *Login) Auth(ctx context.Context, req *login.AuthRequest, rsp *login.AuthResponse) error {
	token := req.Token
	phone := utils.DecryptionTokenByJWT(token)
	user := &model.User{}
	user.GetUserByPhone(phone)

	rsp.AuthStatus = login.AuthStatus_AuthFail
	if user.Id > 0 {
		rsp.AuthStatus = login.AuthStatus_AuthSuccess
		userInfo := &login.UserInfo{Id: int32(user.Id), Name: user.Name, Mobile: user.Mobile, AvatarUrl: user.AvatarUrl}
		rsp.User = userInfo
		return nil
	}

	return errors.New("身份验证失败!")
}