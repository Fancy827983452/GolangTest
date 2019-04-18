package main

import (
	"github.com/kataras/iris"
	"controller"
	"model"
)

func main() {
	app := iris.New()
	//设置静态文件请求路径，HTML页面请求时，src="/static/xxx"
	app.StaticWeb("/static", "./src/view/static")
	// Reload 方法设置为 true 表示开启开发者模式 将会每一次请求都重新加载 views 文件下的所有模板
	app.RegisterView(iris.HTML("./src/view", ".html").Reload(true))

	app.Post("/decode", func(ctx iris.Context) {
		var user model.User
		// 请求参数格式化  请求参数是json类型转化成 User类型
		// 比如 post 参数 {username:'xxxx'} 转成 User 类型
		ctx.ReadJSON(&user)//把 json 类型请求参数 转成结构体
		//ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	app.Get("/login", controller.Login)
	app.Post("/loginPost", controller.LoginPost)

	app.Get("/register", controller.Register)
	app.Post("/registerPost", controller.RegisterPost)

	app.Get("/logout", controller.Logout)

	//app.Party 定义路由组,把相同路由组的放在一个区块(第一个参数设置路由相同的前缀,第二个参数为中间件)
	user:=app.Party("user")
	{
		user.Get("/editInfo",controller.EditUserInfo)
		user.Post("/editInfoPost",controller.EditUserInfoPost)
		user.Get("/editPassword",controller.EditUserPwd)
		user.Post("/editPasswordPost",controller.EditUserPwdPost)
		user.Get("/illcase",controller.UserIllcase)
		user.Post("/illcase/lockRecord",controller.LockRecord)
		user.Post("/illcase/unlockRecord",controller.UnlockRecord)
		user.Get("/visitRecord",controller.VisitUserRecord)
	}

	doctor:=app.Party("doctor")
	{
		doctor.Get("/doctor",controller.Doctor)
	}

	// 为特定HTTP错误注册自定义处理程序方法
	// 当出现 StatusInternalServerError 500错误，将执行第二参数回调方法
	//app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
	//	// ctx.Values() 是一个很有用的东西，主要用来使 处理方法与中间件 通信
	//	errMessage := ctx.Values().GetString("error")//获取自定义错误提示信息
	//	if errMessage != "" {
	//		ctx.Writef("Internal server error: %s", errMessage)
	//		return
	//	}
	//	ctx.Writef("(Unexpected) internal server error")
	//})
	app.Run(iris.Addr(":8080"))
}
