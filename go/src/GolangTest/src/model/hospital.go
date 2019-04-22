package model

import (
	"util"
)

type Hospital struct {
	HospitalId  int		//医院代码
	Name 	    string  //医院名
	Info		string  //医院详细信息
	Location	string  //医院地址
	Grade		string  //医院等级
	Password	string
	Addr		string  //base58地址
	PublicKey   string
	Status      int 	//状态（0：待审核；1：审核通过；2：异常）
}

func HospitalRegister(h Hospital) (bool, error){
	sql :="insert into tbl_medical_institution(Name,Info,Location,Grade,password,addr,publicKey,status) values(?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql,h.Name,h.Info,h.Location,h.Grade,h.Password,h.Addr,h.PublicKey,h.Status)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//判断该医院名是否注册过
func CheckHospitalName(h Hospital) (int, error) {
	sql :="select ifnull(count(*),0) as count from tbl_medical_institution where Name=?"
	var count int
	err := db.QueryRow(sql,h.Name).Scan(&count)
	util.CheckErr(err)
	return count, err
}

func HospitalLogin(h Hospital) (*Hospital, error){
	sql :="select medical_institution_id,Name,Info,Location,Grade,password,addr,publicKey,status from tbl_medical_institution " +
		"where (medical_institution_id=? or Name=?) and password=?"
	err := db.QueryRow(sql,h.Name,h.Name,h.Password).Scan(&h.HospitalId,&h.Name,&h.Info,&h.Location,&h.Grade,&h.Password,&h.Addr,&h.PublicKey,&h.Status)
	util.CheckErr(err)
	return &h, err
}

//判断数据库中是否存在该医院
func CheckHospitalLogin(h Hospital) (int, error) {
	sql :="select ifnull(count(*),0) as count from tbl_medical_institution where (name=? or medical_institution_id=?) and password=?"
	var count int
	err := db.QueryRow(sql,h.Name,h.Name,h.Password).Scan(&count)
	util.CheckErr(err)
	return count, err
}