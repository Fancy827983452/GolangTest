package model

import (
	"util"
	"fmt"
)

type User struct {
	PublicKey string 	//用户公钥
	Name      string 	//用户姓名
	Gender    int   	//用户性别
	BirthDate string 	//出生日期
	IdNum     string 	//身份证号
	PhoneNum  string 	//电话号码
	Location  string 	//用户地址
	Account   int  	//账户金额
	Password  string    //用户密码
}

var U User

//注册
func Register(user User) (bool, error) {
	sql :="insert into tbl_user(user_key,Name,birthdate,gender,id_number,phone_number,account,location,Password) values(?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql,user.PublicKey,user.Name,user.BirthDate,user.Gender,user.IdNum,user.PhoneNum,user.Account,user.Location,user.Password)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//登录
func CheckLogin(user User)(int, error){
	sql :="select ifnull(count(*),0) as count from tbl_user where name=? and password=?"
	var count int
	err := db.QueryRow(sql,user.Name,user.Password).Scan(&count)
	util.CheckErr(err)
	return count, err
}

func Login(user User)(*User, error){
	sql :="select user_key,name,birthdate,gender,id_number,phone_number,location,account,password from tbl_user where name=? and password=?"
	err := db.QueryRow(sql,user.Name,user.Password).Scan(&user.PublicKey,&user.Name,&user.BirthDate,&user.Gender,&user.IdNum,&user.PhoneNum,&user.Location,&user.Account,&user.Password)
	util.CheckErr(err)
	return &user, err
}

func GetAllUser() error {
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		if err := rows.Scan(&U.PublicKey, &U.Name, &U.Gender,
			&U.BirthDate, &U.IdNum, &U.PhoneNum, &U.Location); err == nil {
			fmt.Println(U)
		} else {
			fmt.Println(err)
		}
	}
	return err
}
