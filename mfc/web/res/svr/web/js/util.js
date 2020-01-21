var flyUtil =  flyUtil || {}; 

var util = flyUtil.commonMethod  = {
    addDiv: function(id){   
      return "<div id="+id +"></div>" 
    },
    //ele
    getEleByTag: function(tag)
    {
        return document.getElementsByTagName(tag)
    },
    //ele
    getEleById: function(id)
    {
        return document.getElementById(id)
    },
    //value
    getEleValue:function(id)
    {
        ele = this.getEleById(id)
        return ele.innerHTML
    },
    //set value
    setEleName :function(id, name)
    {
        ele = this.getEleById(id)
        ele.innerHTML = name
    },
    //添加css
    addClass: function(ele, classname)
    {   
        var oldClass = ele.className;
        var pattern = new RegExp('(^|\\s)' + classname + '(\\s|$)');      
        if (!pattern.test(oldClass))
        {         
                ele.className += ' ' + classname;   
        } 
    },    
    //删除css
    delClass: function(ele, classname){ 
      var oldClass = ele.className;
      var pattern = new RegExp('(^|\\s)' + classname + '(\\s|$)');      
      if (!pattern.test(oldClass)) {         
            ele.className = ele.className.replace(pattern, ' ');
        }else {
            if (ele.className != null) { 
                ele.className  = null
            }
        } 
    },
    dropClass:function(id, parent){ 
        ele =  util.getEleById(id)
        if (ele == null) {
            return
        }
        if (parent)
        { 
            if (ele.parentElement.className == null){
                return
            } 
             ele.parentElement.className = null
             return
        }
        if (ele.className == null)
        {
            return
        }
        ele.className = null
    },
    //添加元素
    addElement: function(ele, html){ 
     ele.innerHTML += html 
   },
   //清空
   clearElement:function(ele){
       ele.innerHTML = "";
   },
   //滚动
   addScrollEvt:function(id, callBack){
    ele = util.getEleById(id)
    ele.setAttribute("onscroll",callBack); 
   },
   getScrollTop:function(){
        var scrollTop = 0, bodyScrollTop = 0, documentScrollTop = 0;
        if(document.body){
            bodyScrollTop = document.body.scrollTop;
        }
        if(document.documentElement){
            documentScrollTop = document.documentElement.scrollTop;
        }
        scrollTop = (bodyScrollTop - documentScrollTop > 0) ? bodyScrollTop : documentScrollTop;
        return scrollTop;
    },
    addEvt:function(id, event, callBack){
        ele = util.getEleById(id)
        ele.setAttribute(event,callBack);   
    },
   //添加方法
    addEvent: function(id, callBack){ 
        ele = util.getEleById(id)
        ele.setAttribute("onclick",callBack);  
      },
      addMouseOut:function(id, callBack){
          ele = util.getEleById(id)
          ele.setAttribute("onmouseout",callBack);   
      },
    addMouseOver:function(id, callBack){
        ele = util.getEleById(id)
        ele.setAttribute("onmouseover",callBack);   
    },
    addMouseDown:function(id, callBack){
        ele = util.getEleById(id)
        ele.setAttribute("onmousedown",callBack);   
    },
    timeHms:function(){
        var now = new Date();
        var hour = now.getHours();//得到小时
        var minu = now.getMinutes();//得到分钟
        var sec = now.getSeconds();//得到秒
        var time = "";
        if (hour < 10) hour = "0" + hour;
        if (minu < 10) minu = "0" + minu;
        if (sec < 10) sec = "0" + sec;
        time = hour + ":" + minu + ":" + sec;
        return time
    }

}   