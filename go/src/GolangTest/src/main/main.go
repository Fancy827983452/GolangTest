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
		user.Get("/appointment",controller.UserAppointment)
		user.Get("/appointment/{param}",controller.UserAppointment)
		user.Post("/appointment",controller.UserAppointmentSearchPost)
		user.Post("/addAppointment",controller.UserAddAppointment)
	}

	doctor:=app.Party("doctor")
	{
		doctor.Get("/",controller.DoctorLogin)	//医生登陆
		doctor.Get("/register",controller.DoctorRegister)	//注册在职医生
		doctor.Post("/registerPost",controller.DoctorRegisterPost)
		doctor.Get("/login",controller.DoctorLogin)	//医生登陆
		doctor.Post("/loginPost",controller.DoctorLoginPost)
		doctor.Get("/main",controller.DoctorMain)//医生主界面，包含查看所有当前挂号预约的信息
		doctor.Post("/main",controller.DoctorMainPost)
		doctor.Get("/editInfo",controller.DoctorEditInfo)//修改信息
		doctor.Post("/editInfoPost",controller.DoctorEditInfoPost)
		doctor.Get("/editPwd",controller.DoctorEditPwd)//修改密码
		doctor.Post("/editPwdPost",controller.DoctorEditPwdPost)
		doctor.Get("/visitHistory",controller.VisitHistory)//查看当前医生访问过的病人历史记录
		//doctor.Get("/patientDetails",controller.PatientDetails)	//查看病人详细信息
		doctor.Get("/addCase/{appointmentId}",controller.AddCase)	//添加病例
		doctor.Post("/addCase",controller.AddCasePost)	//添加病例
		doctor.Get("/treatmentHistory",controller.TreatmentHistory)	//查看历史就诊记录
		doctor.Post("/treatmentHistory",controller.TreatmentHistorySearch)
		doctor.Get("/departmentManagement",controller.DepartmentManagement)//科室管理员
		doctor.Get("/viewArrangement",controller.ViewDepartmentArrangement)
		doctor.Get("/setAppointmentNum",controller.SetAppointmentNum)

		doctor.Get("/logout/{param}", controller.Logout)
	}

	hospital:=app.Party("hospital")
	{
		hospital.Get("/",controller.HospitalLogin)
		hospital.Get("/login",controller.HospitalLogin)
		hospital.Post("/loginPost",controller.HospitalLoginPost)
		hospital.Get("/register",controller.HospitalRegister)
		hospital.Post("/registerPost",controller.HospitalRegisterPost)
		hospital.Get("/verifyDoctor",controller.VerifyDoctor)//审核医生注册申请
		hospital.Post("/verifyDoctor/pass",controller.PassDoctor)
		hospital.Post("/verifyDoctor/fail",controller.FailDoctor)
		hospital.Post("/verifyDoctor/withdraw",controller.WithdrawDoctor)
		hospital.Get("/viewDoctors",controller.ViewDoctors)//查看所有医生
		hospital.Post("/setAdmin",controller.SetDoctorAdmin)
		hospital.Post("/cancelAdmin",controller.CancelDoctorAdmin)
		hospital.Post("/searchDoctor/{param}",controller.SearchDoctor)
		hospital.Get("/departmentManagement",controller.HospitalDepartmentManagement)//科室管理
		hospital.Post("/departmentAddPost", controller.DepartmentAddPost)
		hospital.Post("/departmentEditInfoPost", controller.DepartmentEditInfoPost)
		hospital.Get("/logout", controller.Logout)
	}

	admin:=app.Party("admin")
	{
		admin.Get("/",controller.AdminLogin)
		admin.Get("/login",controller.AdminLogin)
		admin.Post("/loginPost",controller.AdminLoginPost)
		admin.Get("/verifyHospitals",controller.VerifyHospitals)
		admin.Get("/jsVerifyHospitals/{status}",controller.JSVerifyHospitals)
		admin.Post("/verifyHospitals/pass",controller.PassHospital)
		admin.Post("/verifyHospitals/fail",controller.FailHospital)
		admin.Get("/logout", controller.Logout)
	}

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
	app.Run(iris.Addr(":8080"))
}
