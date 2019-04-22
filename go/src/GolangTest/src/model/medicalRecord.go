package model

import (
	"fmt"
	"time"
	"util"
)

type MedicalRecord struct {
	ID        	uint32	//medical_record_id
	Time      	string	//添加时间
	Symptom     string  //症状描述
	DeseaseName string	//疾病名称
	Info      	string	//疾病详情
	DepName		string  //就诊科室名
	HospitalName string
	DoctorName 	string  //医生姓名
	UserKey   	string
	DoctorKey 	string
	Status      int
}

type MedicalRecords struct {
	Items    []*MedicalRecord
}

var md MedicalRecord

//添加记录
func AddMedicalRecord(ID uint32, publicKey string, name string, info string, label string,
	time string, doctorKey string) error {

	//todo 分层
	sql := "insert into tbl_medical_record (id, user_key, name, info, " +
		"add_time, doctor_key) values (?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, ID, publicKey, name, info, time, doctorKey)
	if err != nil {
		return err
	}
	query := "select id from tbl_tag where tag.name='" + label + "'"
	fmt.Println(query)
	var labelID uint32
	rows, err := db.Query(query)
	fmt.Println(*rows)
	for rows.Next() {
		if err := rows.Scan(&labelID); err == nil {

		}
	}
	mdclId, _ := res.LastInsertId()
	fmt.Println(mdclId, labelID)
	sql = "insert into medical_tag (id, medical_record_id, tag) values(?,?,?)"
	_, err = db.Exec(sql, 0, mdclId, labelID)
	return err
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
	query := "select medical_record_id,add_time,desease,record_info,department_name,tbl_doctor.name,tbl_medical_institution.name,tbl_medical_record.doctor_key,tbl_medical_record.status,symptom " +
		"from tbl_medical_record " +
		"join tbl_doctor on tbl_medical_record.doctor_key=tbl_doctor.doctor_key " +
		"join tbl_medical_institution_department on tbl_medical_record.department_id=tbl_medical_institution_department.department_id " +
		"join tbl_medical_institution on tbl_medical_record.institution_id=tbl_medical_institution.medical_institution_id " +
		"where userMR_key='" + userKey + "'"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		item := MedicalRecord{}
		err = rows.Scan(&item.ID, &item.Time, &item.DeseaseName, &item.Info,&item.DepName,&item.DoctorName,&item.HospitalName,&item.DoctorKey,&item.Status,&item.Symptom)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

/*
根据ID列表查找相应的病历记录
@return: 病例记录列表
 */
func GetMedicalRecordByLabel(idList []uint32) []MedicalRecord {
	var medicalRecords []MedicalRecord
	for _, id := range idList {
		query := "select * from tbl_medical_record mr where mr.medical_record_id='" + fmt.Sprint(id) + "'"
		rows, err := db.Query(query)
		util.CheckErr(err)
		for rows.Next() {
			var mr MedicalRecord
			if err = rows.Scan(&mr.ID, &mr.DeseaseName, &mr.UserKey, &mr.DoctorKey, &mr.Time, &mr.Info); err == nil {
				medicalRecords = append(medicalRecords, mr)
			} else {
				panic(err)
			}
		}
	}
	return medicalRecords
}

/*
找出时间范围内的医疗记录
 */
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
