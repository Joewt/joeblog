package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Details",
            Router: `/detail/:k`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Editor",
            Router: `/editor`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Method: "IndexCreate",
            Router: `/create`,
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
            Method: "IndexReg",
            Router: `/reg`,
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

    beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/joeblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Reg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
