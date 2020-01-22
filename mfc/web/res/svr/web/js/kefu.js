document.write('<script  type="text/javascript" src="js/util.js"></script>')  
document.write('<script  type="text/javascript" src="js/mem.js"></script>')  
document.write('<script  type="text/javascript" src="js/login.js"></script>')  
document.write('<script  type="text/javascript" src="js/webHttp.js"></script>')  
document.write('<script  type="text/javascript" src="js/websocket.js"></script>') 
document.write('<script  type="text/javascript" src="js/websocket_local.js"></script>') 
document.write('<script  type="text/javascript" src="js/define.js"></script>')  
document.write('<script  type="text/javascript" src="js/chats.js"></script>')  
document.write('<script  type="text/javascript" src="js/imgDlg.js"></script>')   
window.onload = main 
window.setTimeout("timeTick()",1000*90);

function main(){
    localSocket.connect("ws://127.0.0.1:10023/v1/tzj")
    inDiv.showLogin();
}

function timeTick(){
    imSocket.hearMsg();
    window.setTimeout("timeTick()",1000*90);
}