package controller

import (
	"github.com/kataras/iris"
	"encoding/hex"
	"algorithm"
	"model"
	"util"
	"strings"
	"strconv"
	"time"
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
	//u.Gender, _ = strconv.Atoi(ctx.FormValue("sex"))
	u.PhoneNum = ctx.FormValue("tel")
	//u.BirthDate = ctx.FormValue("birthdate")
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
		msg.Message="没有做任何更改！"
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

func VisitUserRecord(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	ctx.View("user/UserVisitRecord.html")
}

func UserAppointment(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)

	//获取医院列表
	records:=model.GetAllHospitals()
	ctx.ViewData("hospitals",util.ParseJson(records))

	//根据医院id读取科室列表
	param:=ctx.Params().Get("param") //获取传递的参数：医院代码&科室代码&周几
	str:=strings.Split(param,"&")
	if(len(str)>0){
		deptList:=model.GetHospitalDepts(str[0]);
		ctx.ViewData("departments",util.ParseJson(deptList))
		//if len(str)>2 {
		//	if str[2]=="-1"{
		//		//获取当前医院的当前科室的所有医生
		//		doctorList:=model.GetDeptValidDoctors(str[0],str[1])
		//		ctx.ViewData("doctors",util.ParseJson(doctorList))
		//	} else {
		//		//判断这一天是周几
		//		str:=strings.Split(str[2],"-")
		//		year, _ :=strconv.Atoi(str[0])
		//		month, _ :=strconv.Atoi(str[1])
		//		day, _ :=strconv.Atoi(str[2])
		//		w:=util.ZellerFunction2Week(year,month,day)
		//		//获取当天排班的医生
		//		doctorList:=model.GetDeptArrangeDoctors(str[0],str[1],string(w))
		//		ctx.ViewData("doctors",util.ParseJson(doctorList))
		//	}
		//}
	}
	ctx.View("user/UserillAppointment.html")
}

func UserAppointmentSearchPost(ctx iris.Context) {
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentUser:=session.Get("currentUser")
	ctx.ViewData("currentUser",currentUser)
	//获取医院列表
	records:=model.GetAllHospitals()
	ctx.ViewData("hospitals",util.ParseJson(records))

	//根据医院id读取科室列表
	param:=ctx.Params().Get("param") //获取传递的参数：医院代码&科室代码&周几
	str:=strings.Split(param,"&")
	if(len(str)>0){
		deptList:=model.GetHospitalDepts(str[0]);
		ctx.ViewData("departments",util.ParseJson(deptList))
	}

	hospitalId:=ctx.FormValue("hospitalName")
	deptId:=ctx.FormValue("deptName")
	day:=ctx.FormValue("selectDate")

	var remain int
	var count int
	//if(day=="-1"){ //如果时间选择的是不限
	//	dates:=util.GetRecentDays(7)//获取近7天的日期
	//	for i:=0;i<len(dates);i++{
	//		count, _ =model.GetCurrentAppointedNum(hospitalId,deptId,dates[i])//获取指定日期的已预约数量
	//		if(count<max){ //如果没有约满
	//			remain=max-count
	//			//判断这一天是周几
	//			str:=strings.Split(dates[i],"-")
	//			year, _ :=strconv.Atoi(str[0])
	//			month, _ :=strconv.Atoi(str[1])
	//			day, _ :=strconv.Atoi(str[2])
	//			w:=util.ZellerFunction2Week(year,month,day)
	//			//获取当天上班的医生信息
	//			doctorList:=model.GetArrangedDoctor(hospitalId,deptId,w)
	//			ctx.ViewData("doctors"+string(i+1),util.ParseJson(doctorList))
	//			ctx.ViewData("date"+string(i+1),dates[i])
	//			ctx.ViewData("remain"+string(i+1),remain)
	//		}else {
	//			remain=0
	//			ctx.ViewData("date"+string(i+1),dates[i])
	//			ctx.ViewData("remain"+string(i+1),remain)
	//		}
	//	}
	//} else if(day!="-1"){//如果指定了时间

	//判断当天是否约满
	//获取最大号数
	max, _ :=model.GetSettedAppointNum(hospitalId,deptId)
	//获取当前已挂号数
	count, _ =model.GetCurrentAppointedNum(hospitalId,deptId,day)
	if(count<max) { //如果没有约满
		remain=max-count
		//判断这一天是周几
		str := strings.Split(day, "-")
		year, _ := strconv.Atoi(str[0])
		month, _ := strconv.Atoi(str[1])
		day1, _ := strconv.Atoi(str[2])
		w := util.ZellerFunction2Week(year, month, day1)
		//获取当天上班的医生信息
		doctorList := model.GetArrangedDoctor(hospitalId, deptId, w)
		ctx.ViewData("doctors",util.ParseJson(doctorList))
		ctx.ViewData("remain",remain)
		ctx.ViewData("hospitalId",hospitalId)
		ctx.ViewData("deptId",deptId)
		ctx.ViewData("day",day)
	}else {
		remain=0
		ctx.ViewData("remain",remain)
	}
	ctx.View("user/UserillAppointment.html")
}

func UserAddAppointment(ctx iris.Context) {
	hospitalId:=ctx.FormValue("hospitalID");
	deptId:=ctx.FormValue("deptID");
	day:=ctx.FormValue("DAY");
	doctorKey:=ctx.FormValue("doctorKey")
	userKey:=ctx.FormValue("publicKey")

	var item model.Appointment
	//获取当前已挂号数量
	appointed, _ :=model.GetCurrentAppointedNum(hospitalId,deptId,day)
	item.Number=strconv.Itoa(appointed+1)
	item.Time=time.Now().Format("2006-01-02 15:04:05")
	item.HospitalId, _ =strconv.Atoi(hospitalId)
	item.DeptId, _ =strconv.Atoi(deptId)
	item.DoctorKey=doctorKey
	item.PatientKey=userKey
	item.AppointDate=day

	result, _ :=model.AddAppointment(item)
	var msg string
	if(result==true){
		msg="预约成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/user/illcase'</script>")
	}

}