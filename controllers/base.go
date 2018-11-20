package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yinrenxin/joeblog/models"
	"errors"
	"github.com/yinrenxin/joeblog/syserror"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"
type MAP_H = map[string]interface{}

type BaseController struct {
	beego.Controller
	User models.User
	IsLogin bool
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	//logs.Info("是否登录",u.Name,ok,this.Ctx.Request.RequestURI)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["islogin"] = this.IsLogin
}


func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}


func (this *BaseController) GetMushString(key, msg string) string {
	k := this.GetString(key)
	if len(k) == 0 {
		this.Abort500(errors.New(msg))
	}
	return k
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

func (this *BaseController) JsonOK(msg, action string) {
	this.Data["json"] = MAP_H{
		"code": 0,
		"msg": msg,
		"action": "/",
	}
	this.ServeJSON()
}

func (this *BaseController) JsonOKH(msg string, data MAP_H) {
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}