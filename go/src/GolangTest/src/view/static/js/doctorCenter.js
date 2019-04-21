window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt);
    // document.getElementById("doctorCenter").href="#";
    document.getElementById("doctorCenter1").innerText=obj.Name;
}