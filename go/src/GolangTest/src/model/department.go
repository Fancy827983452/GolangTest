package model

import "util"

type Department struct {
	HospitalId  int		//医院代码
	Name 	    string  //医院名
	DeptId 		int     //科室代码
	DeptName	string  //科室名
	Info  		string
	location	string
	num 		int     //该科室预约挂号的数量
}

type Departments struct {
	Items    []*Department
}

//根据医院id读取所有科室
func GetHospitalDepts(id string) *Departments {
	var result Departments
	result.Items = []*Department{}
	query := "SELECT tbl_medical_institution_department.medical_institution_id,tbl_medical_institution.name," +
		"department_id,department_name,tbl_medical_institution_department.info,ifnull(num,0) " +
		"FROM tbl_medical_institution_department join tbl_medical_institution " +
		"on tbl_medical_institution_department.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"where medical.tbl_medical_institution_department.medical_institution_id=?"
	rows, err := db.Query(query,id)
	util.CheckErr(err)
	for rows.Next() {
		item := Department{}
		err = rows.Scan(&item.HospitalId,&item.Name,&item.DeptId,&item.DeptName,&item.Info,&item.num)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}


func CheckDepartment(department Department)(int, error){
	sql :="select ifnull(count(*),0) as count from tbl_medical_institution_department where department_name=?"
	var count int
	err := db.QueryRow(sql,department.Name).Scan(&count)
	util.CheckErr(err)
	return count, err
}

func DepartmentAdd(department Department) (bool, error) {
	sql :="insert into tbl_medical_institution_department(medical_institution_id,department_name,info) values(?,?,?)"
	res, err := db.Exec(sql,department.HospitalId,department.Name,department.Info)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//更新科室信息
func UpdateDepartmentInfo(department Department)(int64, error){
	sql :="update tbl_medical_institution_department set department_name=?,info=? where department_id=?"
	res, err := db.Exec(sql,department.Name,department.Info,department.DeptId)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}

//根据医院id和科室id，获取该科室设置的挂号量(max)
func GetSettedAppointNum(hospitalId string,deptId string) (int,error) {
	sql:="select num from tbl_medical_institution_department where medical_institution_id=? and department_id=?"
	var num int
	err := db.QueryRow(sql,hospitalId,deptId).Scan(&num)
	util.CheckErr(err)
	return num,err
}