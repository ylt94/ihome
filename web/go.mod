module ihome/web

go 1.13

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/garyburd/redigo v1.6.3
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/services/getCaptcha v1.0.0
	github.com/services/sendSMS v1.0.0
	golang.org/x/text v0.3.6 // indirect
)

replace github.com/services/getCaptcha v1.0.0 => ../services/getCaptcha

replace github.com/services/sendSMS v1.0.0 => ../services/sendSMS
