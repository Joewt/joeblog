package controllers

import (
	"bytes"
	"github.com/yinrenxin/joeblog/common"
	"github.com/yinrenxin/joeblog/models"
	"github.com/yinrenxin/joeblog/syserror"
	"github.com/PuerkitoBio/goquery"
	_"github.com/astaxie/beego/logs"
)


type ArticleController struct {
	BaseController
}
// @router /create/:key [post]
func (this *ArticleController) Create() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMushString("title", "标题不能为空")
	content := this.GetMushString("content", "内容不能为空")

	_, ok := models.QueryByKeyArt(key)
	if ok == true {
		this.Abort500(syserror.New("已经有该文章了，请勿重复添加",nil))
	}
	summary, err := getSummary(content)
	if err != nil {
		summary = ""
	}
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	err = models.InsertArticle(key, title, content,u.Id, summary)
	if err != nil {
		this.Abort500(syserror.New("系统错误",err))
	}
	this.JsonOK("添加文章成功","/")
}

func getSummary(h string)(string, error){
	var bufBytes bytes.Buffer
	if _, err := bufBytes.Write([]byte(h));err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(&bufBytes)
  	if err != nil {
    	return "", err
	  }
	  
	new_html := doc.Find("body").Text()
	if len(new_html) > 1000 {
		new_html = new_html[:1000]
	}
	return new_html, nil
} 

// @router /detail/:k [get]
func (this *ArticleController) Details() {
	k := this.Ctx.Input.Param(":k")
	id := common.StrToUint(k)
	art, _ := models.GetArtById(id)
	user, _ := models.QueryUserById(art.User_id)
	//u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.Data["art"] = art
	this.Data["user"] = user
	this.TplName = "details.html"
}

// @router /del [post]
func (this *ArticleController) Del() {
	k := this.GetMushString("key","参数错误")
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	if !ok {
		this.Abort500(syserror.New("未登陆",nil))
	}
	art_id := common.StrToUint(k)
	art, err := models.GetArtById(art_id)
	if err != nil {
		this.Abort500(syserror.New("系统错误",nil))
	}
	if art.User_id != u.Id {
		this.Abort500(syserror.New("没有权限",nil))
	}
	err = models.DelArtById(art_id)

	if err != nil {
		this.Abort500(syserror.NoArt())
	}
	this.JsonOK("删除文章成功","/")
}

// @router /editor/:k [get]
func (this *ArticleController) Editor() {
	k := this.Ctx.Input.Param(":k")
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	if !ok{
		this.Abort500(syserror.New("未登陆",nil))
	}
	id := common.StrToUint(k)
	art, _ := models.GetArtById(id)
	if u.Id != art.User_id {
		this.Abort500(syserror.New("没有该文章的编辑权限",nil))
	}
	this.Data["art"] = art
	this.TplName = "art_editor.html"
}