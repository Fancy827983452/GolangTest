package controller

import "github.com/kataras/iris"

func EditInfo(ctx iris.Context) {
	ctx.View("user/UserEditInfo.html")
}

func EditPwd(ctx iris.Context) {
	ctx.View("user/UserEditPassword.html")
}

func Illcase(ctx iris.Context) {
	ctx.View("user/Userillcase.html")
}

func VisitRecord(ctx iris.Context) {
	ctx.View("user/UserVisitRecord.html")
}
