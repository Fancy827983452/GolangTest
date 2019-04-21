// $(function(){
//     var date_now = new Date();
//     var year = date_now.getFullYear();
//     var month = date_now.getMonth()+1 < 10 ? "0"+(date_now.getMonth()+1) : (date_now.getMonth()+1);
//     var date = date_now.getDate() < 10 ? "0"+date_now.getDate() : date_now.getDate();
//     $("#birthdate").attr("max",year+"-"+month+"-"+date);
// })

$(document).ready(function () {
    $("#doctorRegisterForm").bootstrapValidator({
        message: '通用的验证失败消息',
        feedbackIcons:{
            valid:'glyphicon glyphicon-ok',
            invalid:'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            doctorname: {
                validators: {
                    notEmpty: {
                        message: '姓名不能为空！'
                    },
                    stringLength: {
                        /*长度提示*/
                        min: 1,
                        max: 50,
                        message: '姓名不能超过50个字符（或25个中文字符）'
                    }
                }
            },
            password: {
                validators: {
                    notEmpty: {
                        message: '密码不能为空！'
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
            confirmPassword: {
                message: '不合法的密码',
                validators: {
                    notEmpty: {
                        message: '密码不能为空！'
                    },
                    stringLength: {
                        min: 6,
                        max: 20,
                        message: '密码不得少于6个字符，不能超过20个字符'
                    },
                    identical: {//相同
                        field: 'password',
                        message: '两次密码输入不一致'
                    },
                    regexp: {//匹配规则
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '确认密码只能由英文字母、数字、下划线以及小数点组成'
                    }
                }
            },
            telephone:{
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
            gender:{
                validators:{
                    notEmpty:{
                        message:'性别不能为空'
                    }
                }
            },
            birthdate:{
                validators:{
                    notEmpty:{
                        message:'出生日期不能为空'
                    },
                    stringLength:{
                        min:10,
                        max:10,
                        message:'请参照如下格式填写出生日期（1999-01-01）'
                    },
                    regexp: {
                        regexp: /^((19[0-9]\d{1})|(20((0[0-9])|(1[0-9]))))\-((0?[1-9])|(1[0-2]))\-((0?[1-9])|([1-2][0-9])|30|31)$/,
                        message: '请输入合法的出生日期'
                    }
                }
            },
            idnumber:{
                validators:{
                    notEmpty:{
                        message:'身份证号不能为空'
                    },
                    stringLength:{
                        min:18,
                        max:18,
                        message:'身份证号必须为18位数字'
                    },
                    regexp: {
                        regexp: /^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/,
                        message: '请输入合法的身份证号'
                    }
                }
            },
            hospitalId:{
                validators:{
                    notEmpty:{
                        message:'医院代码不能为空'
                    }
                }
            },
            departmentId:{
                validators:{
                    notEmpty:{
                        message:'科室代码不能为空'
                    }
                }
            },
            title:{
                validators:{
                    notEmpty:{
                        message:'职称不能为空'
                    }
                }
            }
        }
    });

    $("#validateBtn").click(function(){
        $("#doctorRegisterForm").bootstrapValidator('validate');
    });
});