package main

import (
	"github.com/kataras/iris"
	"controller"
)

func main() {
	app := iris.New()
	//app.StaticWeb("./static", "/")
	//app.RegisterView(iris.HTML("./view", ".html"))
	tmpl := iris.HTML("C:/Users/LEE/go/src/GolangTest/src/view", ".html")//此处路径好像必须是绝对路径？
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	app.Get("/login", controller.Login)
	app.Post("/loginPost", controller.LoginPost)
	app.Get("/register", controller.Register)
	app.Post("/registerPost", controller.RegisterPost)

	user:=app.Party("user")
	{
		user.Get("/editInfo",controller.EditInfo)
		user.Get("/editPassword",controller.EditPwd)
		user.Get("/illcase",controller.Illcase)
		user.Get("/visitRecord",controller.VisitRecord)
	}

	app.Run(iris.Addr(":8080"))
}
