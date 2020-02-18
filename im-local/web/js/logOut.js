var LoginDiv = LoginDiv || {};
var outDiv = LoginDiv.Method = { 
    div_login:'Login',  
    id_username:"username",
    id_password:"password",
    showLogin:function(){
        eleBody = util.getEleById(imDefine.div_bodys); 
        util.clearElement(eleBody);
        div = util.addEle("div")
        div.id=this.div_login
        eleBody.appendChild(div)
        div.innerHTML = loginHtml 
        util.addEvent("logins", "outDiv.logins()"); 
    },
    logins:function(){ 
        username = util.getEleById(this.id_username).value
        password = util.getEleById(this.id_password).value
        if (username == "") {
            username = "custer"
        } 
        localSocket.connect("ws://" + util.getUrl()+ "/wss", username, password) 
    },
    resLogins:function(msg){   
         data = JSON.parse(msg)
        if (data.Code == 200){  
            memHash.add(imDefine.mem_loginData, data)
            inDiv.showMain();
            localSocket.pubLoginMsg(data.Data) 
        }  
    }
}

var loginHtml='<div class="logo">\
<img  alt="" src="/image/exeIcon.png">\
</div>\
<div class="box">\
<h2>Login</h2>\
<div>\
    <div class="inputbox">\
        <input type="text" name="" required="" id="username" value="">\
        <label>Username</label>\
    </div>\
    <div class="inputbox">\
        <input type="password" name=" " required=""  id="password" value="">\
        <label>Password</label>\
    </div>\
    <input type="button" name="" value="Login" id="logins" >\
</div>\
</div>'