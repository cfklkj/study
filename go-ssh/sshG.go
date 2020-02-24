package main

import (
	"fmt"

	"./cmds"
	"./ssh"
)

func main() {
	ngx := cmds.NewCmds()
	sh := ssh.NewCSSH()
	if _, err := sh.Connect("flylkl", "123", "fly.lkl:22"); err != nil {
		fmt.Println("ssh-err", err)
		return
	}
	sh.PrintMsg = ngx.PrintMsg
	ngx.Run = sh.Run
	// ngx.Cmd_checkUfw()
	// ngx.Cmd_checkUnzip()
	// ngx.Cmd_checkCertbot()
	// ngx.Cmd_checkNginx()
	//--https
	//ngx.Cmd_makeCert("www.baidu.com")
	//ngx.Cmd_nginxReStart()
	//ngx.Cmd_downFile("127.0.0.1:10124/download/readme.zip", "/tmp/abc.7z", true, ngx.Zip_unzip)
	//--port mgr
	//ngx.Cmd_ufwAllow(10012)
	//ngx.Cmd_ufwDelAllow(10012)
	//--android
	//ngx.Cmd_checkJDK()
	//ngx.Cmd_checkSDK("/tmp/android-sdk-linux")
	//ngx.Cmd_checkSDKbuildTools("/tmp/android-sdk-linux", "28.0.3")
	//ngx.Cmd_checkGradle("/tmp/gradle-6.2")
	ngx.Cmd_gradleBuild("/tmp/gradle-6.2", "/tmp/android-sdk-linux", "/tmp/demo")
	sh.Clear()
}
