document.write('<script  type="text/javascript" src="js/util.js"></script>')  
document.write('<script  type="text/javascript" src="js/mem.js"></script>')  
document.write('<script  type="text/javascript" src="js/login.js"></script>')  
document.write('<script  type="text/javascript" src="js/logOut.js"></script>')  
document.write('<script  type="text/javascript" src="js/webHttp.js"></script>')  
document.write('<script  type="text/javascript" src="js/websocket.js"></script>') 
document.write('<script  type="text/javascript" src="js/websocket_local.js"></script>') 
document.write('<script  type="text/javascript" src="js/define.js"></script>')  
document.write('<script  type="text/javascript" src="js/chat.js"></script>')  
document.write('<script  type="text/javascript" src="js/chatLeft.js"></script>')  
document.write('<script  type="text/javascript" src="js/chatRight.js"></script>')  
document.write('<script  type="text/javascript" src="js/imgDlg.js"></script>')   
window.onload = main 
window.setTimeout("timeTick()",1000*90);

function main(){
    localSocket.connect("ws://127.0.0.1:10023/v1/tzj")
    outDiv.showLogin();
}

function timeTick(){
    imSocket.hearMsg();
    window.setTimeout("timeTick()",1000*90);
}