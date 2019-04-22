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

	app.Get("/", controller.Login)
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
		doctor.Get("/",controller.DoctorLogin)	//医生登陆
		doctor.Get("/register",controller.DoctorRegister)	//注册在职医生
		doctor.Post("/registerPost",controller.DoctorRegisterPost)
		doctor.Get("/login",controller.DoctorLogin)	//医生登陆
		doctor.Post("/loginPost",controller.DoctorLoginPost)
		doctor.Get("/main",controller.DoctorMain)//医生主界面，包含查看所有当前挂号预约的信息
		doctor.Get("/editInfo",controller.DoctorEditInfo)//修改信息
		doctor.Post("/editInfoPost",controller.DoctorEditInfoPost)
		doctor.Get("/editPwd",controller.DoctorEditPwd)//修改密码
		doctor.Post("/editPwdPost",controller.DoctorEditPwdPost)
		doctor.Get("/visitHistory",controller.VisitHistory)//查看当前医生访问过的病人历史记录
		doctor.Get("/patientDetails",controller.PatientDetails)	//查看病人详细信息
		doctor.Get("/patientHistoryCase",controller.PatientHistoryCase)	//查看病人历史病历
		doctor.Get("/addCase",controller.AddCase)	//添加病例
		doctor.Get("/patientTreatmentHistory",controller.PatientTreatmentHistory)	//查看病人就诊记录
		doctor.Get("/departmentManagement",controller.DepartmentManagement)//科室管理员
		doctor.Get("/viewArrangement",controller.ViewDepartmentArrangement)
		doctor.Get("/setAppointmentNum",controller.SetAppointmentNum)

		doctor.Get("/logout", controller.Logout)
	}

	hospital:=app.Party("hospital")
	{
		hospital.Get("/",controller.HospitalLogin)
		hospital.Get("/login",controller.HospitalLogin)
		hospital.Post("/loginPost",controller.HospitalLoginPost)
		hospital.Get("/register",controller.HospitalRegister)
		hospital.Post("/registerPost",controller.HospitalRegisterPost)
		hospital.Get("/verifyDoctor",controller.VerifyDoctor)
		hospital.Get("/viewDoctors",controller.ViewDoctors)
		hospital.Get("/departmentManagement",controller.HospitalDepartmentManagement)

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
