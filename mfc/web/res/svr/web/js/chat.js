var layoutChat = layoutChat || {};
var chat = layoutChat.Method = { 
    div_chat:'chat_main',
    chat_isUp:false,
    chat_logFirst:false,
    chat_words:null,
    showChat:function(){
        eleBody = util.getEleById(inDiv.div_right);
        util.clearElement(eleBody); 
        util.addElement(eleBody, chatHtml);   
        chatLeft.show(); 
        chatRight.show(); 
    },   
} 

var chatHtml='<div id="chat_main"></div>'