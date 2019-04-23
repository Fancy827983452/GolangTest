package model

import "util"

type Admin struct {
	ID        string
	Password  string
}

func AdminLogin(a Admin) (*Admin, error){
	sql :="select Id,password from admin where id=? and password=?"
	err := db.QueryRow(sql,a.ID,a.Password).Scan(&a.ID,&a.Password)
	util.CheckErr(err)
	return &a, err
}

//判断数据库中是否存在该账号
func CheckAdminLogin(a Admin) (int, error) {
	sql :="select ifnull(count(*),0) as count from admin where id=? and password=?"
	var count int
	err := db.QueryRow(sql,a.ID,a.Password).Scan(&count)
	util.CheckErr(err)
	return count, err
}