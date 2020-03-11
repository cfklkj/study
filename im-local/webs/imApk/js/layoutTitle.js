 
var layoutTitle = layoutTitle || {};
var titleLayout = layoutTitle.Method = { 
    div_bodys:"bodys",
    div_main:"title",
    main:null,
    title:null,
    back:null,
    show:function(){
        div = util.addEle("div")
        div.setAttribute("class", this.div_main) 
        div.id = this.div_main
        ele = util.getEleById(this.div_bodys)
        ele.appendChild(div) 
        this.main = div 
        this.addBack()
        this.addTitle()
    },
    addTitle:function(){
        h4 = util.addEle("h2")   
        this.main.appendChild(h4)
        this.title = h4
    },
    setTitle:function(msg){
        h4.innerText = msg
    },
    addBack:function(){
        img = util.addEle("img")   
        img.setAttribute("onclick", "clicks.back()")
        this.main.appendChild(img)
        this.back = img 
        this.hideBack()
    },
    showBack:function(){
        this.back.setAttribute("style","")
    },
    hideBack:function(){
        this.back.setAttribute("style","display:none")
    }
}