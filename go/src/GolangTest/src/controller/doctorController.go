package controller

import "github.com/kataras/iris"

func Doctor(ctx iris.Context){
	ctx.View("doctor/Doctor.html")
}
