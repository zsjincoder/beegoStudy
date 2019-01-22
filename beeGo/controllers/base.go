package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code int		`json:"code"`
	Msg  string		`json:"msg"`
	Data interface{}	`json:"data"`
}

type ErrResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  interface{} `json:"errmsg"`
}
