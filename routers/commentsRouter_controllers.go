package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexComment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexUser",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
