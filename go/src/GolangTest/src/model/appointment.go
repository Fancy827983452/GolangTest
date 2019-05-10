package model

import (
	"util"
)

type Appointment struct {
	AppointmentId  int		//预约代码
	Number 	    string  	//号码
	Time 		string  	//挂号时间
	AppointDate string      //预约日期
	DeptId		int  		//科室代码
	DeptName    string		//科室名
	HospitalId  int			//医院代码
	HospitalName string		//医院名
	DoctorKey   string		//医生的公钥（如果挂号时未指定医生，则为空）
	DoctorName  string      //医生名
	PatientKey  string		//病人的公钥
	PatientName string		//病人名
	Status      int			//挂号的状态（0：未看诊；1：已看诊；2：未看诊且过期）
}

type Appointments struct {
	Items    []*Appointment
}

//获取指定医院科室的当前挂号信息
func QueryAppoints(hospitalId string,deptId string,date string,doctorKey string) *Appointments {
	var result Appointments
	result.Items = []*Appointment{}
	var str string //判断有没有指定医生
	//if doctorKey=="-1"{
	//	str=""//如果没有指定医生
	//}else {
		str=" and appointment.doctor_key="+doctorKey
	//}
	query := "select id,number,time,appointDate,appointment.department_id,department_name,appointment.medical_institution_id,tbl_medical_institution.name, " +
		"appointment.status,appointment.doctor_key,tbl_doctor.name,patient_key,tbl_user.name from appointment " +
		"join tbl_doctor on appointment.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_user on patient_key=tbl_user.user_key " +
		"join tbl_medical_institution on tbl_medical_institution.medical_institution_id=appointment.medical_institution_id " +
		"join tbl_medical_institution_department on tbl_medical_institution_department.department_id=appointment.department_id " +
		"where appointment.department_id=? and appointment.medical_institution_id=? and appointDate="+ date +str
	rows, err := db.Query(query,deptId,hospitalId)
	util.CheckErr(err)
	for rows.Next() {
		item := Appointment{}
		err = rows.Scan(&item.AppointmentId,&item.Number,&item.Time,&item.AppointDate,&item.DeptId,&item.DeptName,&item.HospitalId,&item.HospitalName,&item.Status,&item.DoctorKey,&item.DoctorName,&item.PatientKey,&item.PatientName)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//获取指定医院科室指定日期的当前挂号数量
func GetCurrentAppointedNum(hospitalId string,deptId string,date string) (int,error) {
	sql:="select ifnull(count(*),0) as count from appointment where department_id=? and medical_institution_id=? and appointDate=?"
	var count int
	err := db.QueryRow(sql,hospitalId,deptId,date).Scan(&count)
	util.CheckErr(err)
	return count,err
}

//插入记录
func AddAppointment(item Appointment) (bool, error){
	sql :="insert into appointment(number,time,department_id,medical_institution_id,status,doctor_key,patient_key,appointDate) values(?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql,&item.Number,&item.Time,&item.DeptId,&item.HospitalId,0,&item.DoctorKey,&item.PatientKey,&item.AppointDate)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//读取数据库中今天该医院、科室、医生、status=0（未看诊）的预约记录
func DoctorViewAppointments(hospitalId string,deptId string,doctorKey string,status int,date string) *Appointments {
	var result Appointments
	result.Items = []*Appointment{}
	query := "select id,number,time,appointDate,appointment.department_id,department_name," +
		"appointment.medical_institution_id,tbl_medical_institution.name," +
		"appointment.status,appointment.doctor_key,tbl_doctor.name,patient_key,tbl_user.name " +
		"from appointment join tbl_doctor on appointment.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_user on patient_key=tbl_user.user_key " +
		"join tbl_medical_institution on tbl_medical_institution.medical_institution_id=appointment.medical_institution_id " +
		"join tbl_medical_institution_department on tbl_medical_institution_department.department_id=appointment.department_id " +
		"where appointment.department_id=? and appointment.medical_institution_id=? and appointDate=? " +
		"and appointment.doctor_key=? and appointment.status=?;"
	rows, err := db.Query(query,deptId,hospitalId,date,doctorKey,status)
	util.CheckErr(err)
	for rows.Next() {
		item := Appointment{}
		err = rows.Scan(&item.AppointmentId,&item.Number,&item.Time,&item.AppointDate,&item.DeptId,&item.DeptName,&item.HospitalId,&item.HospitalName,&item.Status,&item.DoctorKey,&item.DoctorName,&item.PatientKey,&item.PatientName)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}