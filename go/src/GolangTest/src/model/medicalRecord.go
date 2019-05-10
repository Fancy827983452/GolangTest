package model

import (
	"fmt"
	"time"
	"util"
)

type MedicalRecord struct {
	ID        	int	//medical_record_id
	Time      	string	//添加时间
	Symptom     string  //症状描述
	DeseaseName string	//疾病名称
	Info      	string	//疾病详情
	DepName		string  //就诊科室名
	HospitalName string
	DoctorName 	string  //医生姓名
	UserKey   	string
	DoctorKey 	string
	PatientKey 	string
	PatientName string
	Status      int		//状态（0：默认正常；1：锁定）
	AppointmentId int   //预约代码
	PatientBirth string
	PatientGender string
}

type MedicalRecords struct {
	Items    []*MedicalRecord
}

var md MedicalRecord

//添加记录
func AddMedicalRecord(mr MedicalRecord) (bool,error) {
	sql := "insert into tbl_medical_record (add_time,desease,record_info,status,symptom,appointmentId) values (?,?,?,?,?,?)"
	res, err := db.Exec(sql,mr.Time,mr.DeseaseName,mr.Info,mr.Status,mr.Symptom,mr.AppointmentId)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//获取所有记录
func GetAllMedicalRecord() error {
	query := "select * from medical_record"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	if rows.Next() {
		if err := rows.Scan(&md.ID, &md.Time, &md.DeseaseName, &md.Info, &md.UserKey, &md.DoctorKey); err == nil {
			fmt.Println(md)
		} else {
			return err
		}
	}
	return err
}

func UpdateStatus(id int,status int) (int64, error){
	query:="update tbl_medical_record set status=? where medical_record_id=?"
	res, err := db.Exec(query, status,id)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}

//根据用户key查询病例记录
func GetMedicalRecordByUser(userKey string) *MedicalRecords {
	var result MedicalRecords
	result.Items = []*MedicalRecord{}
	query := "select medical_record_id,add_time,desease,record_info,department_name,tbl_user.name,tbl_doctor.name,tbl_medical_institution.name," +
		"appointment.patient_key,appointment.doctor_key,tbl_medical_record.status,symptom,appointmentId from tbl_medical_record " +
		"join appointment on appointmentId=appointment.id " +
		"join tbl_user on tbl_user.user_key=appointment.patient_key " +
		"join tbl_doctor on appointment.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_medical_institution_department on appointment.department_id=tbl_medical_institution_department.department_id " +
		"join tbl_medical_institution on appointment.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"where appointment.patient_key='" + userKey + "'"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		item := MedicalRecord{}
		err = rows.Scan(&item.ID, &item.Time, &item.DeseaseName, &item.Info,&item.DepName,&item.PatientName,&item.DoctorName,&item.HospitalName,&item.PatientKey,&item.DoctorKey,&item.Status,&item.Symptom,&item.AppointmentId)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//根据医生key查询病例记录
func GetMedicalRecordByDoctor(doctorKey string) *MedicalRecords {
	var result MedicalRecords
	result.Items = []*MedicalRecord{}
	query := "select medical_record_id,add_time,desease,record_info,department_name,tbl_user.name," +
		"tbl_user.birthdate,tbl_user.gender,tbl_doctor.name,tbl_medical_institution.name," +
		"appointment.patient_key,appointment.doctor_key,tbl_medical_record.status,symptom,appointmentId from tbl_medical_record " +
		"join appointment on appointmentId=appointment.id " +
		"join tbl_user on tbl_user.user_key=appointment.patient_key " +
		"join tbl_doctor on appointment.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_medical_institution_department on appointment.department_id=tbl_medical_institution_department.department_id " +
		"join tbl_medical_institution on appointment.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"where appointment.doctor_key='" + doctorKey + "' and tbl_medical_record.status=0"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		item := MedicalRecord{}
		err = rows.Scan(&item.ID, &item.Time, &item.DeseaseName, &item.Info,&item.DepName,&item.PatientName,&item.PatientBirth,&item.PatientGender,
			&item.DoctorName,&item.HospitalName,&item.PatientKey,&item.DoctorKey,&item.Status,&item.Symptom,&item.AppointmentId)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//根据医生key搜索病例记录
func SearchMedicalRecordByDoctor(doctorKey string,name string) *MedicalRecords {
	var result MedicalRecords
	result.Items = []*MedicalRecord{}
	query := "select medical_record_id,add_time,desease,record_info,department_name,tbl_user.name," +
		"tbl_user.birthdate,tbl_user.gender,tbl_doctor.name,tbl_medical_institution.name," +
		"appointment.patient_key,appointment.doctor_key,tbl_medical_record.status,symptom,appointmentId from tbl_medical_record " +
		"join appointment on appointmentId=appointment.id " +
		"join tbl_user on tbl_user.user_key=appointment.patient_key " +
		"join tbl_doctor on appointment.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_medical_institution_department on appointment.department_id=tbl_medical_institution_department.department_id " +
		"join tbl_medical_institution on appointment.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"where appointment.doctor_key='" + doctorKey + "' and tbl_user.name='"+name+"' and tbl_medical_record.status=0"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		item := MedicalRecord{}
		err = rows.Scan(&item.ID, &item.Time, &item.DeseaseName, &item.Info,&item.DepName,&item.PatientName,&item.PatientBirth,&item.PatientGender,
			&item.DoctorName,&item.HospitalName,&item.PatientKey,&item.DoctorKey,&item.Status,&item.Symptom,&item.AppointmentId)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//找出时间范围内的医疗记录
func GetMedicalRecordByTime(startTime time.Time, endTime time.Time) []MedicalRecord {
	var mr MedicalRecord
	var mrs []MedicalRecord
	stime := startTime.Format("2006-1-2")
	etime := endTime.Format("2006-1-2")
	query := "select * from tbl_medical_record mr where mr.add_time > '" + stime + "' and mr.add_time<'" + etime + "'"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		if err = rows.Scan(&mr.ID, &mr.DeseaseName, &mr.Info, &mr.UserKey, &mr.DoctorKey, &mr.Time); err == nil {
			mrs = append(mrs, mr)
		} else {
			util.CheckErr(err)
		}
	}
	return mrs
}
