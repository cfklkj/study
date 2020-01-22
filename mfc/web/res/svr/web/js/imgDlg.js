var layoutImgDlg = layoutImgDlg || {};
var imgDlg = layoutImgDlg.Method = {   
    openDlg:function(){   
        ele = util.getEleById("file-button")
        ele.click();
    },
    changes:function(file){
        if (!file.files || !file.files[0]) {
            return;
        }
        var reader = new FileReader();
        reader.onload = function (evt) { 
            imgDlg.sumitImageFile(evt.target.result)
        }
        reader.readAsDataURL(file.files[0]); 
    },
    doneSubmit:function(res){  
        body = JSON.parse(res)
        if (body.code == 200){
           data = memHash.get(imDefine.mem_loginData)   
           chat.send(actRecords.get(imDefine.act_selectContent),data.image_url + "/" + body.data, imDefine.chat_img) 
        }      
    },
    sumitImageFile:function(base64Codes){   
        var formData = new FormData();    
        formData.append("image",imgDlg.convertBase64UrlToBlob(base64Codes));  //append函数的第一个参数是后台获取数据的参数名,和html标签的input的name属性功能相同  
        data = memHash.get(imDefine.mem_loginData) 
        doHttp.requestProx("POST",data.upUrl,formData, imgDlg.doneSubmit) 
    },
    //convertBase64UrlToBlob函数是将base64编码转换为Blob  
    convertBase64UrlToBlob:function(urlData){  

        var bytes=window.atob(urlData.split(',')[1]);        //去掉url的头，并转换为byte  
    
        //处理异常,将ascii码小于0的转换为大于0  
        var ab = new ArrayBuffer(bytes.length);  
        var ia = new Uint8Array(ab);  
        for (var i = 0; i < bytes.length; i++) {  
            ia[i] = bytes.charCodeAt(i);  
        }   
        return new Blob( [ab] , {type : 'image/png'});  
    } 
} 