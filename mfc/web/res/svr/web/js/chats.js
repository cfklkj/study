var layoutChat = layoutChat || {};
var chat = layoutChat.Method = { 
    div_chat:'chat',
    showChat:function(){
        eleBody = util.getEleById("bodys");
        util.clearElement(eleBody);
        util.addElement(eleBody, util.addDiv(this.div_chat)); 
        login = util.getEleById(this.div_chat);
        util.addElement(login, chatHtml); 
        util.addEvent("talksub", "chat.sendMsg()");  
        util.addScrollEvt("words", "chat.upLog(this)");
    },
    upLog:function(event){
        console.log("upLo-g", event)
        return
        if (event.scrollTop == 0){ 
            logLength = actRecords.get(imDefine.act_selectContentLogLenth); 
            user = actRecords.get(imDefine.act_selectContent);
            for ( i = logLength; i > 0; i--)
            {
                if (i + 17 == logLength) {
                    break
                }
                localSocket.send(localSocket.mkReadMsg(user, i));
                actRecords.add(imDefine.act_selectContentLogLenth, i - 1); 
            }
        }
    },
    sendMsg:function(){
        msgData = util.getEleValue("talkwords") 
        if (msgData == "") {
            return
        }
        user = actRecords.get(imDefine.act_selectContent)
        if (user == null){
            user = "7a57a5a743894a0e"
        }
        msg, data = imSocket.tallMsg(user, msgData)
        imSocket.send(msg)
        chat.addMsg(imSocket.loginUser, data)
        data.Sender = imSocket.loginUser
        localSocket.send(localSocket.mkSetMsg(user, data))
    },
    addMsg:function(sender, jsData, isLog){ 
        //data.Sender, data.Type, data.Data 
        if (sender == imSocket.loginUser)
        {
            str = '<div class="btalk"><span>B说 :' + jsData.Data +'</span></div>'; 
        }else{
            str = '<div class="atalk"><span>A说 :' + jsData.Data +'</span></div>';
        }
        var Words = document.getElementById("words");
        if (isLog) {   //逆
            Words.innerHTML = str + Words.innerHTML; 
        }else{   //正
             Words.innerHTML = Words.innerHTML + str;
        }
    }, 
    altContent_msg:function(user, msg){
        ele = util.getEleById("msgContent").getElementsByTagName("p") 
        for (var i=0; i < ele.length; i++){
            obj = ele[i]
            if (obj.getAttribute("name") == user){
                obj.getElementsByTagName("span")[1].innerText = msg 
                break
            }
        }
    },
    addContent:function(user, msg){ 
        localSocket.send(localSocket.mkSetMsg(user, msg))
        if (memHash.find(user)){ 
            this.altContent_msg(user, msg.Data)  //会话列表
            if(actRecords.get(imDefine.act_selectContent) != user) {   //内容
                chat.addMsg(user, msg) 
            }
            return
        }
        memHash.add(user,1);
        chat.clearContent();
        times = util.timeHms()
        chat.addMsg(user, msg) 
        img = "image/chatCuster.png"
        str = '<p onclick=chat.selectContent(this) name="'+user+'"><img alt="" src=' + img + '></img><span>' + user +':</span><span>' + msg.Data +'</span><span>' + times +'</span></p>';
        var Words = document.getElementById("msgContent");
        Words.innerHTML = Words.innerHTML + str;
    },
    selectContent:function(event){
       user = event.getAttribute("name")  
        if(actRecords.get(imDefine.act_selectContent) != user) { 
            chat.clearContent();
            actRecords.add(imDefine.act_selectContent, user);  
            actRecords.add(imDefine.act_selectContentLogLenth, 0); 
            localSocket.send(localSocket.mkReadMsgLenth(user));
        } 
    },
    clearContent:function(){
        var Words = document.getElementById("words");
        Words.innerHTML = ""; 
    }


}

// window.onscroll = function (){
//     var marginBot = 0;
//     if (document.documentElement.scrollTop){
//     marginBot = document.documentElement.scrollHeight – (document.documentElement.scrollTop+document.body.scrollTop)-document.documentElement.clientHeight;
//     } else {
//     marginBot = document.body.scrollHeight – document.body.scrollTop- document.body.clientHeight;
//     }
//     if(marginBot<=0) {
//     //do something
//     }
//     }
var chatHtml='<div class="talk_con">\
<div class="div-left"><div class="messageContent" id="msgContent"></div></div>\
<div class="div-right">\
<div class="talk_show" id="words">\
    <div class="atalk"><span id="asay">A说：吃饭了吗？</span></div>\
    <div class="btalk"><span id="bsay">B说：还没呢，你呢？</span></div>\
</div>\
<div class="talk_input">\
    <textarea class="talk_word" id="talkwords"></textarea>\
    <input type="button" value="发送" class="talk_sub" id="talksub">\
</div>\
</div>\
</div>'