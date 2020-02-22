var bodyMiddel = bodyMiddel || {};
var bodysMiddel = bodyMiddel.Method = {   
    show:function(ele){  
        this.drawButton(ele)  
    },
    drawButton:function(ele){
        div = util.addEle("div")
        //div.setAttribute("class", "bMiddelButton") 
        div.innerHTML = "-----发送----》"
        div.setAttribute("onclick", "bodysMiddel.sendMsg()")
        ele.appendChild(div) 
    },
    sendMsg:function(){
        data = bodysLeft.div_bottomEdit.innerHTML
        localSocket.send(data)
    } 
}  