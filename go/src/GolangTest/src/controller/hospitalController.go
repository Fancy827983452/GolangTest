package controller

import "github.com/kataras/iris"

func VerifyDoctor(ctx iris.Context)  {
	ctx.View("hospital/HospitalVerifyDoctor.html")
}

func ViewDoctors(ctx iris.Context)  {
	ctx.View("hospital/HospitalViewAllDoctor.html")
}

func HospitalDepartmentManagement(ctx iris.Context)  {
	ctx.View("hospital/HospitalDepartmentManagement.html")
}