package mydb

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var DB = InitDataBase()

func InitDataBase() *sql.DB{
	DB, err := sql.Open("mysql",
		"root:960412@tcp(127.0.0.1:3306)/medical?charset=utf8")
	if err != nil {
		panic(err)
	}
	return DB
}
