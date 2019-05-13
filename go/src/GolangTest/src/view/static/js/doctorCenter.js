window.onload=function () {
    var jsonTxt = document.getElementById("formValue").innerText;
    var obj = JSON.parse(jsonTxt);
    var url = window.location.href;
    document.getElementById("doctorCenter1").innerText = obj.Name;

    //如果处在修改信息页面
    if (url.indexOf("editInfo") != -1) {
        document.getElementById("doctorKey").value = obj.DoctorKey;
        document.getElementById("username").value = obj.Name;
        if (obj.Gender == 0)
            document.getElementById("sex1").checked = true;
        else
            document.getElementById("sex2").checked = true;
        document.getElementById("birthdate").value = obj.BirthDate;
        document.getElementById("idnum").value = obj.IdNum;
        document.getElementById("hospitalName").value = obj.HospitalName;
        document.getElementById("DepName").value = obj.DeptName;
        document.getElementById("tel").value = obj.PhoneNum;
    }

    //如果处在修改密码页面
    if (url.indexOf("editPwd") != -1) {
        document.getElementById("doctorKey").value = obj.DoctorKey;
    }

    //如果处在添加病历页面
    if(url.indexOf("addCase")!=-1){
        document.getElementById("doctorKey").value=obj.DoctorKey;
        document.getElementById("appoint").value=document.getElementById("appointmentId").innerText;
        //读取患者信息
        var patientTxt=document.getElementById("patient").innerText;
        var patient=JSON.parse(patientTxt);
        document.getElementById("name").innerText=patient.Name;
        //根据出生日期计算周岁年龄
        document.getElementById("age").innerText=jsGetAge(patient.BirthDate);
        if(patient.Gender==0)
            document.getElementById("gender").innerText="男";
        else
            document.getElementById("gender").innerText="女";
        //读取主治医生
        document.getElementById("curedoctor").value=obj.Name;
    }

    //如果身份是科室管理员
    if(obj.Role==1) {
        if (document.getElementById("userfunction")) { //左侧菜单
            var ul = document.getElementById("userfunction");
            var li = document.createElement("li");
            li.className = "list-group-item";
            var a = document.createElement("a");
            a.href = "/doctor/departmentManagement";
            a.innerText = "科室管理";
            li.appendChild(a);
            ul.appendChild(li);
        }
        if (document.getElementById("doctorMgrUL")) {  //右侧上方导航菜单
            var li1 = document.getElementById("li1");
            var li2 = document.getElementById("li2");
            var li3 = document.getElementById("li3");
            if (url.indexOf("departmentManagement") != -1) { //处于排班安排页面
                li1.className = "active";
                //显示本科室所有在职医生
                var doctorsTxt = document.getElementById("doctors").innerText;
                if (doctorsTxt) {
                    var doctors = JSON.parse(doctorsTxt).Items;
                    document.getElementById("deptName").innerText = doctors[0].DeptName;
                    var tbody = document.getElementById("tbody");
                    for (var i in doctors) {
                        var tr = document.createElement("tr");
                        tbody.appendChild(tr);
                        var td = document.createElement("td");
                        var check = document.createElement("input");
                        check.setAttribute("type", "checkbox");
                        check.setAttribute("name", "check");
                        check.setAttribute("id", "check" + i);
                        check.setAttribute("type", "checkbox");
                        check.setAttribute("value", doctors[i].DoctorKey);
                        td.appendChild(check);
                        tr.appendChild(td);
                        var td1 = document.createElement("td");
                        td1.innerText = doctors[i].Name;
                        tr.appendChild(td1);
                        var td2 = document.createElement("td");
                        if (doctors[i].Gender == 0)
                            td2.innerText = "男";
                        else
                            td2.innerText = "女";
                        tr.appendChild(td2);
                        var td3 = document.createElement("td");
                        td3.innerText = jsGetAge(doctors[i].BirthDate);
                        tr.appendChild(td3);
                        var td4 = document.createElement("td");
                        switch (doctors[i].Title) {
                            case 0:
                                td4.innerText = "医士、医师、住院医师";
                                break;
                            case 1:
                                td4.innerText = "主治医师";
                                break;
                            case 2:
                                td4.innerText = "副主任医师";
                                break;
                            case 3:
                                td4.innerText = "主任医师";
                                break;
                        }
                        tr.appendChild(td4);
                        var td5 = document.createElement("td");
                        if (doctors[i].Role == 1)
                            td5.innerText = "管理员";
                        tr.appendChild(td5);
                        var td6 = document.createElement("td");
                        switch (doctors[i].Arrange) {
                            case 0:
                                td6.innerText = "星期日";
                                break;
                            case 1:
                                td6.innerText = "星期一";
                                break;
                            case 2:
                                td6.innerText = "星期二";
                                break;
                            case 3:
                                td6.innerText = "星期三";
                                break;
                            case 4:
                                td6.innerText = "星期四";
                                break;
                            case 5:
                                td6.innerText = "星期五";
                                break;
                            case 6:
                                td6.innerText = "星期六";
                                break;
                            default:
                                td6.innerText = "未安排";
                                break;
                        }
                        tr.appendChild(td6);
                    }
                }
            }
            else if (url.indexOf("viewArrangement") != -1) {  //查看排班表页面
                li2.className = "active";
                //显示本科室所有在职医生
                var doctorsTxt = document.getElementById("doctors").innerText;
                if (doctorsTxt) {
                    var doctors = JSON.parse(doctorsTxt).Items;
                    var tbody = document.getElementById("tbody");
                    for (var i in doctors) {
                        //判断是周几
                        var td1 = document.getElementById("td1");
                        var td2 = document.getElementById("td2");
                        var td3 = document.getElementById("td3");
                        var td4 = document.getElementById("td4");
                        var td5 = document.getElementById("td5");
                        var td6 = document.getElementById("td6");
                        var td7 = document.getElementById("td7");
                        var div = document.createElement("div");
                        div.innerText = doctors[i].Name;
                        div.style.height = "30px";
                        div.style.marginTop = "5%";
                        // div.style.border= "thin solid lightgray";
                        switch (doctors[i].Arrange) {
                            case 1:
                                td1.appendChild(div);
                                break;
                            case 2:
                                td2.appendChild(div);
                                break;
                            case 3:
                                td3.appendChild(div);
                                break;
                            case 4:
                                td4.appendChild(div);
                                break;
                            case 5:
                                td5.appendChild(div);
                                break;
                            case 6:
                                td6.appendChild(div);
                                break;
                            case 0:
                                td7.appendChild(div);
                                break;
                        }
                    }
                }
            }
            else if (url.indexOf("setAppointmentNum") != -1) {  //设置挂号数量页面
                li3.className="active";
            }
        }
    }
}

