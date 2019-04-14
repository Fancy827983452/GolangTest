package controller

import (
	"strings"
	"model"
	"util"
	"algorithm"
	"encoding/hex"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"strconv"
)

func Login(ctx iris.Context) {
	ctx.View("Login.html")
}

func Register(ctx iris.Context) {
	ctx.View("Register.html")
}

// 登录：Post提交Login
func LoginPost(ctx iris.Context) {
	var u model.User
	var msg model.Uploador
	u.Name = ctx.FormValue("username")
	u.Password = ctx.FormValue("password")
	if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Password) == "" {//必填项不为空
		msg.Success = false
		msg.Message="请填写用户名和密码！"
	} else {
		u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
		count,err:=model.CheckLogin(u)
		util.CheckErr(err)
		if count==0{
			msg.Success = false
			msg.Message="用户名或密码错误！"
			ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
		} else {
			user,err := model.Login(u);//获取公钥
			util.CheckErr(err)
			msg.Success = true
			msg.Message="登陆成功！"

			//创建唯一的Session ID
			var sess *sessions.Sessions
			sess= sessions.New(sessions.Config{Cookie: "mysession_cookie_name"})
			session:=sess.Start(ctx)
			session.Set("PublicKey",user.PublicKey)//session传参
			session.Set("Name",user.Name)
			session.Set("CurrentUser",user)//存储当前用户的信息

			ctx.HTML("<script>alert('"+msg.Message+"');" +
				"window.location.href='user/editInfo/"+u.Name+"';</script>")//URL传参
		}
	}
	//ctx.JSON(msg)
	//ctx.View("user/UserEditInfo.html")//url不变
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


