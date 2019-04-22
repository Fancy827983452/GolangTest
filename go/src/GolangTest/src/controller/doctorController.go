package controller

import (
	"github.com/kataras/iris"
	"model"
	"encoding/hex"
	"algorithm"
	"util"
)

func DoctorMain(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorMain.html")
}

func DepartmentManagement(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorDepartmentManagement.html")
}

func ViewDepartmentArrangement(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorViewAllWorktime.html")
}

func SetAppointmentNum(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorSetAppointmentNumber.html")
}

func PatientDetails(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/PatientDetails.html")
}

func DoctorEditInfo(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorEditInfo.html")
}

func DoctorEditInfoPost(ctx iris.Context) {
	var doctor model.Doctor
	var msg string
	doctor.DoctorKey = ctx.FormValue("doctorKey")
	doctor.PhoneNum = ctx.FormValue("tel")

	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	var Ace_Key string
	Ace_Key=session.Get("AEC_KEY").(string)

	result, _ := model.UpdateDoctotInfo(doctor); //插入数据库，返回操作结果（true或false）
	if result > 0 {
		msg = "个人信息更新成功！"
		doctor1, err := model.GetDoctorInfoByPublicKey(doctor)
		util.CheckErr(err)
		idnum, _ :=hex.DecodeString(doctor1.IdNum)
		doctor1.IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(Ace_Key)))
		session := sessionMgr.BeginSession(ctx.ResponseWriter(), ctx.Request())
		session.Set("currentDoctor", util.ParseJson(doctor1)) //更新session
		ctx.HTML("<script>alert('" + msg + "');" + "window.location.href='editInfo';</script>")
	} else {
		msg="没有做任何更改！"
		ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
	}
}

func DoctorEditPwd(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorEditPassword.html")
}

func DoctorEditPwdPost(ctx iris.Context)  {
	var doctor model.Doctor
	var msg string
	doctor.DoctorKey = ctx.FormValue("doctorKey")
	doctor.Password=ctx.FormValue("password_old")
	doctor.Password = algorithm.GetMd5String(doctor.Password)
	user, _ :=model.GetDoctorInfoByPublicKey(doctor)

	if doctor.Password!=user.Password{
		msg="原密码错误！"
	}else {
		doctor.Password=ctx.FormValue("password_new")
		doctor.Password = algorithm.GetMd5String(doctor.Password)
		result, _ :=model.UpdateDoctorPwd(doctor)
		if(result>0){
			msg="更改密码成功！请重新登录！"
			sessionMgr.Destroy(ctx.ResponseWriter(),ctx.Request())
			ctx.HTML("<script>alert('"+msg+"');window.location.href='login';</script>")
		} else{
			msg="更改密码失败！"
		}
	}
	ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
}

func VisitHistory(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/VisitHistory.html")
}

func PatientHistoryCase(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorViewHistoryCase.html")
}

func AddCase(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorAddCase.html")
}

func PatientTreatmentHistory(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	ctx.View("doctor/DoctorTreatmentRecord.html")
}