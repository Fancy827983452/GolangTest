package controller

import (
	"github.com/kataras/iris"
	"strings"
	"model"
	"util"
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
		u.Password = util.GetMd5String(u.Password)//MD5加密password
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
			count,err:=model.GetUserCount();
			util.CheckErr(err)
			u.PublicKey=strconv.Itoa(count+1)//获取当前user表中的条目数，加一，转为string
			u.PublicKey=util.GetMd5String(u.PublicKey)//生成公钥
			u.Password = util.GetMd5String(u.Password)//MD5加密password
			result, err := model.RegisterPost(u);//插入数据库，返回操作结果（true或false）
			util.CheckErr(err)
			msg.Success = result
			msg.Message="注册成功！"
		}
	}
	ctx.JSON(msg)
}


