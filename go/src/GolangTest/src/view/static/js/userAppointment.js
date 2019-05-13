window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("userCenter").href="/user/editInfo";
    document.getElementById("userCenter1").innerText=obj.Name;
    setHospital();
    setDepartment();
    setDate();
    setUrl();
    passValue();
    setSelectAfterPost();
    setHiddenField();
}

$("#hospitalName").change(function(){
    window.location.href='/user/appointment/'+this.value;
});

$("#deptName").change(function(){
    var p1=document.getElementById('hospitalName').value;
    window.location.href='/user/appointment/'+p1+'&'+this.value;
});

// $("#selectDate").change(function(){
//     var p1=document.getElementById('hospitalName').value;
//     var p2=document.getElementById('deptName').value;
//     window.location.href='/user/appointment/'+p1+'&'+p2+'&'+this.value;
// });

//根据url路径设置下拉框的值
function setUrl() {
    var url=window.location.href.split("/");
    if(url.length>1) {
        var str=url[url.length-1].split("&");
        var select = document.getElementById("hospitalName");//设置医院
        var options = select.children;
        if(str.length >0) {
            for (var i in options) {
                if (options[i].value == str[0])
                    options[i].selected = true;
            }
            if (str.length > 1) {
                var dept = document.getElementById("deptName");
                options = dept.children;
                for (var i in options) {
                    if (options[i].value == str[1])
                        options[i].selected = true;
                }
                if (str.length > 2) {
                    var date = document.getElementById("selectDate");
                    options = date.children;
                    for (var i in options) {
                        if (options[i].value == str[2])
                            options[i].selected = true;
                    }
                }
            }
        }
    }
}

//设置日期
function setDate() {
    var mydate=new Date();//读取当前日期
    var weekday=["周日","周一","周二","周三","周四","周五","周六"];

    var selectDate=document.getElementById("selectDate");
    // var op2=document.createElement("option");
    // op2.innerText="不限";
    // op2.value="-1";
    // selectDate.appendChild(op2);

    var op1=document.createElement("option");
    op1.value=mydate.getFullYear()+"-"+parseInt(mydate.getMonth()+1)+"-"+mydate.getDate();
    op1.innerText=parseInt(mydate.getMonth()+1) +"月"+mydate.getDate() +"日("+weekday[mydate.getDay()]+")";
    selectDate.appendChild(op1);

    for(var i=1;i<7;i++){
        var op=document.createElement("option");
        mydate.setDate(mydate.getDate()+1);
        op.value=mydate.getFullYear()+"-"+parseInt(mydate.getMonth()+1)+"-"+mydate.getDate();
        op.innerText=parseInt(mydate.getMonth()+1) +"月"+mydate.getDate() +"日("+weekday[mydate.getDay()]+")";
        selectDate.appendChild(op);
    }
}

//读取医院信息
function setHospital() {
    var hospitalsTxt = document.getElementById("hospitals").innerText;
    if (hospitalsTxt) {
        var hospitals = JSON.parse(hospitalsTxt).Items;
        var selectHospital=document.getElementById("hospitalName");
        for (var i in hospitals) {
            var op = document.createElement("option");
            op.innerText = hospitals[i].Name;
            op.value=hospitals[i].HospitalId;
            selectHospital.appendChild(op);
        }
    }
}

//根据选中的医院读取科室
function setDepartment() {
    var departmentsTxt = document.getElementById("departments").innerText;
    if (departmentsTxt) {
        var departments = JSON.parse(departmentsTxt).Items;
        var selectDept=document.getElementById("deptName");
        for (var i in departments) {
            var op = document.createElement("option");
            op.innerText = departments[i].DeptName;
            op.value=departments[i].DeptId;
            selectDept.appendChild(op);
        }
    }
}

//根据选中的医院和科室读取医生
function setDoctor() {
    var doctorsTxt = document.getElementById("doctors").innerText;
    if (doctorsTxt) {
        var doctors = JSON.parse(doctorsTxt).Items;
        var selectDoctor=document.getElementById("doctorName");
        for (var i in doctors) {
            var op = document.createElement("option");
            op.innerText = doctors[i].Name;
            op.value=doctors[i].DoctorKey;
            selectDoctor.appendChild(op);
        }
    }
}

//post提交表单，搜索符合条件的记录
function search() {
    //判断是否全部都有选择
    var hospical=document.getElementById("hospitalName").value;
    var dept=document.getElementById("deptName").value;
    var day=document.getElementById("selectDate").value;
    // var doctor=document.getElementById("doctorName").value;
    if(hospical=="null" || dept=="null" ||day=="null")
        alert("请选择！");
    else {
        var appointForm = document.getElementById("appointForm");
        appointForm.action = '/user/appointment';
        appointForm.method = 'post';
        appointForm.submit();
    }
}

function passValue() {
    var tbody=document.getElementById("tbody");
    var remainField=document.getElementById("remainField");
    var remain=document.getElementById("remain").innerText;
    var day = document.getElementById("day").innerText;
    if(remain && day) {
        var h4=document.createElement("h4");
        h4.setAttribute("style","float:left;margin-left:10px;");
        h4.innerText="日期："+day;
        remainField.appendChild(h4);

        var h4_1=document.createElement("h4");
        h4_1.setAttribute("style","float:left;margin-left:50px;");
        h4_1.innerText="剩余号数："+remain;
        remainField.appendChild(h4_1);

        if (remain == "0") {
            var tr = document.createElement("tr");
            tbody.appendChild(tr);
            var td = document.createElement("td");
            td.innerText = "暂无余号！";
            tr.appendChild(td);
        }
        else {
            var doctorsTxt = document.getElementById("doctors").innerText;
            if (doctorsTxt) {
                var doctors = JSON.parse(doctorsTxt).Items;
                for (var i in doctors) {
                    var tr = document.createElement("tr");
                    tbody.appendChild(tr);
                    var td = document.createElement("input");
                    td.value = doctors[i].DoctorKey;
                    tr.appendChild(td);
                    td.setAttribute("hidden","hidden");
                    td.setAttribute("id","doctorKey");
                    td.setAttribute("name","doctorKey");
                    var td1 = document.createElement("td");
                    td1.innerText = doctors[i].Name;
                    tr.appendChild(td1);
                    var td2 = document.createElement("td");
                    var btn = document.createElement("button");
                    btn.innerText = "预约";
                    btn.setAttribute("class", "btn btn-primary");
                    btn.setAttribute("type", "submit");
                    td2.appendChild(btn)
                    tr.appendChild(td2);
                }
            }
        }
    }
}

function setSelectAfterPost() {
    var hospitalId=document.getElementById("hospitalId").innerText;
    var deptId=document.getElementById("deptId").innerText;
    if(hospitalId && deptId){ //如果三者均不为空
        var select = document.getElementById("hospitalName");//设置医院
        var options = select.children;
        for (var i in options) {
            if (options[i].value == hospitalId)
                options[i].selected = true;
        }
        var dept = document.getElementById("deptName");
        options = dept.children;
        for (var i in options) {
            if (options[i].value == deptId)
                options[i].selected = true;
        }
    }
}

function setHiddenField() {
    var span1=document.getElementById("hospitalId").innerText;
    var span2=document.getElementById("deptId").innerText;
    var span3=document.getElementById("day").innerText;
    var hospitalID=document.getElementById("hospitalID");
    hospitalID.value=span1;
    var deptID=document.getElementById("deptID");
    deptID.value=span2;
    var DAY=document.getElementById("DAY");
    DAY.value=span3;
}