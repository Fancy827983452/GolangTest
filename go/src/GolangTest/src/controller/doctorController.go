package controller

import (
	"github.com/kataras/iris"
	"model"
	"encoding/hex"
	"algorithm"
	"util"
	"strings"
	"time"
	"strconv"
)

func DoctorMain(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	doctorStatus:=session.Get("doctorStatus")
	ctx.ViewData("currentDoctor", currentDoctor)
	ctx.ViewData("doctorStatus", doctorStatus)
	if(doctorStatus.(int)==2 || doctorStatus.(int)==3){ //空闲或忙碌状态
		str:=currentDoctor.(string)
		record:=strings.Split(str,",")[0]//取DoctorKey的key和value
		value:=strings.Split(record,":")[1]//取value
		//获取医生公钥
		doctorKey:=value[1:len(value)-1]//去除前后的双引号
		//读取数据库中今天该医院、科室、医生、status=0（未看诊）的预约记录
		//获取医院id（不用去除前后的双引号，因为结构体中定义的是int型）
		hospitalId:=strings.Split(strings.Split(str,",")[7],":")[1]
		//获取科室id
		deptId:=strings.Split(strings.Split(str,",")[9],":")[1]
		//获取今天日期
		dateStr:=time.Now().Format("2006-01-02 15:04:05")
		date:=strings.Split(dateStr," ")[0]
		appoints:=model.DoctorViewAppointments(hospitalId,deptId,doctorKey,0,date)
		ctx.ViewData("appoints",util.ParseJson(appoints))
	}
	ctx.View("doctor/DoctorMain.html")
}

func DoctorMainPost(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	doctorStatus:=session.Get("doctorStatus")
	str := currentDoctor.(string)
	record := strings.Split(str, ",")[0]   //取DoctorKey的key和value
	value := strings.Split(record, ":")[1] //取value
	//获取医生公钥
	doctorKey := value[1:len(value)-1] //去除前后的双引号
	if (doctorStatus.(int) == 1 || doctorStatus.(int) == 4) {
		model.UpdateDoctorStatus(doctorKey, 2) //更改status为2（空闲）
	} else if (doctorStatus.(int) == 2 || doctorStatus.(int) == 3) {
		model.UpdateDoctorStatus(doctorKey, 4) //更改status为4（挂起）
	}
	//更新session
	doctor1, err := model.GetDoctorInfoByPublicKey(doctorKey)
	util.CheckErr(err)
	idnum, _ :=hex.DecodeString(doctor1.IdNum)
	doctor1.IdNum=string(algorithm.AEC_CRT_Crypt(idnum,[]byte(doctor1.Aec_Key)))
	session.Set("currentDoctor", util.ParseJson(doctor1))
	session.Set("doctorStatus", doctor1.Status)
	ctx.HTML("<script>window.location.href='/doctor/main';</script>")
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
		doctor1, err := model.GetDoctorInfoByPublicKey(doctor.DoctorKey)
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
	user, _ :=model.GetDoctorInfoByPublicKey(doctor.DoctorKey)

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

func AddCase(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)

	appointmentId:=ctx.Params().Get("appointmentId")
	//根据预约号读取患者信息
	user, _ :=model.GetAppointedUserInfo(appointmentId)
	ctx.ViewData("patient",util.ParseJson(user))
	ctx.ViewData("appointmentId",appointmentId)
	ctx.View("doctor/DoctorAddCase.html")
}

func AddCasePost(ctx iris.Context){
	var mr model.MedicalRecord
	var msg string
	mr.DeseaseName=ctx.FormValue("illname")
	mr.Symptom=ctx.FormValue("illdescribe")
	mr.Info=ctx.FormValue("illdetail")
	appointId:=ctx.FormValue("appoint")
	mr.AppointmentId, _ =strconv.Atoi(appointId)
	//获取今天日期
	mr.Time=time.Now().Format("2006-01-02 15:04:05")
	mr.Status=0
	result, _ :=model.AddMedicalRecord(mr)
	if result==true {
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/doctor/main';</script>")
	}else {
		msg="操作失败！"
		ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
	}
}

func TreatmentHistory(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	str:=currentDoctor.(string)
	record:=strings.Split(str,",")[0]//取DoctorKey的key和value
	value:=strings.Split(record,":")[1]//取value
	//获取医生公钥
	doctorKey:=value[1:len(value)-1]//去除前后的双引号
	records:=model.GetMedicalRecordByDoctor(doctorKey)
	ctx.ViewData("history",util.ParseJson(records))

	ctx.View("doctor/DoctorTreatmentRecord.html")
}

func TreatmentHistorySearch(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	str:=currentDoctor.(string)
	record:=strings.Split(str,",")[0]//取DoctorKey的key和value
	value:=strings.Split(record,":")[1]//取value
	//获取医生公钥
	doctorKey:=value[1:len(value)-1]//去除前后的双引号
	//获取搜索的病人姓名
	patientName:=ctx.FormValue("patientName")
	records:=model.SearchMedicalRecordByDoctor(doctorKey,patientName)
	ctx.ViewData("history",util.ParseJson(records))

	ctx.View("doctor/DoctorTreatmentRecord.html")
}

