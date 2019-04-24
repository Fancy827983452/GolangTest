//由于使用confirm弹窗确实用户操作，点击取消时会有页面跳转（？）
//onbeforeunload在即将离开当前页面（刷新或关闭）时触发
//由用户进行判断是否跳转页面
window.onbeforeunload=function(){
    return '确定要离开吗？'
}

function loading() {
    var jsonTxt=document.getElementById("formValue").innerText;
    if(jsonTxt) {
        var obj = JSON.parse(jsonTxt);
        document.getElementById("hospitalCenter1").innerText=obj.Name;
        document.getElementById("publicKey").value = obj.PublicKey;

        var doctorsTxt = document.getElementById("doctors").innerText;
        if(doctorsTxt) {
            var doctors = JSON.parse(doctorsTxt).Items;
            var tbody = document.getElementById("tbody");
            for (var i in doctors) {
                var tr = document.createElement("tr");
                tbody.appendChild(tr);
                var td1 = document.createElement("td");
                var check = document.createElement("input");
                check.setAttribute("type", "checkbox");
                check.setAttribute("name", "check");
                check.setAttribute("id", "check" + i);
                check.setAttribute("type", "checkbox");
                check.setAttribute("value",doctors[i].DoctorKey);
                td1.appendChild(check);
                tr.appendChild(td1);
                var td3 = document.createElement("td");
                td3.innerText = doctors[i].Name;
                tr.appendChild(td3);
                var td4 = document.createElement("td");
                if (doctors[i].Gender == 0)
                    td4.innerText = "男";
                else
                    td4.innerText = "女";
                tr.appendChild(td4);
                var td5 = document.createElement("td");//年龄
                td5.innerText = jsGetAge(doctors[i].BirthDate);
                tr.appendChild(td5);
                var td8 = document.createElement("td");//科室
                td8.innerText = doctors[i].DeptName;
                tr.appendChild(td8);
                var td9 = document.createElement("td");//职称
                if(doctors[i].Title==0)
                    td9.innerText ="医士、医师、住院医师";
                else if(doctors[i].Title==1)
                    td9.innerText ="主治医师";
                else if(doctors[i].Title==2)
                    td9.innerText ="副主任医师";
                else if(doctors[i].Title==3)
                    td9.innerText ="主任医师";
                tr.appendChild(td9);
                var td10 = document.createElement("td");//身份
                if(doctors[i].Role==1)
                    td10.innerText="管理员";
                tr.appendChild(td10);
            }
        }
        //根据url路径设置下拉框的值
        var url=window.location.href;
        var select=document.getElementById("select");
        var options = select.children;
        if(url.indexOf("searchDoctor/name")!=-1)
            options[0].selected=true;
        else if(url.indexOf("searchDoctor/gender")!=-1)
            options[1].selected=true;
        else if(url.indexOf("searchDoctor/department_name")!=-1)
            options[2].selected=true;
        else if(url.indexOf("searchDoctor/title")!=-1)
            options[3].selected=true;
        else if(url.indexOf("searchDoctor/role")!=-1)
            options[4].selected=true;
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

function confirmRole() {
    if (confirm('确定设置为管理员吗？')) {
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
                document.userForm.action = '/hospital/setAdmin';
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

function confirmCancel()
{
    if(confirm('确定取消管理员身份吗？'))
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
                document.userForm.action='/hospital/cancelAdmin';
                document.userForm.method='post';
                document.userForm.submit();
            }
        }
        return true;
    }
    else
        return false;
}

function confirmWithdraw()
{
    if(confirm('确定注销以上账户吗？'))
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
                document.userForm.action='/hospital/verifyDoctor/withdraw';
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
    }
    return true;
}

//根据出生日期计算周岁年龄
function jsGetAge(strBirthday)
{
    var returnAge;
    var strBirthdayArr=strBirthday.split("-");
    var birthYear = strBirthdayArr[0];
    var birthMonth = strBirthdayArr[1];
    var birthDay = strBirthdayArr[2];

    d = new Date();
    var nowYear = d.getFullYear();
    var nowMonth = d.getMonth() + 1;
    var nowDay = d.getDate();

    if(nowYear == birthYear)
        returnAge = 0;//同年 则为0岁
    else
    {
        var ageDiff = parseInt(nowYear - birthYear) ; //年之差
        if(ageDiff > 0) {
            if(nowMonth == birthMonth) {
                var dayDiff = nowDay - birthDay;//日之差
                if(dayDiff < 0)
                    returnAge = ageDiff - 1;
                else
                    returnAge = ageDiff ;
            }
            else {
                var monthDiff = nowMonth - birthMonth;//月之差
                if(monthDiff < 0)
                    returnAge = ageDiff - 1;
                else
                    returnAge = ageDiff ;
            }
        }
        else
            returnAge = -1;//返回-1 表示出生日期输入错误 晚于今天
    }
    return returnAge;//返回周岁年龄
}

function search() {
    var param=document.getElementById("select").value;
    document.userForm.action='/hospital/searchDoctor/'+param;
    document.userForm.method='post';
    document.userForm.submit();
}