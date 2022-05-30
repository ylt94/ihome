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

func GetReturn(data interface{}, inputCode ...string) map[string]interface{} {
	res := make(map[string]interface{}, 3)
	code :=  utils.RECODE_OK
	errmsg := ""

	if len(inputCode) >= 1 {
		code = inputCode[0]
	}
	if len(inputCode) >= 2 {
		errmsg = inputCode[1]
	}

	res["errno"] = code
	res["errmsg"] = errmsg
	res["data"] = data

	return res
}

func GetServiceError(msg string) *microResult {
	microRes := &microResult{}
	err := json.Unmarshal([]byte(msg), microRes)
	if err != nil {
		panic(err.Error())
	}
	return microRes
}