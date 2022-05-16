package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	sendSMS "ihome/services/sendSMS/proto/sendSMS"
	"ihome/services/sendSMS/utils"
	"math/rand"
	"strconv"
	"time"
)

type SendSMS struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *SendSMS) Send(ctx context.Context, req *sendSMS.Request, rsp *sendSMS.Response) error {
	log.Info("Received SendSMS.Call request")
	codeLne := req.CodeLen
	if codeLne <= 0 {
		codeLne = 4
	}

	//生成验证码
	rand.Seed(time.Now().UnixNano())
	var code string
	for i := 0; i < int(codeLne); i++ {
		code += string(rand.Intn(10))
	}

	rsp.Phone = req.Phone
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		return  err
	}
	rsp.Code = int32(codeInt)

	//验证码存起来
	key := sendSMS.SMSType_name[int32(sendSMS.SMSType_Register)] + "_" + req.Phone
	redis := utils.Redis()
	_, err = redis.Do("setex", key, time.Second * 5 * 60, rsp.Code)
	if err != nil {
		return err
	}

	return nil
}
