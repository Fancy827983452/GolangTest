package controller

import (
	"github.com/kataras/iris"
	"encoding/hex"
	"algorithm"
	"model"
	"util"
	"strconv"
)

func EditUserInfo(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.View("user/UserEditInfo.html")
}

func EditUserInfoPost(ctx iris.Context) {
	var u model.User
	var msg model.Uploador
	u.PublicKey = ctx.FormValue("publicKey")
	u.Name = ctx.FormValue("username")
	u.Gender, _ = strconv.Atoi(ctx.FormValue("sex"))
	u.PhoneNum = ctx.FormValue("tel")
	u.BirthDate = ctx.FormValue("birthdate")
	u.Location = ctx.FormValue("location")
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	var Ace_Key string
	Ace_Key=session.Get("AEC_KEY").(string)
	u.Location=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.Location),[]byte(Ace_Key)))
	result, _ := model.UpdateInfo(u); //插入数据库，返回操作结果（true或false）
	if result > 0 {
		msg.Success = true
		msg.Message = "个人信息更新成功！"
		user, err := model.GetInfoByPublicKey(u)
		util.CheckErr(err)
		idnum, _ :=hex.DecodeString(user.IdNum)
		user.IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(Ace_Key)))
		location, _ :=hex.DecodeString(user.Location)
		user.Location=string(algorithm.AEC_CRT_Crypt(location,[]byte(Ace_Key)))
		session := sessionMgr.BeginSession(ctx.ResponseWriter(), ctx.Request())
		session.Set("currentUser", util.ParseJson(user)) //更新session
		ctx.HTML("<script>alert('" + msg.Message + "');" +
			"window.location.href='editInfo';</script>")
	} else {
		msg.Message="个人信息更新失败！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}
}

func EditUserPwd(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.View("user/UserEditPassword.html")
}

func EditUserPwdPost(ctx iris.Context){
	var u model.User
	var msg string
	u.PublicKey = ctx.FormValue("publicKey")
	u.Password=ctx.FormValue("password_old")
	u.Password = algorithm.GetMd5String(u.Password)
	user, _ :=model.GetInfoByPublicKey(u)

	if u.Password!=user.Password{
		msg="原密码错误！"
	}else {
		u.Password=ctx.FormValue("password_new")
		u.Password = algorithm.GetMd5String(u.Password)
		result, _ :=model.UpdatePwd(u)
		if(result>0){
			msg="更改密码成功！请重新登录！"
			sessionMgr.Destroy(ctx.ResponseWriter(),ctx.Request())
			ctx.HTML("<script>alert('"+msg+"');window.location.href='/login';</script>")
		} else{
			msg="更改密码失败！"
		}
	}
	ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
}

func UserIllcase(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.View("user/Userillcase.html")
}

func VisitUserRecord(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.View("user/UserVisitRecord.html")
}
