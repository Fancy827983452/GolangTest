window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt)
    document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("userCenter").href="/user/editInfo";
    document.getElementById("userCenter1").innerText=obj.Name;
}

$(document).ready(function () {
    $("#editUserPwdForm").bootstrapValidator({
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
        $("#editUserPwdForm").bootstrapValidator('validate');
    });
});