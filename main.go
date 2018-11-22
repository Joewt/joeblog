package main

import (
	_ "github.com/yinrenxin/joeblog/routers"
	_ "github.com/yinrenxin/joeblog/models"
	"github.com/astaxie/beego"
	"strings"

	"encoding/gob"
	"github.com/yinrenxin/joeblog/models"
)

func main() {
	initTemplate()
	initSession()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) (bool) {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
	beego.AddFuncMap("add", func(a, b int64)(int64){
		return a + b
	})
}

func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "joeblog"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./data/session"
}