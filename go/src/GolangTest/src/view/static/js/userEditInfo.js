$(function(){
    var date_now = new Date();
    var year = date_now.getFullYear();
    var month = date_now.getMonth() + 1 < 10 ? "0" + (date_now.getMonth() + 1) : (date_now.getMonth() + 1);
    var date = date_now.getDate() < 10 ? "0" + date_now.getDate() : date_now.getDate();
    $("#birthdate").attr("max", year + "-" + month + "-" + date);
})

window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt)
    document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("username").value=obj.Name;
    document.getElementById("userCenter").href="/user/editInfo";
    document.getElementById("userCenter1").innerText=obj.Name;
    if(obj.Gender==0)
        document.getElementById("sex1").checked=true;
    else
        document.getElementById("sex2").checked=true;
    document.getElementById("birthdate").value=obj.BirthDate.substr(0,10);
    document.getElementById("idnum").value=obj.IdNum;
    document.getElementById("tel").value=obj.PhoneNum;
    document.getElementById("location").value=obj.Location;
}

$(document).ready(function () {
    $("#userEditInfoForm").bootstrapValidator({
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

            },
            birthdate:{
                validators:{
                    notEmpty:{
                        message:'出生日期不能为空'
                    }
                }
            },
            location:{
                validators:{
                    notEmpty:{
                        message:'家庭住址不能为空'
                    }
                },
                stringLength:{
                    max:200,
                    message:'家庭住址不能超过200个字符（或100个中文字符）'
                }
            }
        }
    });

    $("#validateBtn").click(function(){
        $("#userEditInfoForm").bootstrapValidator('validate');
    });
});