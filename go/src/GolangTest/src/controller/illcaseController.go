package controller

import (
	"github.com/kataras/iris"
	"model"
	"util"
	"encoding/json"
	"strings"
	"strconv"
	"fmt"
)

func UserIllcase(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.ResponseWriter().Header().Set("content-type", "text/html")

	user:=model.User{}
	json.Unmarshal([]byte(currentUser.(string)),&user)//interface -> 结构体
	records:=model.GetMedicalRecordByUser(user.PublicKey)
	ctx.ViewData("illcase",util.ParseJson(records))
	ctx.View("user/Userillcase.html")
}

func LockRecord(ctx iris.Context){
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			id, _ :=strconv.Atoi(Ids[i])
			//fmt.Println("id=",id)
			model.UpdateStatus(id,1)
		}
			msg="锁定记录成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/user/illcase';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func UnlockRecord(ctx iris.Context){
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			id, _ :=strconv.Atoi(Ids[i])
			//fmt.Println("id=",id)
			model.UpdateStatus(id,0)
		}
		msg="解锁记录成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/user/illcase';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}
