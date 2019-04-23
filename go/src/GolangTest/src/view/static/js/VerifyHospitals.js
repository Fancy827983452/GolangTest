//由于使用confirm弹窗确实用户操作，点击取消时会有页面跳转（？）
//onbeforeunload在即将离开当前页面（刷新或关闭）时触发
//由用户进行判断是否跳转页面
// window.onbeforeunload=function(){
//     return '确定要离开吗？'
// }

function loading() {
    var jsonTxt=document.getElementById("formValue").innerText;
    if(jsonTxt) {
        var hospitalsTxt = document.getElementById("hospitals").innerText;
        if (hospitalsTxt) {
            var hospitals = JSON.parse(hospitalsTxt).Items;
            var tbody = document.getElementById("tbody");
            for (var i in hospitals) {
                var tr = document.createElement("tr");
                tbody.appendChild(tr);
                var td1 = document.createElement("td");
                var check = document.createElement("input");
                check.setAttribute("type", "checkbox");
                check.setAttribute("name", "check");
                check.setAttribute("id", "check" + i);
                check.setAttribute("type", "checkbox");
                check.setAttribute("value", hospitals[i].HospitalId);
                td1.appendChild(check);
                tr.appendChild(td1);
                var td2 = document.createElement("td");
                td2.innerText = hospitals[i].Name;
                tr.appendChild(td2);
                var td3 = document.createElement("td");
                td3.innerText = hospitals[i].Grade;
                tr.appendChild(td3);
                var td4 = document.createElement("td");
                td4.innerText = hospitals[i].Location;
                tr.appendChild(td4);
                var td5 = document.createElement("td");
                var status = hospitals[i].Status;
                if (status == -1)
                    td5.innerText = "不通过";
                else if (status == 0)
                    td5.innerText = "待审核";
                else if (status == 1)
                    td5.innerText = "审核通过";
                else if (status == 2)
                    td5.innerText = "异常";
                tr.appendChild(td5);
            }
            //根据url路径设置下拉框的值
            var url=window.location.href;
            var select=document.getElementById("selectStatus");
            var options = select.children;
            if(url.indexOf("jsVerifyHospitals/-1")!=-1)
                options[1].selected=true;
            else if(url.indexOf("jsVerifyHospitals/0")!=-1)
                options[2].selected=true;
            else if(url.indexOf("jsVerifyHospitals/1")!=-1)
                options[3].selected=true;
            else if(url.indexOf("jsVerifyHospitals/2")!=-1)
                options[4].selected=true;
            else
                options[0].selected=true;
        }
    }
    else {
        alert("登录已过期，请重新登陆！")
        window.location.href="login";
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

// function changeStatus() {
//     var select=document.getElementById("selectStatus");
//     if(select.value!="") {//如果选中了
//         var index = select.selectedIndex;
//         var status = select.options[index].value;
//         document.userForm.action = '/admin/jsVerifyHospitals/?status='+status;
//         document.userForm.method = 'get';
//         document.userForm.submit();
//     }
// }

function confirmVerify() {
    if (confirm('确定审核通过以上记录吗？')) {
        var str = "";
        $("input[name='check']:checkbox:checked").each(function () {
            str += $(this).val() + ",";
        })
        $("#selectedItem").val(str);
        var check = document.getElementsByName("check");
        var flag = false;
        for (var c = 0; c < check.length; c++) {
            if (check[c].checked == true) {
                flag = true;
                document.userForm.action = '/admin/verifyHospitals/pass';
                document.userForm.method = 'post';
                document.userForm.submit();
            }
        }
        return true;
    }
    else {
        return false;
    }
}

function confirmDeny()
{
    if(confirm('确定审核不通过以上记录吗？'))
    {
        var str="";
        $("input[name='check']:checkbox:checked").each(function(){
            str+=$(this).val()+",";
        })
        $("#selectedItem").val(str);
        var check=document.getElementsByName("check");
        var flag=false;
        for(var c=0;c<check.length;c++)
        {
            if(check[c].checked==true)
            {
                flag=true;
                document.userForm.action='/admin/verifyHospitals/fail';
                document.userForm.method='post';
                document.userForm.submit();
            }
        }
        return true;
    }
    else
        return false;
}

function checkSelected() {
    //判断是否选择记录
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
        // window.history.back(-1);
    }
    return true;
}