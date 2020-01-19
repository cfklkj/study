var Socket = Socket || {};
var imSocket = Socket.Method = { 
    ws:null,  
    name:"",
    hearMsg:function(){
        msg = {"Operation":imDefine.Opt_heart, "Body":{}}
        return JSON.stringify(msg)
    },
    loginMsg:function(userid, userSig, appid){
       msg = {"Operation":imDefine.Opt_login, "Body":{ "UserId":userid,"UserSig" : userSig,"AppId" : appid }}
       name = userid;
       return JSON.stringify(msg)
    },
    tallMsg:function(toUser, data){
        body = {"Type":"txt","Desc":"","Data":data,"Ext":""}
        msg = {"Operation":imDefine.Opt_msg, "Body":{"ConverstaionType":imDefine.Conversation_c2c,"ConverstaionId":toUser,"MsgId":0,"Rand":0,"Time":0,"IsSelf":false,"Status":0,"Sender":"","TIMElemet":[body]}}
        return JSON.stringify(msg)
    },
    send:function(toUser, msg){
            //发送文本
        ws.send(imSocket.tallMsg(toUser, msg)); 
    },
    connect:function(url, user, sig, appid){   //"ws://localhost:8080/msg"
        
        ws = new WebSocket(url);
        //readyState属性返回实例对象的当前状态，共有四种。
        //CONNECTING：值为0，表示正在连接。
        //OPEN：值为1，表示连接成功，可以通信了。
        //CLOSING：值为2，表示连接正在关闭。
        //CLOSED：值为3，表示连接已经关闭，或者打开连接失败
        //例如：if (ws.readyState == WebSocket.CONNECTING) { }
        
        //【用于指定连接成功后的回调函数】
        ws.onopen = function (evt) {
            console.log("Connection open ...");
            ws.send(imSocket.loginMsg(user, sig, appid));
        };
        //ws.addEventListener('open', function (event) {
        //    ws.send('Hello Server!');
        //};
        
        //【用于指定收到服务器数据后的回调函数】
        //【服务器数据有可能是文本，也有可能是二进制数据，需要判断】
        ws.onmessage = function (event) { 
            if(event.data.length < 3)  //换行
            {
                return
            } 
            data = JSON.parse(event.data) 
            console.log( data);
            if(data.Body.Status == imDefine.Err_null) {
                if( data.Operation == imDefine.Opt_login){
                    chat.showChat();
                    return
                }
            }
            if (data.Operation == imDefine.Opt_msg){
            }
            if (data.Operation == imDefine.Opt_revmsg){
                chat.addMsg(data.Body.Sender, data.Body.TIMElemet[0]) 
            }
        };
        
        //[【于指定连接关闭后的回调函数。】
        ws.onclose = function (evt) {
            console.log("Connection closed.");
        }; 
        //webSocket.onerror 用于指定报错时的回调函数
        
        ws.onerror = function (event) {
        };
        
        ws.addEventListener("error", function (event) {
        });
    }
}
 