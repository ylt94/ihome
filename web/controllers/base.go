package controllers

import (
	"github.com/ylt94/ihome/web/utils"
)


func getReturn(data interface{}, inputCode ...string) map[string]interface{} {
	res := make(map[string]interface{}, 3)
	code :=  utils.RECODE_OK
	msg := ""

	if len(inputCode) >= 1 {
		code = inputCode[0]
	}
	if len(inputCode) >= 2 {
		msg = inputCode[1]
	}

	res["errno"] = code
	res["errmsg"] = msg
	res["data"] = data

	return res
}