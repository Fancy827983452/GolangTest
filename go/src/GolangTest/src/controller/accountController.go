package controller

import (
	"model"
	"util"
	"algorithm"
	"encoding/hex"
	"github.com/kataras/iris"
	//"github.com/kataras/iris/sessions"
	"strconv"
	//"fmt"
)

func Login(ctx iris.Context) {
	ctx.View("Login.html")
}

func Register(ctx iris.Context) {
	ctx.View("Register.html")
}

// 注册：Post提交Register
func RegisterPost(ctx iris.Context) {
	var u model.User
	var msg model.Uploador
	u.Name = ctx.FormValue("username")
	u.Password = ctx.FormValue("password")
	confirmPassword:=ctx.FormValue("confirmPassword");
	u.IdNum = ctx.FormValue("idnumber")
	u.PhoneNum = ctx.FormValue("telephone")
	u.Gender, _ =strconv.Atoi(ctx.FormValue("sex"))
	u.BirthDate=ctx.FormValue("birthdate")
	u.Location=ctx.FormValue("location")
	u.Account=0
	u.PublicKey="tmp"

	if confirmPassword==u.Password {//两次密码输入一致
		privateKey,publicKey,err:=algorithm.GetKey();
		err=algorithm.SavePrivateKey("privateKey",privateKey)//保存私钥到本地
		util.CheckErr(err)
		u.PublicKey=hex.EncodeToString(algorithm.PublicKeyToByte(publicKey))//设置user_key字段为公钥
		u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
		result, err := model.Register(u);//插入数据库，返回操作结果（true或false）
		util.CheckErr(err)
		msg.Success = result
		msg.Message="注册成功！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	} else{
		msg.Success = false
		msg.Message="两次输入密码不一致！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}
}


