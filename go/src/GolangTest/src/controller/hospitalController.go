package controller

import (
	"github.com/kataras/iris"
	"model"
	"encoding/json"
	"util"
	"encoding/hex"
	"algorithm"
)

func VerifyDoctor(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)

	hospital:=model.Hospital{}
	json.Unmarshal([]byte(currentHospital.(string)),&hospital)//interface -> 结构体
	records:=model.GetUnverifiedDoctors(hospital.HospitalId)
	Items := records.Items //解密
	for i:= range Items {
		aec_key:=Items[i].Aec_Key
		idnum, _ :=hex.DecodeString(Items[i].IdNum)
		Items[i].IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(aec_key)))
	}
	ctx.ViewData("doctors",util.ParseJson(records))
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