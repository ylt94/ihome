module github.com/ylt94/ihome/services/registerAndLogin

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/services/getCaptcha v1.0.0 => ../services/getCaptcha

replace github.com/services/sendSMS v1.0.0 => ../services/sendSMS
