var MgrDiv = MgrDiv || {};
var inDiv = MgrDiv.Method = { 
    div_main:'in_main',  
    div_left:'in_left',  
    div_right:'in_right',  
    l_me:"l_me",
    l_chat:"l_chat",
    l_black:"l_black",
    l_set:"l_set",
    showMain:function(){
        eleBody = util.getEleById("bodys");
        util.clearElement(eleBody);  
        util.addElement(eleBody, mainHtml);
        this.paintLeft();
    }, 
    paintLeft:function(){
        eleBody = util.getEleById(this.div_left); 
        util.addElement(eleBody, leftHtml); 
        memHash.add(imDefine.mem_Lchoice, this.l_chat)
        this.paintRight();
    },
    paintRight:function(){
       switch( memHash.get(imDefine.mem_Lchoice)){
           case this.l_chat:
               chat.showChat()
       }
    },
    isChoiceChat:function(){
        return memHash.get(imDefine.mem_Lchoice) == this.l_chat
    }
}
 
var mainHtml='<div id="in_main" class="in_main">\
<div id="in_left" class="in-left"></div>\
<div id="in_right" class="in-right"></div>\
<div>'

var leftHtml='<div id="l_me"><img  alt="" src="/image/me.png"/></div>\
<div id="l_chat"><img  alt="" src="/image/info.png"/></div>\
<div id="l_black"><img  alt="" src="/image/black.png"/></div>\
<div id="l_set"><img  alt="" src="/image/set.png"/></div>\
'