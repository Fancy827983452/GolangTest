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

func WithdrawDoctor(ctx iris.Context)  {//注销医生
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
			model.UpdateDoctorRole(Ids[i],0)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/hospital/viewDoctors';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func ViewDoctors(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	if(currentHospital!=nil){
		hospital:=model.Hospital{}
		json.Unmarshal([]byte(currentHospital.(string)),&hospital)//interface -> 结构体
		records:=model.GetValidDoctors(hospital.HospitalId)
		Items := records.Items //解密
		for i:= range Items {
			aec_key:=Items[i].Aec_Key
			idnum, _ :=hex.DecodeString(Items[i].IdNum)
			Items[i].IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(aec_key)))
		}
		ctx.ViewData("doctors",util.ParseJson(records))
	}
	ctx.View("hospital/HospitalViewAllDoctor.html")
}

func SearchDoctor(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	param:=ctx.Params().Get("param")
	name:=ctx.FormValue("department")
	//fmt.Println("param=",param)
	//fmt.Println("name=",name)

	if(currentHospital!=nil){
		hospital:=model.Hospital{}
		json.Unmarshal([]byte(currentHospital.(string)),&hospital)//interface -> 结构体
		records:=model.GetSelectedDoctors(hospital.HospitalId,param,name)
		Items := records.Items //解密
		for i:= range Items {
			aec_key:=Items[i].Aec_Key
			idnum, _ :=hex.DecodeString(Items[i].IdNum)
			Items[i].IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(aec_key)))
		}
		ctx.ViewData("doctors",util.ParseJson(records))
	}
	ctx.View("hospital/HospitalViewAllDoctor.html")
}

func HospitalDepartmentManagement(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.View("hospital/HospitalDepartmentManagement.html")
}

func SetDoctorAdmin(ctx iris.Context)  {//设置医生为科室管理员
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateDoctorRole(Ids[i],1)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/hospital/viewDoctors';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func CancelDoctorAdmin(ctx iris.Context)  {//取消医生的科室管理员身份
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateDoctorRole(Ids[i],0)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/hospital/viewDoctors';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}