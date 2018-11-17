package controllers

import (
	"github.com/yinrenxin/joeblog/models"
	"github.com/yinrenxin/joeblog/syserror"
	"fmt"
)

type UserController struct {
	BaseController
}


// @router /login [post]
func (this *UserController) Login() {
	email := this.GetMushString("email", "邮箱不能为空")
	pwd   := this.GetMushString("password", "密码不能为空")
	id, err := models.QueryUserByEmailAndPwd(email, pwd)

	if err != nil {
		this.Abort500(syserror.New("登录失败",err))
	}

	this.SetSession(SESSION_USER_KEY,id)

	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"action": "/",
	}

	fmt.Println("主键id： ", id)
	this.ServeJSON()

}