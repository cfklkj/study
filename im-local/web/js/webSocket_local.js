var lSocket = lSocket || {};
var localSocket = lSocket.Method = { 
    wsLocal:null, 
    loginUser:"",  
    url:"",
    isLoginUser:function(user){
        return this.loginUser == user
    },
    hearMsg:function(){
        msg = {"Operation":imDefine.Opt_heart, "Body":{}}
        return JSON.stringify(msg)
    },
    ReadMsgLenth:function(user){
        var data =  {"Act":imDefine.Act_msgLen, "ConversationId": user}  
        wsLocal.send(JSON.stringify(data));  
    },
    mkReadMsg:function(user, index){ 
       var data =  {"Act":imDefine.Act_read, "ConversationId": user, "Index":index}  
       return JSON.stringify(data)
    }, 
    mkSetMsg:function(user, msg){ 
       msgStr = JSON.stringify(msg)
       var data =  {"Act":imDefine.Act_write, "ConversationId": user, "Data":msgStr}  
       return JSON.stringify(data)
    }, 
    pubMsg:function(msg){ 
        var data =  {"Act":imDefine.Act_pub,  "Data":msg}  
        wsLocal.send(JSON.stringify(data)); 
    },
    pubLoginMsg:function(user){
        this.loginUser = user
        var data =  {"Act":imDefine.Act_Login, "ConversationId": user}  
        this.pubMsg(data)
    },
    onlineMsg:function(to){
        dataMsg = {"type":imDefine.Act_Login, "content":this.loginUser}
        var data =  {"Act":imDefine.Act_route, "ConversationId": to, "Data":dataMsg}  
        wsLocal.send(JSON.stringify(data)); 
        return dataMsg 
    },
    tallMsg:function(to, data, dataType){
        dataMsg = {"type":dataType, "content":data}
        var data =  {"Act":imDefine.Act_route, "ConversationId": to, "Data":dataMsg}  
        wsLocal.send(JSON.stringify(data)); 
        return dataMsg
    }, 
    getLoginInfo: function(userName, password) {  
      //  body = 'username=' +userName+ "&password=" +password; 
        body = {"Username":userName, "Password":password}
        var data =  {"Act":imDefine.Act_Login, "Data":JSON.stringify(body)}   
        wsLocal.send(JSON.stringify(data)); 
    },
    send:function(msg){
         //发送文本  
         wsLocal.send(msg); 
    },
    connect:function(url, username, password){   //"wsLocal://localhost:8080/msg" 
        wsLocal = new WebSocket(url);
        //readyState属性返回实例对象的当前状态，共有四种。
        //CONNECTING：值为0，表示正在连接。
        //OPEN：值为1，表示连接成功，可以通信了。
        //CLOSING：值为2，表示连接正在关闭。
        //CLOSED：值为3，表示连接已经关闭，或者打开连接失败
        //例如：if (wsLocal.readyState == WebSocket.CONNECTING) { }
        
        //【用于指定连接成功后的回调函数】
        wsLocal.onopen = function (evt) {
            console.log("Connection open ..."); 
            localSocket.getLoginInfo(username, password)
           // localSocket.send(localSocket.mkSetMsg("dfd", "ddd"))
          //  wsLocal.send(localSocket.hearMsg())
        };
        //wsLocal.addEventListener('open', function (event) {
        //    wsLocal.send('Hello Server!');
        //};
        
        //【用于指定收到服务器数据后的回调函数】
        //【服务器数据有可能是文本，也有可能是二进制数据，需要判断】
        wsLocal.onmessage = function (event) { 
            if(event.data.length < 15)  //换行
            {
                return
            }   
            data = JSON.parse(event.data)   
            if (data.Code == imDefine.Act_read){
                if (data.Data.Data == ""){
                    return
                }  
                data = JSON.parse(data.Data.Data) 
                chatRight.addMsg(data.Sender , data, true)
            } 
            if (data.Code == imDefine.Act_msgLen){
                logLength = data.Data.Index 
                actRecords.add(imDefine.act_selectContentLogLenth, logLength); 
                chatRight.chat_logFirst = true;
                chatRight.upLog(); 
            }
            if (data.Code == imDefine.Act_Login){ 
                var tbase64 = new Base64()
                outDiv.resLogins(tbase64.decode(data.Data.Data))
            }
            if (data.Code == imDefine.Act_pub) { 
                inDiv.pubInfo(data.Data) 
            }
            if (data.Code == imDefine.Act_route){
                inDiv.routeInfo(data.Data)
            }
        };
        
        //[【于指定连接关闭后的回调函数。】
        wsLocal.onclose = function (evt) {
            console.log("Connection closed."); 
        }; 
        //webSocket.onerror 用于指定报错时的回调函数
        
        wsLocal.onerror = function (event) {
        };
        
        wsLocal.addEventListener("error", function (event) {
        });
    }
}
 