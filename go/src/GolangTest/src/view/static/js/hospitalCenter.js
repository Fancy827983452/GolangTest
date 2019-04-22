window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    // var url=window.location.href;
    // document.getElementById("doctorCenter").href="#";
    document.getElementById("hospitalCenter1").innerText=obj.Name;
}

$("#quanxuan").click(function(){//给全选按钮加上点击事件
    var xz = $(this).prop("checked");//判断全选按钮的选中状态
    var ck = $(".cb").prop("checked",xz);  //让class名为qx的选项的选中状态和全选按钮的选中状态一致。
})