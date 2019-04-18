window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("userCenter").href="/user/editInfo";
    document.getElementById("userCenter1").innerText=obj.Name;

    var illcaseTxt=document.getElementById("illcase").innerText;
    var illcase=JSON.parse(illcaseTxt).Items;
    var tbody=document.getElementById("tbody");
    for(var i in illcase){
        var tr=document.createElement("tr");
        tbody.appendChild(tr);
        var td1=document.createElement("td");
        var check=document.createElement("input");
        check.setAttribute("type","checkbox");
        check.setAttribute("name","check");
        check.setAttribute("id","check"+i);
        check.setAttribute("type","checkbox");
        check.setAttribute("value",illcase[i].ID);
        td1.appendChild(check);
        tr.appendChild(td1);
        // var td2=document.createElement("td");
        // td2.innerText=illcase[i].ID;
        // tr.appendChild(td2);
        var td3=document.createElement("td");
        td3.innerText=illcase[i].Time.substring(0,10);
        tr.appendChild(td3);
        var td4=document.createElement("td");
        td4.innerText=illcase[i].DeseaseName;
        tr.appendChild(td4);
        var td5=document.createElement("td");
        td5.innerText=illcase[i].DepName;
        tr.appendChild(td5);
        var td6=document.createElement("td");
        td6.innerText=illcase[i].DoctorName;
        tr.appendChild(td6);
        var td7=document.createElement("td");
        td7.innerText=illcase[i].HospitalName;
        tr.appendChild(td7);
        var td8=document.createElement("td");
        var span=document.createElement("span");
        if(illcase[i].Status==1)
            span.className="glyphicon glyphicon-lock";
        td8.appendChild(span);
        tr.appendChild(td8)
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

function confirmLock(){
    if(confirm('确定锁定以上记录吗？'))
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
                document.userForm.action='/user/illcase/lockRecord';
                document.userForm.method='post';
                document.userForm.submit();
            }
        }
        if(flag==false)
            alert("至少要选择一条记录！");
        return true;
    }
    else
        return false;
}

function confirmUnlock()
{
    if(confirm('确定解锁以上记录吗？'))
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
                document.userForm.action='/user/illcase/unlockRecord';
                document.userForm.method='post';
                document.userForm.submit();
            }
        }
        if(flag==false)
            alert("至少要选择一条记录！");
        return true;
    }
    else
        return false;
}

function modal() {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    var illcaseTxt=document.getElementById("illcase").innerText;
    var illcase=JSON.parse(illcaseTxt).Items;

    document.getElementById("modal_name").innerText=obj.Name;
    //根据出生日期计算周岁年龄
    var birth=obj.BirthDate.substring(0,10);
    var age=jsGetAge(birth);
    document.getElementById("modal_age").innerText=age;
    if(obj.Gender==0)
        document.getElementById("modal_gender").innerText="男";
    else
        document.getElementById("modal_gender").innerText="女";

    //获取已勾选记录的id
    var id = "";
    $("input[name='check']:checkbox:checked").each(function () {
        id = $(this).val();
    })

    for (var i in illcase) {//根据id定位到这条记录
        if (id == illcase[i].ID) {
            document.getElementById("modal_illname").innerText = illcase[i].DeseaseName;
            document.getElementById("modal_symptom").innerText = illcase[i].Symptom;
            document.getElementById("modal_illdetail").innerText = illcase[i].Info;
            document.getElementById("modal_curedoctor").innerText = illcase[i].DoctorName;
            document.getElementById("modal_department").innerText = illcase[i].DepName;
            document.getElementById("modal_hospital").innerText = illcase[i].HospitalName;
            document.getElementById("modal_time").innerText = illcase[i].Time;
        }
    }
}

function checkSelectedOne() {
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