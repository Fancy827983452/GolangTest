window.onload=function () {
    var jsonTxt=document.getElementById("formValue").innerText;
    var obj=JSON.parse(jsonTxt)
    // document.getElementById("publicKey").value=obj.PublicKey;
    document.getElementById("userCenter").href="/user/editInfo";
    document.getElementById("userCenter1").innerText=obj.Name;
}