// 点击行选中该行复选框
$("#table").on("click", "tr", function () {
    var input = $(this).find("input");
    if (!$(input).prop("checked")) {
        $(input).prop("checked", true);
    } else {
        $(input).prop("checked", false);
    }
});

// 多选框 防止事件冒泡
$("#table").on("click", "input", function (event) {
    event.stopImmediatePropagation();
});

// 点击后全选所有复选框
$("input[name='all_check']").change(function () {
    if (this.checked) {
        $("input[name='check']:checkbox").each(function () {
            this.checked = true;
        })
    } else {
        $("input[name='check']:checkbox").each(function () {
            this.checked = false;
        })
    }
});

//判断是否选择记录
function checkSelected() {

    var check=document.getElementsByName("check");
    var flag=false;
    var count=0;
    for(var c=0;c<check.length;c++) {
        if (check[c].checked == true) {
            count++;
        }
    }
    if(count>0)
        flag = true;
    else {
        alert("请至少选择一条记录！");
        window.history.back(-1);
    }
    return true;
}

//bootstrap验证器
$(document).ready(function () {
    $("#doctorEditInfoForm").bootstrapValidator({
        message: '通用的验证失败消息',
        feedbackIcons:{
            valid:'glyphicon glyphicon-ok',
            invalid:'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            tel:{
                validators:{
                    notEmpty: {
                        message: '手机号码不能为空！'
                    },
                    stringLength:{
                        min:11,
                        max:11,
                        message:'手机号码必须为11为数字'
                    },
                    regexp: {
                        regexp: /^1[3|4|5|7|8]{1}[0-9]{9}$/,
                        message: '请输入正确的手机号码'
                    }
                }

            }
        }
    });

    $("#editDoctorPwdForm").bootstrapValidator({
        message: '通用的验证失败消息',
        feedbackIcons:{
            valid:'glyphicon glyphicon-ok',
            invalid:'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            password_old: {
                validators: {
                    notEmpty: {
                        message: '原密码不能为空！'
                    },
                    stringLength: {
                        /*长度提示*/
                        min: 6,
                        max: 20,
                        message: '密码不得少于6个字符，不能超过20个字符'
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '密码只能由英文字母、数字、下划线以及小数点组成'
                    }
                }
            },
            password_new: {
                validators: {
                    notEmpty: {
                        message: '新密码不能为空！'
                    },
                    stringLength: {
                        /*长度提示*/
                        min: 6,
                        max: 20,
                        message: '密码不得少于6个字符，不能超过20个字符'
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '密码只能由英文字母、数字、下划线以及小数点组成'
                    }
                }
            },
            password_new2: {
                message: '不合法的密码',
                validators: {
                    notEmpty: {
                        message: '确认密码不能为空！'
                    },
                    stringLength: {
                        min: 6,
                        max: 20,
                        message: '密码不得少于6个字符，不能超过20个字符'
                    },
                    identical: {//相同
                        field: 'password_new',
                        message: '两次密码输入不一致'
                    },
                    regexp: {//匹配规则
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '确认密码只能由英文字母、数字、下划线以及小数点组成'
                    }
                }
            }
        }
    });

    $("#validateBtn").click(function(){
        $("#doctorEditInfoForm").bootstrapValidator('validate');
        $("#editDoctorPwdForm").bootstrapValidator('validate');
    });
});

