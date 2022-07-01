module github.com/ylt94/ihome/web

go 1.13

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/garyburd/redigo v1.6.3
	github.com/gin-gonic/gin v1.7.7
	github.com/micro/go-micro/v2 v2.9.1
	github.com/ylt94/ihome/services/getCaptcha v0.0.0-20220526174807-23acb797b957
	github.com/ylt94/ihome/services/index v0.0.0-20220627153528-98ebfd08d09f
	github.com/ylt94/ihome/services/registerAndLogin v0.0.0-00010101000000-000000000000
	github.com/ylt94/ihome/services/sendSMS v0.0.0-20220526174807-23acb797b957
	github.com/ylt94/ihome/services/user v0.0.0-20220630083056-33f1c50fe656
	golang.org/x/text v0.3.6 // indirect
)

replace github.com/ylt94/ihome/services/getCaptcha => ../services/getCaptcha

replace github.com/ylt94/ihome/services/sendSMS => ../services/sendSMS

replace github.com/ylt94/ihome/services/registerAndLogin => ../services/registerAndLogin
