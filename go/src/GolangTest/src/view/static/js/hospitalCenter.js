window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("hospitalCenter1").innerText=obj.Name;

    var departmentTxt=document.getElementById("department").innerText;
    var department=JSON.parse(departmentTxt).Items;
    var tbody=document.getElementById("tbody");
    for(var d in department){
        var tr=document.createElement("tr");
        tbody.appendChild(tr);
        var td1=document.createElement("td");
        var check=document.createElement("input");
        check.setAttribute("type","checkbox");
        check.setAttribute("name","check");
        check.setAttribute("id","check"+d);
        check.setAttribute("type","checkbox");
        check.setAttribute("value",department[d].DeptId);
        td1.appendChild(check);
        tr.appendChild(td1);
        // var td2=document.createElement("td");
        // td2.innerText=illcase[i].ID;
        // tr.appendChild(td2);
        var td3=document.createElement("td");
        td3.innerText=department[d].DeptId;
        tr.appendChild(td3);
        var td4=document.createElement("td");
        td4.innerText=department[d].DeptName;
        tr.appendChild(td4);
        var td5=document.createElement("td");
        td5.innerText=department[d].Info;
        tr.appendChild(td5);
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

$(document).ready(function () {
    $("#departmentForm").bootstrapValidator({
        message: '通用的验证失败消息',
        feedbackIcons:{
            valid:'glyphicon glyphicon-ok',
            invalid:'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            DepartmentName: {
                validators: {
                    notEmpty: {
                        message: '科室名不能为空！'
                    },
                    stringLength: {
                        min: 1,
                        max: 50,
                        message: '科室名不能超过50个字符（或25个中文字符）'
                    }
                }
            },
            detail: {
                validators: {
                    notEmpty: {
                        message: '信息不能为空！'
                    },
                }
            },
        }
    });
    $("#validateBtn").click(function(){
        $("#departmentForm").bootstrapValidator('validate');
    });
});


function modal1() {
    var departmentTxt=document.getElementById("department").innerText;
    var department=JSON.parse(departmentTxt).Items;

    //获取已勾选记录的id
    var id = "";

    $("input[name='check']:checkbox:checked").each(function () {
        id = $(this).val();
    })

    for (var d in department) {//根据id定位到这条记录
        if (id == department[d].DeptId) {
            document.getElementById("DepartmentName2").value = department[d].DeptName;
            document.getElementById("detail2").innerText = department[d].Info;
            document.getElementById("DepartmentId").value=department[d].DeptId;
        }
    }
}


function checkSelected1() {
    //判断是否只选择了一条记录
    var check=document.getElementsByName("check");
    var flag=false;
    var count=0;
    for(var c=0;c<check.length;c++) {
        if (check[c].checked == true) {
            count++;
        }
    }
    if(count==1)
        flag = true;
    else if(count==0)
        alert("请选择一条记录！");
    else
        alert("只能选择一条记录！");
    return true;
}