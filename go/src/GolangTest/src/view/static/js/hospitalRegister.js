$(document).ready(function () {
    $("#hospitalRegisterForm").bootstrapValidator({
        message: '通用的验证失败消息',
        feedbackIcons:{
            valid:'glyphicon glyphicon-ok',
            invalid:'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            hospitalname: {
                validators: {
                    notEmpty: {
                        message: '医院名不能为空！'
                    },
                    stringLength: {
                        /*长度提示*/
                        min: 1,
                        max: 50,
                        message: '医院名不能超过50个中文字符！'
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
                        message: '确认密码不能为空！'
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
            detailinfo:{
                validators:{
                    notEmpty: {
                        message: '医院简介不能为空！'
                    },
                    stringLength:{
                        max:1000,
                        message:'医院简介不能超过500个中文字符！'
                    }
                }

            },
            address:{
                validators:{
                    notEmpty:{
                        message:'医院地址不能为空！'
                    }
                }
            },
            grade:{
                validators:{
                    notEmpty:{
                        message:'医院等级不能为空！'
                    }
                }
            }
        }
    });

    $("#validateBtn").click(function(){
        $("#hospitalRegisterForm").bootstrapValidator('validate');
    });
});