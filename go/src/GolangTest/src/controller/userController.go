package controller

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/sessions"
	//"util"
	//"fmt"
)

func EditPwd(ctx iris.Context) {
	username:=ctx.Params().Get("username")//获取路由参数
	ctx.ViewData("username",username)//向数据模板传值
	ctx.View("user/UserEditPassword.html")
}

func Illcase(ctx iris.Context) {
	username:=ctx.Params().Get("username")//获取路由参数
	ctx.ViewData("username",username)//向数据模板传值
	ctx.View("user/Userillcase.html")
}

func VisitRecord(ctx iris.Context) {
	username:=ctx.Params().Get("username")//获取路由参数
	ctx.ViewData("username",username)//向数据模板传值
	ctx.View("user/UserVisitRecord.html")
}
