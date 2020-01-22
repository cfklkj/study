var layoutChatRight = layoutChatRight || {};
var chatRight = layoutChatRight.Method = {  
    div_top:'chat_l_t',
    div_middle:'chat_l_m',
    div_bottom:'chat_l_b',
    chat_isUp:false,
    chat_logFirst:false,
    chat_words:null,
    show:function(){
        eleBody = util.getEleById(chat.div_chat);
        util.addElement(eleBody, chatRightHtml);  
 
        util.addEvent("talksub", "chat.sendMsg()");  
        util.addScrollEvt(this.div_middle, "chatRight.upLog(this)");
        util.addEvent("fileUpdate-button", "imgDlg.openDlg()")
        util.addEvt("file-button", "onchange", "imgDlg.changes(this)")
        this.chat_words = util.getEleById(this.div_middle)
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
        this.send(user, msgData, imDefine.chat_chat)
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
}
 
var chatRightHtml='<div class="chat_r">\
<div  class="chat_r_t" >\
<span id="chat_r_t_name">Test</span>\
<i></i>\
</div>\
<div  id="chat_r_m"></div>\
<div  id="chat_r_b">\
<input type="file" id="file-button" style="display: none;" name="qiniuPic" accept="image/*" >\
<span index="1"><img  alt="" src="/image/file.png"/></span>\
<textarea  id="talkwords"></textarea>\
<input type="button" value="发送" id="talksub"></input>\
</div>\
</div>'

// var chatHtml='<div class="talk_con">\
// <div class="div-right">\
// <div class="talk_show" id="words">\
//     <div class="atalk"><span id="asay">A说：吃饭了吗？</span></div>\
//     <div class="btalk"><span id="bsay">B说：还没呢，你呢？</span></div>\
// </div>\
// <div class="talk_input">\
//     <textarea class="talk_word" id="talkwords"></textarea>\
//     <input type="file" id="file-button" style="display: none;" name="qiniuPic" accept="image/*" >\
//     <div class="file_scan" style="width: 100px;float: left; margin-top: 20px;" id="fileUpdate-button">\
//         添加浏览\
//     </div>\
//     <input type="button" value="发送" class="talk_sub" id="talksub">\
// </div>\
// </div>\
// </div>'