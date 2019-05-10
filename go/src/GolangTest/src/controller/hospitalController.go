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
	"strconv"
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
	ctx.ResponseWriter().Header().Set("content-type", "text/html")

	if(currentHospital!=nil) {
		hospital := model.Hospital{}
		json.Unmarshal([]byte(currentHospital.(string)), &hospital) //interface -> 结构体
		records := model.GetHospitalDepts(strconv.Itoa(hospital.HospitalId))
		ctx.ViewData("department", util.ParseJson(records))
	}
	ctx.View("hospital/HospitalDepartmentManagement.html")
}

func DepartmentAddPost(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	hospital := model.Hospital{}
	json.Unmarshal([]byte(currentHospital.(string)), &hospital) //interface -> 结构体

	var d model.Department
	var msg model.Uploador
	d.HospitalId = hospital.HospitalId
	d.Name = ctx.FormValue("DepartmentName")
	d.Info = ctx.FormValue("detail")
	count, _ :=model.CheckDepartment(d)
	if(count>0) {//判断科室是否已经注册过
		msg.Success = false
		msg.Message="该科室已注册过！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}else {
		result, _ := model.DepartmentAdd(d);//插入数据库，返回操作结果（true或false）
		if(result==true){
			msg.Success = result
			msg.Message="添加成功！"
			ctx.HTML("<script>alert('"+msg.Message+"');window.location.href='/hospital/departmentManagement';</script>")
		}
	}
}

func DepartmentEditInfoPost(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentHospital:=session.Get("currentHospital")
	ctx.ViewData("currentHospital",currentHospital)
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	hospital := model.Hospital{}
	json.Unmarshal([]byte(currentHospital.(string)), &hospital) //interface -> 结构体

	var d model.Department
	var msg model.Uploador
	d.Name = ctx.FormValue("DepartmentName2")
	d.Info = ctx.FormValue("detail2")
	d.DeptId, _ = strconv.Atoi(ctx.FormValue("DepartmentId"))

	result, _ := model.UpdateDepartmentInfo(d); //插入数据库，返回操作结果（true或false）
	if result > 0 {
		msg.Success = true
		msg.Message = "科室信息更新成功！"
		ctx.HTML("<script>alert('" + msg.Message + "');" +
			"window.location.href='departmentManagement';</script>")
	} else {
		msg.Message="没有做任何更改！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}
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