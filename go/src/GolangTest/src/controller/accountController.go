package controller

import (
	"model"
	"util"
	"algorithm"
	"encoding/hex"
	"github.com/kataras/iris"
	"strconv"
)

//用户
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
		//birth, _ :=hex.DecodeString(user.BirthDate)
		//user.BirthDate=string(algorithm.AEC_CRT_Crypt(birth,[]byte(user.Ace_Key)))
		location, _ :=hex.DecodeString(user.Location)
		user.Location=string(algorithm.AEC_CRT_Crypt(location,[]byte(user.Ace_Key)))
		session.Set("currentUser",util.ParseJson(user))
		session.Set("AEC_KEY",user.Ace_Key)
		ctx.HTML("<script>alert('"+msg.Message+"');window.location.href='user/editInfo';</script>")//URL传参
	}
}

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
		privateKey,publicKey,_:=algorithm.GetKey();
		u.PublicKey=hex.EncodeToString(algorithm.PublicKeyToByte(publicKey))//设置user_key字段为公钥
		pk := algorithm.GetMd5String(u.PublicKey)//MD5加密公钥

		u.Ace_Key=algorithm.GetRandomString(16)//生成随机ACE秘钥
		//使用AEC算法对称加密数据
		//u.Name=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.Name),[]byte(u.Ace_Key)))
		u.IdNum=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.IdNum),[]byte(u.Ace_Key)))
		//u.BirthDate=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.BirthDate),[]byte(u.Ace_Key)))
		//u.PhoneNum=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.PhoneNum),[]byte(u.Ace_Key)))
		u.Location=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(u.Location),[]byte(u.Ace_Key)))
		u.Password = algorithm.GetMd5String(u.Password)//MD5加密password
		u.Addr=hex.EncodeToString(algorithm.GetAddress(u.PublicKey))//base58根据公钥生成地址

		result, _ := model.Register(u);//插入数据库，返回操作结果（true或false）
		if(result==true){
			msg.Success = result
			msg.Message="注册成功！"
			algorithm.SavePrivateKey("privateKey_"+pk,privateKey)//保存私钥到本地
			ctx.HTML("<script>alert('"+msg.Message+"');window.history.back(-1);</script>")
		}
	}
}

//医生
func DoctorRegister(ctx iris.Context){
	ctx.View("doctor/DoctorRegister.html")
}

func DoctorRegisterPost(ctx iris.Context){
	var doctor model.Doctor
	var msg string
	doctor.Name=ctx.FormValue("username")
	doctor.HospitalId, _ =strconv.Atoi(ctx.FormValue("hospitalId"))
	doctor.DeptId, _ =strconv.Atoi(ctx.FormValue("departmentId"))
	doctor.Gender, _ =strconv.Atoi(ctx.FormValue("sex"))
	doctor.BirthDate=ctx.FormValue("birthdate")
	doctor.IdNum=ctx.FormValue("idnumber")
	doctor.PhoneNum=ctx.FormValue("telephone")
	doctor.Password=ctx.FormValue("password")
	doctor.Status=0;
	doctor.Role=0;
	doctor.Title, _ =strconv.Atoi(ctx.FormValue("title"))

	//判断手机号是否被注册过
	count1, _ :=model.CheckDoctorPhone(doctor)
	if count1>0 {
		msg="改手机号已注册过！";
		ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
	}else {
		//判断医院代码和科室代码是否存在
		count2, _ :=model.CheckDoctorIDs(doctor)
		if count2==0{
			msg="医院代码或科室代码错误！"
			ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
		} else {
			//公钥
			privateKey,publicKey,_:=algorithm.GetKey();
			doctor.DoctorKey=hex.EncodeToString(algorithm.PublicKeyToByte(publicKey))//设置user_key字段为公钥
			pk := algorithm.GetMd5String(doctor.DoctorKey)//MD5加密公钥
			//生成随机ACE秘钥
			doctor.Aec_Key=algorithm.GetRandomString(16)
			//使用AEC算法对称加密数据
			doctor.IdNum=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(doctor.IdNum),[]byte(doctor.Aec_Key)))
			//doctor.BirthDate=hex.EncodeToString(algorithm.AEC_CRT_Crypt([]byte(doctor.BirthDate),[]byte(doctor.Aec_Key)))
			//base58地址
			doctor.Addr=hex.EncodeToString(algorithm.GetAddress(doctor.DoctorKey))//base58根据公钥生成地址
			//MD5加密密码
			doctor.Password = algorithm.GetMd5String(doctor.Password)
			//插入数据库，返回操作结果（true或false）
			result, err := model.DoctorRegister(doctor);
			util.CheckErr(err)
			if result==true{
				msg="注册成功！请等待医院方审核！"
				algorithm.SavePrivateKey("privateKey_"+pk,privateKey)//保存私钥到本地
			}else {
				msg="注册失败！"
			}
			ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
		}
	}
}

func DoctorLogin(ctx iris.Context){
	ctx.View("doctor/DoctorLogin.html")
}

func DoctorLoginPost(ctx iris.Context){
	var doctor model.Doctor
	var msg string
	doctor.Name=ctx.FormValue("username")
	doctor.Password=ctx.FormValue("password")
	doctor.Password = algorithm.GetMd5String(doctor.Password)//MD5加密password
	//fmt.Println("doctorname="+doctor.Name)
	//fmt.Println("password="+doctor.Password)
	count,_:=model.CheckDoctorLogin(doctor)
	if count==0{
		msg="用户名或密码错误！"
		ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
	} else {
		doctor,err := model.DoctorLogin(doctor);
		util.CheckErr(err)
		if doctor.Status==0{
			msg="请耐心等待院方审核！"
			ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
		}else {
			msg="欢迎您，"+doctor.Name+"！"
			//获取session管理器
			session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
			idnum, _ :=hex.DecodeString(doctor.IdNum)
			doctor.IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(doctor.Aec_Key)))
			//birth, _ :=hex.DecodeString(doctor.BirthDate)
			//doctor.BirthDate=string(algorithm.AEC_CRT_Crypt(birth,[]byte(doctor.Aec_Key)))
			session.Set("currentDoctor",util.ParseJson(doctor))
			session.Set("AEC_KEY",doctor.Aec_Key)
			ctx.HTML("<script>alert('"+msg+"');window.location.href='main';</script>")
		}
	}
}

func Logout(ctx iris.Context) {
	sessionMgr.Destroy(ctx.ResponseWriter(),ctx.Request())
	ctx.HTML("<script>alert('登出成功！');window.location.href='login';</script>")
}