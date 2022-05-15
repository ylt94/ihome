package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	log "github.com/micro/go-micro/v2/logger"
	"ihome/services/getCaptcha/utils"
	"image/color"

	getCaptcha "ihome/services/getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	log.Info("Received GetCaptcha.Call request")
	cap := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	cap.SetFont("./config/comic.ttf")
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img, number := cap.Create(4,captcha.NUM)
	var imJson []byte
	imJson, err := json.Marshal(img)
	if err != nil {
		return err
	}
	rsp.Image = imJson
	rsp.Number = number

	//缓存验证码
	redis, err := utils.Redis()
	if err != nil {
		return err
	}
	defer redis.Close()

	_, err = redis.Do("setex", req.Uuid, 5*60, number)
	if err != nil {
		return err
	}

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetCaptcha) Stream(ctx context.Context, req *getCaptcha.StreamingRequest, stream getCaptcha.GetCaptcha_StreamStream) error {
	log.Infof("Received GetCaptcha.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&getCaptcha.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetCaptcha) PingPong(ctx context.Context, stream getCaptcha.GetCaptcha_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&getCaptcha.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
