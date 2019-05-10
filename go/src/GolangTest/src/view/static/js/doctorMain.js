window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    var url=window.location.href;
    // document.getElementById("doctorCenter").href="#";
    document.getElementById("doctorCenter1").innerText=obj.Name;

    if(url.indexOf("main")!=-1) { //如果处在预约看诊的页面
        //获取当前登录的医生的status
        var doctorStatus=document.getElementById("doctorStatus").innerText;
        if(doctorStatus=="1" || doctorStatus=="4")
        {
            //显示开始工作按钮
            var btn1=document.getElementById("startBtn");
            var btn2=document.getElementById("SuspendBtn");
            btn1.style.display="inline";
            btn2.style.display="none";
        }
        else if(doctorStatus=="2" || doctorStatus=="3"){
            var btn1=document.getElementById("startBtn");
            var btn2=document.getElementById("SuspendBtn");
            btn1.style.display="none";
            btn2.style.display="inline";
            //显示表格
            // var tbody=document.getElementById("tbody");
            var appointTxt=document.getElementById("appoints").innerText;
            if(appointTxt){
                var appoints = JSON.parse(appointTxt).Items;
                for (var i in appoints) {
                    var tr = document.createElement("tr");
                    var td1 = document.createElement("td");
                    td1.innerText=appoints[i].Number;
                    tr.appendChild(td1);
                    var td2 = document.createElement("td");
                    td2.innerText = appoints[i].PatientName;
                    tr.appendChild(td2);
                    var td31 = document.createElement("td");
                    td31.innerText = appoints[i].AppointDate;
                    tr.appendChild(td31);
                    var td3 = document.createElement("td");
                    td3.innerText = appoints[i].Time;
                    tr.appendChild(td3);
                    var td4= document.createElement("td");
                    if(appoints[i].Status=="0")
                        td4.innerText ="等待就诊";
                    tr.appendChild(td4);
                    var a = document.createElement("a");
                    a.innerText = "看诊";
                    a.setAttribute("class", "btn btn-primary");
                    a.setAttribute("id", "consult");
                    a.setAttribute("href", "/doctor/addCase/"+appoints[i].AppointmentId);
                    tr.appendChild(a);
                    tbody.appendChild(tr);
                }
            }
        }
    }

    if(url.indexOf("treatmentHistory")!=-1) { //如果处在诊治历史页面
        var historyTxt=document.getElementById("history").innerText;
        if(historyTxt){
            var historys = JSON.parse(historyTxt).Items;
            for (var i in historys) {
                var tr = document.createElement("tr");
                var td=document.createElement("td");
                var check=document.createElement("input");
                check.setAttribute("type","checkbox");
                check.setAttribute("name","check");
                check.setAttribute("id","check"+i);
                check.setAttribute("type","checkbox");
                check.setAttribute("value",historys[i].ID);
                td.appendChild(check);
                tr.appendChild(td);
                var td1 = document.createElement("td");
                td1.innerText=historys[i].PatientName;
                tr.appendChild(td1);
                var age = document.createElement("td");
                age.innerText=jsGetAge(historys[i].PatientBirth);
                tr.appendChild(age);
                var td2 = document.createElement("td");
                td2.innerText = historys[i].DeseaseName;
                tr.appendChild(td2);
                var td3 = document.createElement("td");
                td3.innerText = historys[i].Time;
                tr.appendChild(td3);
                tbody.appendChild(tr);
            }
        }
    }

    if(obj.Role==1)//如果身份是科室管理员
    {
        if(document.getElementById("userfunction")) {
            var ul = document.getElementById("userfunction");
            var li = document.createElement("li");
            li.className = "list-group-item";
            var a = document.createElement("a");
            a.href = "/doctor/departmentManagement";
            a.innerText = "科室管理";
            li.appendChild(a);
            ul.appendChild(li);
        }
        if(document.getElementById("doctorMgrUL")){
            var li1=document.getElementById("li1");
            var li2=document.getElementById("li2");
            var li3=document.getElementById("li3");
            if(url.indexOf("departmentManagement")!=-1)
                li1.className="active";
            else if(url.indexOf("viewArrangement")!=-1)
                li2.className="active";
            else if(url.indexOf("setAppointmentNum")!=-1)
                li3.className="active";
        }
    }
}

function modal(obj) {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    var recordTxt=document.getElementById("history").innerText;
    var records=JSON.parse(recordTxt).Items;
    document.getElementById("modal_curedoctor").innerText = obj.Name;

    //获取已勾选记录的id
    var id = "";
    $("input[name='check']:checkbox:checked").each(function () {
        id = $(this).val();
    })

    for (var i in records) {//根据id定位到这条记录
        if (id == records[i].ID) {
            document.getElementById("modal_illname").innerText = records[i].DeseaseName;
            document.getElementById("modal_symptom").innerText = records[i].Symptom;
            document.getElementById("modal_illdetail").innerText = records[i].Info;
            document.getElementById("modal_department").innerText = records[i].DepName;
            document.getElementById("modal_hospital").innerText = records[i].HospitalName;
            document.getElementById("modal_name").innerText = records[i].PatientName;
            document.getElementById("modal_time").innerText = records[i].Time;
            //根据出生日期计算周岁年龄
            var birth = records[i].PatientBirth;
            var age = jsGetAge(birth);
            document.getElementById("modal_age").innerText = age;
            if (records[i].PatientGender == 0)
                document.getElementById("modal_gender").innerText = "男";
            else
                document.getElementById("modal_gender").innerText = "女";
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