var LoginDiv = LoginDiv || {};
var outDiv = LoginDiv.Method = { 
    div_login:'Login',  
    showLogin:function(){
        eleBody = util.getEleById("bodys");
        util.clearElement(eleBody);
        util.addElement(eleBody, util.addDiv(this.div_login)); 
        login = util.getEleById(this.div_login);
        util.addElement(login, loginHtml); 
        util.addEvent("logins", "outDiv.logins()");
    },
    logins:function(){
        localSocket.getLoginInfo("user01", "123456")
    },
    resLogins:function(msg){  
        if (msg.Data.indexOf('{') < 0){
            console.debug(msg.Data)
            return
        }
        data = JSON.parse(msg.Data)
        if (data.code == 200){ 
            memHash.add(imDefine.mem_loginData, data.data)
            imSocket.connect("ws://122.51.33.81:10031/v1/tzj", data.data.userid, data.data.sign, data.data.appid)
        }
        console.debug(JSON.parse(msg.Data)) 
    }
}

var loginHtml='<div class="logo">\
<img  alt="" src="/image/exeIcon.png">\
</div>\
<div class="box">\
<h2>Login</h2>\
<div>\
    <div class="inputbox">\
        <input type="text" name="" required="">\
        <label>Username</label>\
    </div>\
    <div class="inputbox">\
        <input type="password" name=" " required="">\
        <label>Password</label>\
    </div>\
    <input type="button" name="" value="Login" id="logins" >\
</div>\
</div>'