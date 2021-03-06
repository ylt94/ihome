package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	sendSMS "github.com/ylt94/ihome/services/sendSMS/proto/sendSMS"
	"github.com/ylt94/ihome/services/sendSMS/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type SendSMS struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *SendSMS) Send(ctx context.Context, req *sendSMS.Request, rsp *sendSMS.Response) error {
	log.Info("Received SendSMS.Call request")
	codeLen := req.CodeLen
	if codeLen <= 0 {
		codeLen = 4
	}

	//生成验证码
	rand.Seed(time.Now().UnixNano())
	var code string
	var codeString strings.Builder
	for i := 0; i < int(codeLen); i++ {
		number := strconv.Itoa(rand.Intn(10))
		codeString.WriteString(number)
	}
	code = codeString.String()

	rsp.Phone = req.Phone
	codeInt, err := strconv.ParseInt(code, 10, 32)
	if err != nil {
		return  err
	}

	rsp.Code = int32(codeInt)

	//验证码存起来
	key := sendSMS.SMSType_name[int32(sendSMS.SMSType_Register)] + "_" + req.Phone

	redis := utils.Redis()
	_, err = redis.Do("setex", key, 5 * 60, rsp.Code)
	if err != nil {
		return err
	}

	return nil
}
