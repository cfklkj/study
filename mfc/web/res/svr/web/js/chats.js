var layoutChat = layoutChat || {};
var chat = layoutChat.Method = { 
    div_chat:'chat',
    chat_isUp:false,
    chat_logFirst:false,
    chat_words:null,
    showChat:function(){
        eleBody = util.getEleById("bodys");
        util.clearElement(eleBody);
        util.addElement(eleBody, util.addDiv(this.div_chat)); 
        login = util.getEleById(this.div_chat);
        util.addElement(login, chatHtml); 
        util.addEvent("talksub", "chat.sendMsg()");  
        util.addScrollEvt("words", "chat.upLog(this)");
        util.addEvent("fileUpdate-button", "imgDlg.openDlg()")
        util.addEvt("file-button", "onchange", "imgDlg.changes(this)")
        this.chat_words = util.getEleById("words")
    }, 
    upLog:function(event){ 
        if ( this.chat_logFirst || event.scrollTop == 0){  
            this.chat_logFirst = false
            logLength = actRecords.get(imDefine.act_selectContentLogLenth);
            if (logLength < 1) {
                return
            }  
            if(this.chat_isUp)
            {
                return
            }
            this.chat_isUp = true
            user = actRecords.get(imDefine.act_selectContent);
            for ( i = logLength; i > 0; i--)
            {
                if (i + 17 == logLength) {
                    break
                }
                localSocket.send(localSocket.mkReadMsg(user, i));
                actRecords.add(imDefine.act_selectContentLogLenth, i - 1); 
            }
            this.chat_isUp = false
        }
    },
    send:function(user, msgData, dataType){
        data = imSocket.tallMsg(user, msgData, dataType) 
        chat.addMsg(imSocket.loginUser, data)
        data.Sender = imSocket.loginUser
        localSocket.send(localSocket.mkSetMsg(user, data))
    },
    sendMsg:function(){ 
        user = actRecords.get(imDefine.act_selectContent)
        if (user == null){
            //user = "7a57a5a743894a0e"
            return false
        }
        msgData = util.getEleValue("talkwords") 
        if (msgData == "") {
            return
        }
        chat.send(user, msgData, imDefine.chat_chat)
    },
    addMsg:function(sender, jsData, isLog){ 
        //data.Sender, data.Type, data.Data 
        body = JSON.parse(jsData.Data)
        htmls = ""
        switch(body.type){
            case imDefine.chat_chat:
                htmls = body.content;
                break;
            case imDefine.chat_img:
                htmls = '<img alt="" src=' + body.content + '></img>'
                break; 
        }
        if (htmls == "")
        {
            return
        }
        if (sender == imSocket.loginUser)
        {
            str = '<div class="btalk"><span>' + htmls +'</span></div>'; 
        }else{
            str = '<div class="atalk"><span>' + htmls +'</span></div>';
        }
        var Words = document.getElementById("words");
        if (isLog) {   //逆
            Words.innerHTML = str + Words.innerHTML; 
            this.chat_words.scrollTop = this.chat_words.scrollHeight /2
        }else{   //正
             Words.innerHTML = Words.innerHTML + str;
             this.chat_words.scrollTop = this.chat_words.scrollHeight 
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
 
var chatHtml='<div class="talk_con">\
<div class="div-left"><div class="messageContent" id="msgContent"></div></div>\
<div class="div-right">\
<div class="talk_show" id="words">\
    <div class="atalk"><span id="asay">A说：吃饭了吗？</span></div>\
    <div class="btalk"><span id="bsay">B说：还没呢，你呢？</span></div>\
</div>\
<div class="talk_input">\
    <textarea class="talk_word" id="talkwords"></textarea>\
    <input type="file" id="file-button" style="display: none;" name="qiniuPic" accept="image/*" >\
    <div class="file_scan" style="width: 100px;float: left; margin-top: 20px;" id="fileUpdate-button">\
        添加浏览\
    </div>\
    <input type="button" value="发送" class="talk_sub" id="talksub">\
</div>\
</div>\
</div>'