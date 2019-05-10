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