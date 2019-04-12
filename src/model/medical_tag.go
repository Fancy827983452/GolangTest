package model

import (
	"fmt"
	"util"
)

//根据标签ID列表获取相关的病历ID列表
func GetMedicalIDByTagID(tagIDList []uint32) []uint32 {
	var medicalIDList []uint32
	var id uint32
	for _, tagID := range tagIDList {
		query := "select medical_record_id from medical_tag mt where mt.tag_id = '"+
			 fmt.Sprint(tagID) + "'"
		rows, err := db.Query(query)
		util.CheckErr(err)

		for rows.Next() {
			if err = rows.Scan(&id); err == nil {
				medicalIDList = append(medicalIDList, id)
			} else {
				panic(err)
			}
		}
	}
	return medicalIDList
}
