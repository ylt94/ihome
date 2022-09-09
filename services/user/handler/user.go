package handler

import (
	"context"
	"errors"
	"github.com/ylt94/ihome/services/user/model"

	log "github.com/micro/go-micro/v2/logger"

	user "github.com/ylt94/ihome/services/user/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Info(ctx context.Context, req *user.InfoRequest, rsp *user.InfoResponse) error {
	log.Info("Received User.Info request")
	userInfo := new(model.User)
	db := model.Db()
	db.Table("user").Where("id", req.UserId).Select("id","name", "mobile","real_name","id_card","avatar_url").First(userInfo)
	if userInfo.Id == 0 {
		return errors.New("用户未注册")
	}
	rsp.UserId = userInfo.Id
	rsp.Name = userInfo.Name
	rsp.Mobile = userInfo.Mobile
	rsp.RealName = userInfo.RealName
	rsp.IdCard = userInfo.IdCard
	rsp.AvatarUrl = userInfo.AvatarUrl
	return nil
}

func (e *User) UploadAvatar(ctx context.Context, req *user.AvatarRequest, rsp *user.AvatarResponse) error {
	log.Info("Received User.UploadAvatar request")
	if req.AvatarUrl == "" {
		return errors.New("图片路径为空")
	}
	if req.UserId == 0 {
		return errors.New("缺少参数:UserId")
	}

	db := model.Db()
	db.Table("user").Where("id", req.UserId).Update("avatar_url", req.AvatarUrl)
	rsp.AvatarUrl = req.AvatarUrl
	rsp.Status = user.AvatarStatus_AvatarStatusSucc
	return nil
}

func (e *User) UpdateName(ctx context.Context, req *user.UpNameRequest, rsp *user.UpNameResponse) error {
	log.Info("Received User.UpdateName request")
	if req.Name == "" {
		return errors.New("请输入用户名称")
	}
	if req.UserId == 0 {
		return errors.New("缺少参数:UserId")
	}

	db := model.Db()
	db.Table("user").Where("id", req.UserId).Update("name", req.Name)
	rsp.Name = req.Name
	return nil
}


