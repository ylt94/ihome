package controllers

import (
	"encoding/json"
	"github.com/ylt94/ihome/web/utils"
)

type microResult struct{
	Id     string `json:"id"`
	Code   int `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func getReturn(data interface{}, inputCode ...string) map[string]interface{} {
	res := make(map[string]interface{}, 3)
	code :=  utils.RECODE_OK
	errmsg := ""

	if len(inputCode) >= 1 {
		code = inputCode[0]
	}
	if len(inputCode) >= 2 {
		microRes := &microResult{}
		err := json.Unmarshal([]byte(inputCode[1]), microRes)
		if err != nil {
			panic(err.Error())
		}
		errmsg = microRes.Detail
	}

	res["errno"] = code
	res["errmsg"] = errmsg
	res["data"] = data

	return res
}