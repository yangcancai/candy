package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)
// 扩展prepare函数
type ExtPreparerInterface interface {
	ExtPrepare()
}
type ApiController struct {
	beego.Controller
	Request map[string]interface{}
}
type ApiResp struct {
	Code int `json:"code"`
	Err_msg string `json:"err_msg"`
}
// 接口入口校验函数
func (this *ApiController) Prepare(){
	// 做一些接口校验之类的工作
	if this.Auth() == false {
		return
	}
	// 如果子类中有实现ExtPreparerInterface接口则执行该接口
	if app, ok := this.AppController.(ExtPreparerInterface); ok {
		app.ExtPrepare()
	}
}
func (this *ApiController) Auth() bool{
	// json格式
	if this.Ctx.Input.Header("Content-Type") == "application/json"{
		err := json.Unmarshal([]byte(this.Ctx.Input.RequestBody), &this.Request)
		if err != nil {
			this.Data["json"] = &ApiResp{1, "invalid json format"}
			this.ServeJSON()
			this.StopRun()
			return false
		}
	}else if this.Ctx.Input.Header("Content-Type") == "application/x-www-form-urlencoded"{
	}else{
		this.Data["json"] = &ApiResp{1,"error Content-Type"}
		this.ServeJSON()
		this.StopRun()
		return false
	}
	if this.Ctx.Input.Header("appid") == "cam" &&
		this.Ctx.Input.Header("appkey") == "123456"{
		return true
	}else{
		this.Data["json"] = &ApiResp{1,"auth error"}
		this.ServeJSON()
		this.StopRun()
		return false
	}
}
