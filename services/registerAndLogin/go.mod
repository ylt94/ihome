module github.com/ylt94/ihome/services/registerAndLogin

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220520162116-f2ea8d2a96ea
	github.com/ylt94/ihome/services/sendSMS v0.0.0-20220520162116-f2ea8d2a96ea
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220520155318-0235a592b59a => ../services/getCaptcha

replace github.com/ylt94/ihome/services/sendSMS v0.0.0-20220520155318-0235a592b59a => ../services/sendSMS

replace github.com/ylt94/ihome/services/registerAndLogin => ../services/registerAndLogin
