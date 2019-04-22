package controller

import "github.com/kataras/iris"

func VerifyDoctor(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.View("hospital/HospitalVerifyDoctor.html")
}

func ViewDoctors(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.View("hospital/HospitalViewAllDoctor.html")
}

func HospitalDepartmentManagement(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.View("hospital/HospitalDepartmentManagement.html")
}