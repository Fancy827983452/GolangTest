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
	"fmt"
	"encoding/json"
)

func DoctorMain(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor", currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	if(doctor.Status==2 || doctor.Status==3){ //空闲或忙碌状态
		//获取今天日期
		dateStr:=time.Now().Format("2006-01-02 15:04:05")
		date:=strings.Split(dateStr," ")[0]
		appoints:=model.DoctorViewAppointments(doctor.HospitalId,doctor.DeptId,doctor.DoctorKey,0,date)
		ctx.ViewData("appoints",util.ParseJson(appoints))
	}
	ctx.View("doctor/DoctorMain.html")
}

func DoctorMainPost(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	if doctor.Status == 1 || doctor.Status == 4 {
		model.UpdateDoctorStatus(doctor.DoctorKey, 2) //更改status为2（空闲）
		doctor.Status=2;
	} else if doctor.Status == 2 || doctor.Status == 3 {
		model.UpdateDoctorStatus(doctor.DoctorKey, 4) //更改status为4（挂起）
		doctor.Status=4;
	}
	//更新session
	session.Set("currentDoctor", util.ParseJson(doctor))
	ctx.HTML("<script>window.location.href='/doctor/main';</script>")
}

func DepartmentManagement(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	//根据医院代码和科室代码读取所有在职医生
	doctors:=model.GetDeptValidDoctors(doctor.HospitalId,doctor.DeptId)
	ctx.ViewData("doctors",util.ParseJson(doctors))
	ctx.View("doctor/DoctorDepartmentManagement.html")
}

func SetDeptArrangement(ctx iris.Context){ //安排出诊
	var msg string
	selectedItems:=ctx.FormValue("selectedItem")//获取选中的id
	//获取选择的时间
	arrange, _ :=strconv.Atoi(ctx.FormValue("arrange"))
	if len(selectedItems)>0 {
		//fmt.Println("selectedItems="+selectedItems)
		Ids:=strings.Split(selectedItems,",")//切割取出每一个id
		//fmt.Println("Ids=",Ids)
		var length=len(Ids)
		//遍历Ids，挨个做update
		for i:=0;i<length;i++ {
			model.UpdateDoctorArrange(Ids[i],arrange)
		}
		msg="操作成功！"
		ctx.HTML("<script>alert('"+msg+"');window.location.href='/doctor/departmentManagement';</script>")
	} else {
		fmt.Println("selectedItems is empty")
	}
}

func ViewDepartmentArrangement(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	//根据医院代码和科室代码读取所有在职医生
	doctors:=model.GetDeptValidDoctors(doctor.HospitalId,doctor.DeptId)
	ctx.ViewData("doctors",util.ParseJson(doctors))
	ctx.View("doctor/DoctorViewAllWorktime.html")
}

func ViewAppointmentNum(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	//获取最大挂号数
	num, _ :=model.GetSettedAppointNum(doctor.HospitalId,doctor.DeptId)
	ctx.ViewData("num",num)
	ctx.View("doctor/DoctorSetAppointmentNumber.html")
}

func SetAppointmentNum(ctx iris.Context){
	num, _ :=strconv.Atoi(ctx.FormValue("num"))
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	result, _ :=model.UpdateSettedAppointNum(num,doctor.HospitalId,doctor.DeptId)
	var msg string
	if result > 0 {
		msg = "更新成功！"
		ctx.HTML("<script>alert('" + msg + "');window.location.href='/doctor/setAppointmentNum';</script>")
	} else {
		msg="没有做任何更改！"
		ctx.HTML("<script>alert('"+msg+"');window.history.back(-1);</script>")
	}
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
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	records:=model.GetMedicalRecordByDoctor(doctor.DoctorKey)
	ctx.ViewData("history",util.ParseJson(records))
	ctx.View("doctor/DoctorTreatmentRecord.html")
}

func TreatmentHistorySearch(ctx iris.Context){
	session:=sessionMgr.BeginSession(ctx.ResponseWriter(),ctx.Request())
	currentDoctor:=session.Get("currentDoctor")
	ctx.ViewData("currentDoctor",currentDoctor)
	//interface -> 结构体
	ctx.ResponseWriter().Header().Set("content-type", "text/html")
	doctor := model.Doctor{}
	json.Unmarshal([]byte(currentDoctor.(string)), &doctor)
	//获取搜索的病人姓名
	patientName:=ctx.FormValue("patientName")
	records:=model.SearchMedicalRecordByDoctor(doctor.DoctorKey,patientName)
	ctx.ViewData("history",util.ParseJson(records))
	ctx.View("doctor/DoctorTreatmentRecord.html")
}

