package model

import "util"

type Doctor struct {
	DoctorKey string 	//公钥
	Name      string 	//姓名
	Gender    int   	//性别
	BirthDate string 	//出生日期
	IdNum     string 	//身份证号
	PhoneNum  string 	//电话号码
	Password  string    //密码
	HospitalId  int	//就职医院的代码
	HospitalName string
	DeptId		 int	//科室代码
	DeptName     string
	Title     int    	//职称(0：初级职称；1：中级职称；2：副高级职称；3：高级职称)
	Status    int 		//状态（-1：审核不通过；0：待审核；1：离线；2：空闲；3：忙碌；4：挂起）
	Role      int       //角色（0：普通医生；1：管理员）
	Aec_Key   string    //信息加密的对称加密秘钥
	Addr 	  string    //记录的地址
}

type Doctors struct {
	Items    []*Doctor
}

func DoctorRegister(doctor Doctor) (bool, error){
	sql :="insert into tbl_doctor(doctor_key,Name,birthdate,gender,id_number,phone_number,medical_institution_id,department_id," +
		"password,aec_key,addr,title,status,role) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql,doctor.DoctorKey,doctor.Name,doctor.BirthDate,doctor.Gender,doctor.IdNum,doctor.PhoneNum,doctor.HospitalId,doctor.DeptId,doctor.Password,doctor.Aec_Key,doctor.Addr,doctor.Title,doctor.Status,doctor.Role)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	util.CheckErr(err)
	return result > 0, nil
}

//判断医院代码和科室代码是否存在
func CheckDoctorIDs(doctor Doctor) (int,error){
	sql :="select ifnull(count(*),0) as count from tbl_doctor where medical_institution_id=? and department_id=?"
	var count int
	err := db.QueryRow(sql,doctor.HospitalId,doctor.DeptId).Scan(&count)
	util.CheckErr(err)
	return count, err
}

//判断手机号是否注册过
func CheckDoctorPhone(doctor Doctor) (int, error) {
	sql :="select ifnull(count(*),0) as count from tbl_doctor where phone_number=?"
	var count int
	err := db.QueryRow(sql,doctor.PhoneNum).Scan(&count)
	util.CheckErr(err)
	return count, err
}

//判断数据库中是否存在该用户
func CheckDoctorLogin(doctor Doctor) (int, error) {
	sql :="select ifnull(count(*),0) as count from tbl_doctor where (name=? or phone_number=?) and password=?"
	var count int
	err := db.QueryRow(sql,doctor.Name,doctor.Name,doctor.Password).Scan(&count)
	util.CheckErr(err)
	return count, err
}

func DoctorLogin(doctor Doctor)(*Doctor, error){
	sql :="select doctor_key,tbl_doctor.Name,birthdate,gender,id_number,phone_number,tbl_doctor.medical_institution_id,tbl_medical_institution.name," +
		"tbl_doctor.department_id,tbl_medical_institution_department.department_name,tbl_doctor.password,aec_key,tbl_doctor.addr,title,tbl_doctor.status,role from tbl_doctor " +
		"join tbl_medical_institution on tbl_doctor.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"join tbl_medical_institution_department on tbl_doctor.department_id=tbl_medical_institution_department.department_id " +
		"where (tbl_doctor.name=? or phone_number=?) and tbl_doctor.password=?"
	err := db.QueryRow(sql,doctor.Name,doctor.Name,doctor.Password).Scan(&doctor.DoctorKey,&doctor.Name,&doctor.BirthDate,&doctor.Gender,&doctor.IdNum,&doctor.PhoneNum,&doctor.HospitalId,&doctor.HospitalName,&doctor.DeptId,&doctor.DeptName,&doctor.Password,&doctor.Aec_Key,&doctor.Addr,&doctor.Title,&doctor.Status,&doctor.Role)
	util.CheckErr(err)
	return &doctor, err
}

//修改个人信息
func UpdateDoctotInfo(doctor Doctor)(int64, error){
	sql :="update tbl_doctor set phone_number=? where doctor_key=?"
	res, err := db.Exec(sql,doctor.PhoneNum,doctor.DoctorKey)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}

//修改密码
func UpdateDoctorPwd(doctor Doctor)(int64, error){
	sql :="update tbl_doctor set password=? where doctor_key=?"
	res, err := db.Exec(sql,doctor.Password,doctor.DoctorKey)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}

