package main

import (
	"github.com/kataras/iris"
	"controller"
	"model"
	"util"
	"algorithm"
	"strconv"
)

func main() {
	app := iris.New()
	//app.StaticWeb("./src/view/static", "/")
	// Reload 方法设置为 true 表示开启开发者模式 将会每一次请求都重新加载 views 文件下的所有模板
	app.RegisterView(iris.HTML("./src/view", ".html").Reload(true))

	//获取session管理器
	sessionMgr:=util.NewSessionManager()

	// 为特定HTTP错误注册自定义处理程序方法
	// 当出现 StatusInternalServerError 500错误，将执行第二参数回调方法
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// ctx.Values() 是一个很有用的东西，主要用来使 处理方法与中间件 通信
		errMessage := ctx.Values().GetString("error")//获取自定义错误提示信息
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Post("/decode", func(ctx iris.Context) {
		var user model.User
		// 请求参数格式化  请求参数是json类型转化成 User类型
		// 比如 post 参数 {username:'xxxx'} 转成 User 类型
		ctx.ReadJSON(&user)//把 json 类型请求参数 转成结构体
		//ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	app.Get("/login", controller.Login)
	app.Post("/loginPost", func (ctx iris.Context) {
			var u model.User
			var msg model.Uploador
			u.Name = ctx.FormValue("username")
			u.Password = ctx.FormValue("password")
			u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
			count,err:=model.CheckLogin(u)
			util.CheckErr(err)
			if count==0{
				msg.Success = false
				msg.Message="用户名或密码错误！"
				ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
			} else {
				user,err := model.Login(u);//获取公钥
				util.CheckErr(err)
				msg.Success = true
				msg.Message="登陆成功！"
				//获取session管理器
				session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
				session.Set("currentUser",util.ParseJson(user))
				ctx.HTML("<script>alert('"+msg.Message+"');" +
				"window.location.href='user/editInfo/"+user.Name+"';</script>")//URL传参
			}
		})

	app.Get("/register", controller.Register)
	app.Post("/registerPost", controller.RegisterPost)

	app.Get("/logout", func(ctx iris.Context) {
		sessionMgr.Destroy(ctx.ResponseWriter(),ctx.Request())
		ctx.HTML("<script>alert('登出成功！');window.location.href='/login';</script>")
	})

	//app.Party 定义路由组,把相同路由组的放在一个区块(第一个参数设置路由相同的前缀,第二个参数为中间件)
	user:=app.Party("user")
	{
		user.Get("/editInfo/{username}",func (ctx iris.Context) {
			username:=ctx.Params().Get("username")//获取路由参数
			ctx.ViewData("username",username)//向数据模板传值
			session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
			currentUser:=session.Get("currentUser")
			//fmt.Println("currentUser",currentUser)
			ctx.ViewData("currentUser",currentUser)
			ctx.View("user/UserEditInfo.html")
		})
		user.Post("/editInfoPost",func (ctx iris.Context) {
			var u model.User
			var msg model.Uploador
			u.PublicKey = ctx.FormValue("publicKey")
			u.Name = ctx.FormValue("username")
			u.Gender, _ = strconv.Atoi(ctx.FormValue("sex"))
			u.PhoneNum = ctx.FormValue("tel")
			u.BirthDate = ctx.FormValue("birthdate")
			u.Location = ctx.FormValue("location")
			//ctx.JSON(u)
			result, _ := model.UpdateInfo(u); //插入数据库，返回操作结果（true或false）
			if result > 0 {
				msg.Success = true
				msg.Message = "个人信息更新成功！"
				user, err := model.GetInfoByPublicKey(u)
				util.CheckErr(err)
				session := sessionMgr.BeginSession(ctx.ResponseWriter(), ctx.Request())
				session.Set("currentUser", util.ParseJson(user)) //更新session
				ctx.HTML("<script>alert('" + msg.Message + "');" +
					"window.location.href='editInfo/" + u.Name + "';</script>")
			} else {
				msg.Message="个人信息更新失败！"
				ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
			}
		})
		user.Get("/editPassword/{username}",controller.EditPwd)
		user.Get("/illcase/{username}",controller.Illcase)
		user.Get("/visitRecord/{username}",controller.VisitRecord)
	}

	app.Run(iris.Addr(":8080"))
}
