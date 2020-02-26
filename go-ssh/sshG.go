package main

import (
	"fmt"

	"./cmds"
	"./ssh"
)

func main() {
	ngx := cmds.NewCmds()
	sh := ssh.NewCSSH()
	rst := false
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
	//rst = ngx.Cmd_makeCert("im.guiruntang.club")
	//rst = ngx.Cmd_nginxReStart()
	//rst = ngx.Cmd_downFile("127.0.0.1:10124/download/readme.zip", "/tmp/abc.7z", true, ngx.Zip_unzip)
	//--port mgr
	//rst = ngx.Cmd_ufwAllow(10012)
	//rst = ngx.Cmd_ufwDelAllow(10012)
	//--android
	//ngx.Cmd_checkJDK()
	//ngx.Cmd_checkSDK("/tmp/android-sdk-linux")
	//ngx.Cmd_checkSDKbuildTools("/tmp/android-sdk-linux", "28.0.3")
	//ngx.Cmd_checkGradle("/tmp/gradle-6.2")
	//rst = ngx.Cmd_gradleBuild("/tmp/gradle-6.2", "/tmp/android-sdk-linux", "/tmp/demo")
	//--fail2ban
	// ngx.Cmd_checkFail2ban()
	// rst =  ngx.Cmd_restartFail2ban()
	//--kill port's pid
	//rst = ngx.Cmd_killRunPortPID(80)
	//---scp
	//rst = sh.SCPupFile("sshG.go", "/tmp/")
	//rst = sh.SCPDownFile("/tmp/sshG.go", "d:/")
	//---file mod
	//ngx.Cmd_showMode("/tmp")
	//rst = ngx.Cmd_chown("im", "/tmp/java_url3.7z")
	rst = ngx.Cmd_chmodR("755", "/tmp/java_url3.7z")
	//	rst = ngx.Cmd_chown("root:root", "/tmp/java_url3.7z")
	fmt.Println("rst", rst)
	sh.Clear()
}
