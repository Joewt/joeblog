package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yinrenxin/joeblog/models"
	"errors"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User models.User
	IsLogin bool
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)

	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
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