var imHttp = imHttp || {};
var doHttp = imHttp.Method = {  
    checkData: function(res){ 
        data = JSON.parse(res)
        if (data == null || data.Code != 200) {        
            alert(data, res) 
            return null
        }else{ 
            return data
        }      
    },
    request: function(GOrP, URL, data, callBack){ 
        var xhr = new XMLHttpRequest();
        xhr.open(GOrP, URL, true); 
        xhr.send(data);  
        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4) { // 读取完成
                if (xhr.status == 200) { 
                     data = doHttp.checkData(xhr.responseText)
                     if (data != null){ 
                        callBack(data)
                     } 
                     return
                }
            }             
            if(xhr.status == 500){
                if(xhr.responseText == "token timeout")
                {
                    window.location.reload();
                }else
                    alert(xhr.responseText);
            }else{
                console.debug("request error", xhr.status);
            }
        }
    }, //get prolist
    getLoginInfo: function(url, userName, password, resFunc) {
       // var data = {Url:url, Data:{name:userName, password:password}}    //application/json;charset=utf-8
        body = 'username=' +userName+ "&password=" +password + '';
        var data =  {"Url":url, "Data":body, "ContentType":"application/x-www-form-urlencoded;"}   
        this.request("post","/prox", JSON.stringify(data), resFunc)  
    },  
}
 