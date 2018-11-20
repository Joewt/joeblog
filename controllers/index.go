package controllers

type IndexController struct {
	BaseController
}


// @router / [get]
func (this *IndexController) Index() {
	//this.Abort("404")
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