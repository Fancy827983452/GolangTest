package controller

import (
	"github.com/kataras/iris"
	"model"
	"encoding/json"
	"util"
	"strings"
	"fmt"
	"strconv"
)

//显示全部医院信息
func VerifyHospitals(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentAdmin:=session.Get("currentAdmin")
	ctx.ViewData("currentAdmin",currentAdmin)
	if(currentAdmin!=nil){
		admin:=model.Admin{}
		json.Unmarshal([]byte(currentAdmin.(string)),&admin)//interface -> 结构体
		records:=model.GetAllHospitals()
		ctx.ViewData("hospitals",util.ParseJson(records))
	}
	ctx.View("admin/VerifyHospital.html")
}

//显示选择状态的医院信息
func JSVerifyHospitals(ctx iris.Context)  {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentAdmin:=session.Get("currentAdmin")
	ctx.ViewData("currentAdmin",currentAdmin)

	if(currentAdmin!=nil){
		admin:=model.Admin{}
		json.Unmarshal([]byte(currentAdmin.(string)),&admin)//interface -> 结构体
		param:=ctx.Params().Get("status")//获取url传入的参数
		//fmt.Println("status=",status)
		status, _ :=strconv.Atoi(param)
		records:=model.GetUnverifiedHospitals(status)
		ctx.ViewData("hospitals",util.ParseJson(records))
	}
	ctx.View("admin/VerifyHospital.html")
}

func PassHospital(ctx iris.Context)  {//审核医院注册通过
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateHospitalStatus(Ids[i],1)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/admin/verifyHospitals';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func FailHospital(ctx iris.Context)  {//审核医院注册不通过
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateHospitalStatus(Ids[i],-1)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/admin/verifyHospitals';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}