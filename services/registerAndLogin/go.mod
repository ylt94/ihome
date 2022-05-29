module github.com/ylt94/ihome/services/registerAndLogin

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220526174807-23acb797b957
	github.com/ylt94/ihome/services/sendSMS v0.0.0-20220526174807-23acb797b957

	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)

replace github.com/ylt94/ihome/services/sendSMS v0.0.0-20220526174807-23acb797b957 => ../sendSMS

replace github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220526174807-23acb797b957 => ../getCaptcha