func GetDoctorInfoByPublicKey(doctor Doctor)(*Doctor, error){
	sql :="select doctor_key,tbl_doctor.Name,birthdate,gender,id_number,phone_number,tbl_doctor.medical_institution_id,tbl_medical_institution.name," +
		"tbl_doctor.department_id,tbl_medical_institution_department.department_name,password,aec_key,addr,title,status,role from tbl_doctor " +
		"join tbl_medical_institution on tbl_doctor.medical_institution_id=tbl_medical_institution.medical_institution_id " +
		"join tbl_medical_institution_department on tbl_doctor.department_id=tbl_medical_institution_department.department_id " +
		"where doctor_key=?"
		err := db.QueryRow(sql,doctor.DoctorKey).Scan(&doctor.DoctorKey,&doctor.Name,&doctor.BirthDate,&doctor.Gender,&doctor.IdNum,&doctor.PhoneNum,&doctor.HospitalId,&doctor.HospitalName,&doctor.DeptId,&doctor.DeptName,&doctor.Password,&doctor.Aec_Key,&doctor.Addr,&doctor.Title,&doctor.Status,&doctor.Role)
		util.CheckErr(err)
		return &doctor, err
}

//读取数据库中所有待审核的医生
func GetUnverifiedDoctors(hospitalId int) *Doctors {
	var result Doctors
	result.Items = []*Doctor{}
	query := "select doctor_key,name,birthdate,gender,id_number,phone_number,department_name,title,aec_key,status from tbl_doctor " +
		"join tbl_medical_institution_department on tbl_medical_institution_department.department_id=tbl_doctor.department_id " +
		"where tbl_doctor.medical_institution_id=? and (status='0' or status='-1')"
	rows, err := db.Query(query,hospitalId)
	util.CheckErr(err)
	for rows.Next() {
		item := Doctor{}
		err = rows.Scan(&item.DoctorKey,&item.Name,&item.BirthDate,&item.Gender,&item.IdNum,&item.PhoneNum,&item.DeptName,&item.Title,&item.Aec_Key,&item.Status)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//读取数据库中所有在职医生
func GetValidDoctors(hospitalId int) *Doctors {
	var result Doctors
	result.Items = []*Doctor{}
	query := "select doctor_key,name,birthdate,gender,id_number,phone_number,department_name,title,aec_key,status,role from tbl_doctor " +
		"join tbl_medical_institution_department on tbl_medical_institution_department.department_id=tbl_doctor.department_id " +
		"where tbl_doctor.medical_institution_id=? and (status=1 or status=2 or status=3 or status=4)"
	rows, err := db.Query(query,hospitalId)
	util.CheckErr(err)
	for rows.Next() {
		item := Doctor{}
		err = rows.Scan(&item.DoctorKey,&item.Name,&item.BirthDate,&item.Gender,&item.IdNum,&item.PhoneNum,&item.DeptName,&item.Title,&item.Aec_Key,&item.Status,&item.Role)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//根据下拉框选中的参数找到相应的医生
func GetSelectedDoctors(hospitalId int,param string,name string) *Doctors {
	var result Doctors
	result.Items = []*Doctor{}
	query := "select doctor_key,name,birthdate,gender,id_number,phone_number,department_name,title,aec_key,status,role from tbl_doctor " +
		"join tbl_medical_institution_department on tbl_medical_institution_department.department_id=tbl_doctor.department_id " +
		"where tbl_doctor.medical_institution_id=? and (status=1 or status=2 or status=3 or status=4) and "+param+"=?"
	rows, err := db.Query(query,hospitalId,name)
	util.CheckErr(err)
	for rows.Next() {
		item := Doctor{}
		err = rows.Scan(&item.DoctorKey,&item.Name,&item.BirthDate,&item.Gender,&item.IdNum,&item.PhoneNum,&item.DeptName,&item.Title,&item.Aec_Key,&item.Status,&item.Role)
		util.CheckErr(err)
		result.Items = append(result.Items, &item)
	}
	return &result
}

//审核医生注册申请
func UpdateDoctorStatus(id string,status int) (int64,error){
	sql :="update tbl_doctor set status=? where doctor_key=?"
	res, err := db.Exec(sql,status,id)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}

//设置医生为科室管理员
func UpdateDoctorRole(id string,role int) (int64,error){
	sql :="update tbl_doctor set role=? where doctor_key=?"
	res, err := db.Exec(sql,role,id)
	util.CheckErr(err)
	result, err := res.RowsAffected()
	return result, nil
}