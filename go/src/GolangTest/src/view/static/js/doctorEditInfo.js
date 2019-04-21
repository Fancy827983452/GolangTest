window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt)
    // document.getElementById("doctorCenter").href="#";
    document.getElementById("doctorCenter1").innerText=obj.Name;
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