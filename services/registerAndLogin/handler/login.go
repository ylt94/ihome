package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	user, err := user.GetUserByPhone(req.Phone)
	if err != nil {
		return err
	}

	md5Pwd := utils.EncryptionByMD5(req.Pwd)
	if md5Pwd != user.PasswordHash {
		return errors.New("密码错误!")
	}
	rsp.Token = utils.CreateTokenByJWT(req.Phone)
	rsp.Status = login.LoginStatus_Success

	return nil
}