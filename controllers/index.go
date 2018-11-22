package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/joeblog/syserror"
	"github.com/yinrenxin/joeblog/models"
)

type IndexController struct {
	BaseController
}


// @router / [get]
func (this *IndexController) Index() {
	page, err := this.GetInt64("page",1)
	title := this.GetString("title")
	var limit int64
	limit = 2
	if err != nil || page <= 0 {
		page = 1
	}

	art, num,_:= models.QueryArtByPage(title,page, limit)
	count, _ := models.QueryArtCount(title)
	totalPage := count/limit
	if count%limit != 0 {
		totalPage = totalPage + 1
	}
	logs.Info(num)
	this.Data["art"] = art
	this.Data["totalpage"] = totalPage
	this.Data["page"] = page
	this.Data["title"] = title


	this.TplName = "index.html"
}

// @router /about [get]
func (this *IndexController) IndexAbout() {
	this.TplName = "about.html"
}


// @router /message [get]
func (this *IndexController) IndexMessage() {
	this.TplName = "message.html"
}

// @router /comment [get]
func (this *IndexController) IndexComment() {
	this.TplName = "comment.html"
}

// @router /user [get]
func (this *IndexController) IndexUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) IndexReg() {
	this.TplName = "reg.html"
}

// @router /create [get]
func (this *IndexController) IndexCreate() {
	if this.IsAdmin == false {
		this.Abort401(syserror.New("权限不足",nil))
	}
	this.Data["key"] = this.UUID()
	this.TplName = "editor.html"
}