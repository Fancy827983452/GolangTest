package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"util"
)

func EditInfo(ctx iris.Context) {
	username:=ctx.Params().Get("username")//获取路由参数
	ctx.ViewData("username",username)//向数据模板传值

	sess := sessions.New(sessions.Config{Cookie: "mysession_cookie_name"})
	session:=sess.Start(ctx)
	currentUser:=util.ParseJson(session.Get("CurrentUser"))//json格式的数据
	ctx.ViewData("currentUser",currentUser)

	//把获得的动态数据username 绑定在 ./src/views/user/UserEditInfo.html模板 语法 {{ .username }}
	ctx.View("user/UserEditInfo.html")
}

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
