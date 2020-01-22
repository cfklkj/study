var layoutChatLeft = layoutChatLeft || {};
var chatLeft = layoutChatLeft.Method = {  
    div_top:'chat_l_t',
    div_bottom:'chat_l_b',
    show:function(){
        eleBody = util.getEleById(chat.div_chat);
        util.addElement(eleBody, chatLeftHtml);  
    },   
    altContent_msg:function(user, msg, times){
        ele = util.getEleById(this.div_bottom).getElementsByTagName("p") 
        for (var i=0; i < ele.length; i++){
            obj = ele[i]
            if (obj.getAttribute("name") == user){
                obj.getElementsByTagName("span")[1].innerText = times
                obj.getElementsByTagName("span")[2].innerText =  msg 
                break
            }
        }
    },
    addContent:function(user, msg){ 
        times = util.timeHms() 
        data = JSON.parse(msg.Data)
        datas = data.content.substring(0, 16)
        localSocket.send(localSocket.mkSetMsg(user, msg))
        if (memHash.find(user)){ 
            this.altContent_msg(user, datas, times)  //会话列表 
            return
        }
        memHash.add(user,1); 
        if (!inDiv.isChoiceChat())
        {
            return
        }
        img = "image/chatCuster.png"
        str = '<p onclick=chatLeft.selectContent(this) name="'+user+'"><img alt="" src=' + img + 
            '></img><span index="1">' + user.substring(0, 11) +'</span><span index="2">' + times  +
            '</span><span index="3">' + datas +'</span></p>';
        content = document.getElementById(this.div_bottom);
        content.innerHTML = content.innerHTML + str;
    },
    selectContent:function(event){
       user = event.getAttribute("name")  
        if(actRecords.get(imDefine.act_selectContent) != user) { 
            this.clearContent();
            actRecords.add(imDefine.act_selectContent, user);  
            actRecords.add(imDefine.act_selectContentLogLenth, 0); 
            localSocket.send(localSocket.mkReadMsgLenth(user));
        } 
    },
    clearContent:function(){
        content = document.getElementById(this.div_bottom);
        content.innerHTML = ""; 
    } 
}
 
var chatLeftHtml='<div class="chat_l">\
<div  id="chat_l_t" ><div class="chat_l_t_search">\
<span><img  alt="" src="/image/search2.png"/></span>\
<input type="text" name="" required="" maxlength="11" placeholder="search">\
</div>\
</div>\
<div id="chat_l_b"></div>\
</div>'
