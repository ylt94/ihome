module github.com/ylt94/ihome/services/registerAndLogin

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220520155318-0235a592b59a => ../services/getCaptcha

replace github.com/ylt94/ihome/services/sendSMS v0.0.0-20220520155318-0235a592b59a => ../services/sendSMS
