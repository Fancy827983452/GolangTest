package controller

import (
	"github.com/kataras/iris"
	"strings"
	"model"
	"util"
	"algorithm"
	"encoding/hex"
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
		user, err := model.LoginPost(u);//查询数据库，返回User对象
		util.CheckErr(err)
		if user.PublicKey=="" {
			msg.Success = false
			msg.Message="用户名或密码错误！"
		}	else {
			msg.Success = true
			msg.Message="登陆成功！"
			msg.Path="PublicKey="+user.PublicKey+",Name="+u.Name;
		}
	}
	ctx.JSON(msg)
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
	if confirmPassword==u.Password {//两次密码输入一致
		if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Password) == "" {//必填项不为空
			msg.Success = false
			msg.Message="请填写所有带*的必填项！"
		} else {
			privateKey,publicKey,err:=algorithm.GetKey();
			err=algorithm.SavePrivateKey("privateKey",privateKey)//保存私钥到本地
			util.CheckErr(err)
			err=algorithm.SavePublicKey("publicKey",publicKey)//保存公钥到本地
			util.CheckErr(err)
			u.PublicKey=hex.EncodeToString(algorithm.PublicKeyToByte(publicKey))//设置user_key字段为公钥
			u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
			result, err := model.RegisterPost(u);//插入数据库，返回操作结果（true或false）
			util.CheckErr(err)
			msg.Success = result
			msg.Message="注册成功！"
		}
	}
	ctx.JSON(msg)
}


