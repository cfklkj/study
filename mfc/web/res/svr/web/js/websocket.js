var Socket = Socket || {};
var imSocket = Socket.Method = { 
    ws:null,  
    loginUser:"", 
    isConnect:false,
    heartTick:new Date().getTime(),  
    upHeartTick:function(){
        imSocket.heartTick = new Date().getTime();
    },
    hearMsg:function(){
        nowTick = new Date().getTime();
        step = nowTick - imSocket.heartTick 
        if (step < 90*1000)
            return 
            imSocket.heartTick = nowTick
        if (this.isConnect) { 
            msg = {"Operation":imDefine.Opt_heart, "Body":{}}
            this.ws.send(JSON.stringify(msg))
        }
    },
    loginMsg:function(userid, userSig, appid){
       msg = {"Operation":imDefine.Opt_login, "Body":{ "UserId":userid,"UserSig" : userSig,"AppId" : appid }}
       imSocket.loginUser = userid;
       this.ws.send(JSON.stringify(msg)) 
    },
    tallMsg:function(toUser, data, dataType){
        dataMsg = {"type":dataType, "content":data}
        body = {"Type":"txt","Desc":"","Data": JSON.stringify(dataMsg),"Ext":""}
        msg = {"Operation":imDefine.Opt_msg, "Body":{"ConverstaionType":imDefine.Conversation_c2c,"ConverstaionId":toUser,"MsgId":0,"Rand":0,"Time":0,"IsSelf":false,"Status":0,"Sender":"","TIMElemet":[body]}}
      
        this.ws.send(JSON.stringify(msg))
        return body
    }, 
    connect:function(url, user, sig, appid){   //"ws://localhost:8080/msg"
        
        this.ws = new WebSocket(url);
        //readyState属性返回实例对象的当前状态，共有四种。
        //CONNECTING：值为0，表示正在连接。
        //OPEN：值为1，表示连接成功，可以通信了。
        //CLOSING：值为2，表示连接正在关闭。
        //CLOSED：值为3，表示连接已经关闭，或者打开连接失败
        //例如：if (ws.readyState == WebSocket.CONNECTING) { }
        
        //【用于指定连接成功后的回调函数】
        this.ws.onopen = function (evt) { 
            this.isConnect = true
            imSocket.loginMsg(user, sig, appid); 
        };
        //this.ws.addEventListener('open', function (event) {
        //    this.ws.send('Hello Server!');
        //};
        
        //【用于指定收到服务器数据后的回调函数】
        //【服务器数据有可能是文本，也有可能是二进制数据，需要判断】
        this.ws.onmessage = function (event) { 
            if(event.data.length < 3)  //换行
            {
                return
            } 
            imSocket.upHeartTick();
            data = JSON.parse(event.data) 
            console.log( data);
            if(data.Body.Status == imDefine.Err_null) {
                if( data.Operation == imDefine.Opt_login){
                    inDiv.showMain();
                   // chat.showChat(); 
                    return
                }
            }
            if (data.Operation == imDefine.Opt_msg){
            }
            if (data.Operation == imDefine.Opt_revmsg){
                chatLeft.addContent(data.Body.Sender, data.Body.TIMElemet[0])
            }
        };
        
        //[【于指定连接关闭后的回调函数。】
        this.ws.onclose = function (evt) {
            console.log("Connection closed.");
            this.isConnect = false
        }; 
        //webSocket.onerror 用于指定报错时的回调函数
        
        this.ws.onerror = function (event) {
        };
        
        this.ws.addEventListener("error", function (event) {
        });
    }
}
 