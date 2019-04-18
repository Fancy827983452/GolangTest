function checkNull() {
    var username = document.getElementById("username").value;
    var pwd = document.getElementById("password").value;
    if (username.length == 0 || pwd.length == 0) {
        alert("用户名或密码不能为空！");
        return false;
    }
}