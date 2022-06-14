package handler

import (
	"context"
	"errors"

	log "github.com/micro/go-micro/v2/logger"

	auth "github.com/ylt94/ihome/services/user/proto/auth"
	"github.com/ylt94/ihome/services/user/model"
)

type Auth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth)CheckRealNameAuth (ctx context.Context, req *auth.CheckAuthRequest, rsp *auth.CheckAuthResponse) error {
	log.Info("Received auth.CheckRealNameAuth request")
	db := model.Db()
	if req.UserId <= 0 {
		return errors.New("缺少参数:UserId")
	}
	user := model.User{}

	fileds := []string{"id","name", "mobile", "real_name", "id_card", "avatar_url"}
	db.Table("user").Where("id = ?", req.UserId).Select(fileds).First(&user)
	if user.Id == 0 {
		return errors.New("用户未注册")
	}
	rsp.UserId = user.Id
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.RealName
	rsp.IdCard = user.IdCard
	rsp.AvatarUrl = user.AvatarUrl
	return nil
}

func (e *Auth) RealNameAuth(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Info("Received auth.RealNameAuth request")
	if req.UserId <= 0 {
		return errors.New("缺少参数:UserId")
	}
	if req.IdCard == "" {
		return errors.New("请输入身份证号")
	}
	if req.RealName == "" {
		return errors.New("请输入真实姓名")
	}

	updateData := map[string]interface{}{"id_card":req.IdCard,"real_name":req.RealName}
	db := model.Db()
	db.Model(&User{}).Where("id", req.UserId).Updates(updateData)
	rsp.Status = auth.AuthStatus_AuthSucc
	return nil
}