//根据出生日期计算周岁年龄
function jsGetAge(strBirthday) {
    var returnAge;
    var strBirthdayArr=strBirthday.split("-");
    var birthYear = strBirthdayArr[0];
    var birthMonth = strBirthdayArr[1];
    var birthDay = strBirthdayArr[2];

    d = new Date();
    var nowYear = d.getFullYear();
    var nowMonth = d.getMonth() + 1;
    var nowDay = d.getDate();

    if(nowYear == birthYear)
        returnAge = 0;//同年 则为0岁
    else
    {
        var ageDiff = parseInt(nowYear - birthYear) ; //年之差
        if(ageDiff > 0) {
            if(nowMonth == birthMonth) {
                var dayDiff = nowDay - birthDay;//日之差
                if(dayDiff < 0)
                    returnAge = ageDiff - 1;
                else
                    returnAge = ageDiff ;
            }
            else {
                var monthDiff = nowMonth - birthMonth;//月之差
                if(monthDiff < 0)
                    returnAge = ageDiff - 1;
                else
                    returnAge = ageDiff ;
            }
        }
        else
            returnAge = -1;//返回-1 表示出生日期输入错误 晚于今天
    }
    return returnAge;//返回周岁年龄
}

function check() {
    var illdescribe=document.createElement("illdescribe").value;
    var illname=document.getElementById("illname").value;
    var illdetail=document.getElementById("illdetail").value;
    if(illdescribe.length==0 || illname.length==0 || illdetail.length==0){
        alert("请填写完整！");
        return false;
    }
    else
        return true;
}

//安排排班
function setArrange() {
    var str="";
    $("input[name='check']:checkbox:checked").each(function(){
        str+=$(this).val()+",";
    })
    $("#selectedItem").val(str);
    document.getElementById("selectedItem").value=str;
    var check=document.getElementsByName("check");
    for(var c=0;c<check.length;c++){
        if(check[c].checked==true){
            var arrangeForm=document.getElementById("arrangeForm");
            arrangeForm.action="/doctor/setArrangement";
            arrangeForm.method="post";
            arrangeForm.submit();
        }
    }
}