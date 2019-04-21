package controller

import "github.com/kataras/iris"

func HospitalLogin(ctx iris.Context)  {
	ctx.View("hospital/HospitalLogin.html")
}

func HospitalManagement(ctx iris.Context)  {
	ctx.View("hospital/HospitalManagement.html")
}

func VerifyDoctor(ctx iris.Context)  {
	ctx.View("hospital/HospitalVerifyDoctor.html")
}