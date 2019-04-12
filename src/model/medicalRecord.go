package model

import (
	"fmt"
	"time"
	"util"
)

type MedicalRecord struct {
	ID        uint32
	UserKey   string
	Name      string
	Info      string
	Time      string
	DoctorKey string
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
		if err := rows.Scan(&md.ID, &md.UserKey, &md.Name, &md.Info, &md.Time, &md.DoctorKey); err == nil {
			fmt.Println(md)
		} else {
			return err
		}
	}
	return err
}

/*
根据用户key查询病例记录
 */
func GetMedicalRecordByUser(userKey string) []MedicalRecord {
	var medicalRecords []MedicalRecord
	query := "select * from tbl_medical_record mr where mr.userMR_key='" + userKey + "'"
	rows, err := db.Query(query)
	util.CheckErr(err)
	for rows.Next() {
		var mr MedicalRecord
		if err = rows.Scan(&mr.ID, &mr.Name, &mr.UserKey, &mr.DoctorKey,
			&mr.Time, &mr.Info); err == nil {
			medicalRecords = append(medicalRecords, mr)
		} else {
			panic(err)
		}
	}
	return medicalRecords
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
			if err = rows.Scan(&mr.ID, &mr.Name, &mr.UserKey, &mr.DoctorKey,
				&mr.Time, &mr.Info); err == nil {
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
		if err = rows.Scan(&mr.ID, &mr.Name, &mr.Info, &mr.UserKey, &mr.DoctorKey, &mr.Time); err == nil {
			mrs = append(mrs, mr)
		} else {
			util.CheckErr(err)
		}
	}
	return mrs
}
