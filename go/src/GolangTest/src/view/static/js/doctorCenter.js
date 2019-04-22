window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    var url=window.location.href;
    // document.getElementById("doctorCenter").href="#";
    document.getElementById("doctorCenter1").innerText=obj.Name;

    if(url.indexOf("editInfo")!=-1){ //如果处在修改信息页面
        document.getElementById("doctorKey").value=obj.DoctorKey;
        document.getElementById("username").value=obj.Name;
        if(obj.Gender==0)
            document.getElementById("sex1").checked=true;
        else
            document.getElementById("sex2").checked=true;
        document.getElementById("birthdate").value=obj.BirthDate;
        document.getElementById("idnum").value=obj.IdNum;
        document.getElementById("hospitalName").value=obj.HospitalName;
        document.getElementById("DepName").value=obj.DeptName;
        document.getElementById("tel").value=obj.PhoneNum;
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

    $("#validateBtn").click(function(){
        $("#doctorEditInfoForm").bootstrapValidator('validate');
    });
});