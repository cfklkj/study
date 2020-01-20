var LoginDiv = LoginDiv || {};
var inDiv = LoginDiv.Method = { 
    div_login:'Login', 
    showLogin:function(){
        eleBody = util.getEleById("bodys");
        util.clearElement(eleBody);
        util.addElement(eleBody, util.addDiv(this.div_login)); 
        login = util.getEleById(this.div_login);
        util.addElement(login, loginHtml); 
        util.addEvent("logins", "inDiv.logins()");
    },
    logins:function(){
        doHttp.getLoginInfo("http://134.175.145.46:8030/im/info", "admin", "123456",this.resLogins)
    },
    resLogins:function(msg){  
        data = JSON.parse(msg.Data)
        if (data.code == 200){ 
            imSocket.connect("ws://122.51.33.81:10031/v1/tzj", data.data.userid, data.data.sign, data.data.appid)
        }
        console.debug(JSON.parse(msg.Data)) 
    }
}

var loginHtml='<div class="box">\
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
    <input type="button" name="" value="submit" id="logins" >\
</div>\
</div>'