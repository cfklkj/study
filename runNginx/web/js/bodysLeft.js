var bodyLeft = bodyLeft || {};
var bodysLeft = bodyLeft.Method = {  
    div_main:"bodys",
    div_left:null,
    div_middle:null,
    div_bottomEdit:null,
    show:function(ele){ 
        this.drawTop(ele)
        this.drawBottom(ele)  
    },
    drawTop:function(ele){
        div = util.addEle("div")
        div.setAttribute("class", "bLeftTop") 
        ele.appendChild(div)
        //this.topButtons(div)
        this.drawNginx(div)
    }, 
    drawNginx:function(ele){
        bt = util.addEle("button")
        bt.innerHTML = "检测nginx"
        bt.setAttribute("onclick", "bodysLeft.setBottomEdit(mkMsgsNginx.checkNginx())")
        ele.appendChild(bt)
        bt = util.addEle("button")
        bt.innerHTML = "重启nginx"
        bt.setAttribute("onclick", "bodysLeft.setBottomEdit(mkMsgsNginx.reloadNginx())")
        ele.appendChild(bt)
        bt = util.addEle("button")
        bt.innerHTML = "下载文件"
        bt.setAttribute("onclick", "bodysLeft.setBottomEdit(mkMsgsNginx.downFile('192.168.152.1:10024/download/jVrDvQoss.data', '/tmp/jVrDvQoss.data', 'zip'))")
        ele.appendChild(bt)
        bt = util.addEle("button")
        bt.innerHTML = "生成cert"
        bt.setAttribute("onclick", "bodysLeft.setBottomEdit(mkMsgsNginx.newCert('wxd.baidu.com'))")
        ele.appendChild(bt)  
        bt = util.addEle("button")
        bt.innerHTML = "更新cert"
        bt.setAttribute("onclick", "bodysLeft.setBottomEdit(mkMsgsNginx.renewCert())")
        ele.appendChild(bt)

    }, 
    drawBottom:function(ele){
        div = util.addEle("div")
        div.setAttribute("class", "bLeftBottom")
        ele.appendChild(div)
        this.bottomEdit(div)
    },  
    bottomEdit:function(ele){
        txt = util.addEle("textarea") 
        ele.appendChild(txt)
        this.div_bottomEdit = txt
    },
    setBottomEdit:function(msg){
        this.div_bottomEdit.innerText = msg
    }
}  