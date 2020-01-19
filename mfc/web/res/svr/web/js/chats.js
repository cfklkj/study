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
    },
    sendMsg:function(){
        msgData = "hello"
        imSocket.send("7a57a5a743894a0e", msgData)
        chat.addSendMsg(null, msgData)
    },
    addMsg:function(sender, msg){
        str = '<div class="atalk"><span>A说 :' + msg.Data +'</span></div>';
        var Words = document.getElementById("words");
        Words.innerHTML = Words.innerHTML + str;
    },
    addSendMsg:function(msgType, msg){
        str = '<div class="btalk"><span>B说 :' + msg +'</span></div>';
        var Words = document.getElementById("words");
        Words.innerHTML = Words.innerHTML + str;
    }

}

var chatHtml='<div class="talk_con">\
<div class="talk_show" id="words">\
    <div class="atalk"><span id="asay">A说：吃饭了吗？</span></div>\
    <div class="btalk"><span id="bsay">B说：还没呢，你呢？</span></div>\
</div>\
<div class="talk_input">\
    <select class="whotalk" id="who">\
        <option value="0">A说：</option>\
        <option value="1">B说：</option>\
    </select>\
    <input type="text" class="talk_word" id="talkwords">\
    <input type="button" value="发送" class="talk_sub" id="talksub">\
</div>\
</div>'