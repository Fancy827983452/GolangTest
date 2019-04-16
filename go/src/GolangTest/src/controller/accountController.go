package controller

import (
	"model"
	"util"
	"algorithm"
	"encoding/hex"
	"github.com/kataras/iris"
	"strconv"
)

func Login(ctx iris.Context) {
	ctx.View("Login.html")
}

func Register(ctx iris.Context) {
	ctx.View("Register.html")
}

func LoginPost (ctx iris.Context) {
	var u model.User
	var msg model.Uploador
	u.Name = ctx.FormValue("username")
	u.Password = ctx.FormValue("password")
	u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
	count,err:=model.CheckLogin(u)
	util.CheckErr(err)
	if count==0{
		msg.Success = false
		msg.Message="用户名或密码错误！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	} else {
		user,err := model.Login(u);
		util.CheckErr(err)
		msg.Success = true
		msg.Message="登陆成功！"
		//获取session管理器
		session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
		idnum, _ :=hex.DecodeString(user.IdNum)
		user.IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(user.Ace_Key)))
		location, _ :=hex.DecodeString(user.Location)
		user.Location=string(algorithm.AEC_CRT_Crypt(location,[]byte(user.Ace_Key)))
		session.Set("currentUser",util.ParseJson(user))
		session.Set("AEC_KEY",user.Ace_Key)
		ctx.HTML("<script>alert('"+msg.Message+"');window.location.href='user/editInfo';</script>")//URL传参
	}
}

// 注册：Post提交Register
func RegisterPost(ctx iris.Context) {
	var u model.User
	var msg model.Uploador
	u.Name = ctx.FormValue("username")
	u.Password = ctx.FormValue("password")
	u.IdNum = ctx.FormValue("idnumber")
	u.PhoneNum = ctx.FormValue("telephone")
	u.Gender, _ =strconv.Atoi(ctx.FormValue("sex"))//string转int
	u.BirthDate=ctx.FormValue("birthdate")
	u.Location=ctx.FormValue("location")
	u.PublicKey="tmp"

	count, _ :=model.CheckLogin(u)
	if(count>0) {//判断手机号是否已经注册过
		msg.Success = false
		msg.Message="该手机号已注册过！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}else {
		privateKey,publicKey,err:=algorithm.GetKey();
		err=algorithm.SavePrivateKey("privateKey",privateKey)//保存私钥到本地
		util.CheckErr(err)
		u.PublicKey=hex.EncodeToString(algorithm.PublicKeyToByte(publicKey))//设置user_key字段为公钥
		u.Ace_Key=algorithm.GetRandomString(16)//生成随机ACE秘钥
		//使用AEC算法对称加密数据
		//u.Name=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.Name),[]byte(u.Ace_Key)))
		u.IdNum=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.IdNum),[]byte(u.Ace_Key)))
		//u.PhoneNum=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.PhoneNum),[]byte(u.Ace_Key)))
		u.Location=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.Location),[]byte(u.Ace_Key)))
		u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
		u.Addr=hex.EncodeToString(algorithm.GetAddress(u.PublicKey))//base58根据公钥生成地址

		//fmt.Println("u.Addr="+u.Addr)

		result, err := model.Register(u);//插入数据库，返回操作结果（true或false）
		util.CheckErr(err)
		msg.Success = result
		msg.Message="注册成功！"
		ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
	}
}

func Logout(ctx iris.Context) {
	sessionMgr.Destroy(ctx.ResponseWriter(),ctx.Request())
	ctx.HTML("<script>alert('登出成功！');window.location.href='/login';</script>")
}


