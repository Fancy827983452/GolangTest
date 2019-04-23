package controller

import (
	"github.com/kataras/iris"
	"model"
	"encoding/json"
	"util"
	"encoding/hex"
	"algorithm"
	"strings"
	"fmt"
)

func VerifyDoctor(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)

	if(currentHospital!=nil){
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
	}

	ctx.View("hospital/HospitalVerifyDoctor.html")
}

func PassDoctor(ctx iris.Context)  {//审核医生注册通过
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateDoctorStatus(Ids[i],1)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/hospital/verifyDoctor';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func FailDoctor(ctx iris.Context)  {//审核医生注册不通过
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateDoctorStatus(Ids[i],-1)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/hospital/verifyDoctor';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
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