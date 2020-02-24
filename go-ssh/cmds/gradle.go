package cmds

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (c *Cmds) Cmd_environmemt(evtStr string) { // PATH=$PATH:xx
	sh := "export " + evtStr
	c.Run(str_environment, sh)
}
func (c *Cmds) Cmd_checkJDK() {
	c.BindCheck(str_jdkCheck, c.checkJDK)
	sh := "javac -version"
	c.Run(str_jdkCheck, sh)
}
func (c *Cmds) cmd_installJDK() {
	c.BindCheck(str_jdk, c.checkJDK)
	sh := "apt install default-jdk -y"
	c.Run(str_jdk, sh)

}

func (c *Cmds) Cmd_checkSDK(path string) {
	c.status[str_sdkCheck] = path
	c.BindCheck(str_sdkCheck, c.checkAndroidSdk)
	sh := "export ANDROID_HOME=" + path +
		"\nexport PATH=$PATH:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools" +
		"\nandroid -h"
	c.Run(str_sdkCheck, sh)
}

func (c *Cmds) cmd_checkSDKtools(path string) {
	c.status[str_sdkMgrCheck] = path
	c.BindCheck(str_sdkMgrCheck, c.checkAndroidSdkMgr)
	sh := "ls " + path + "/tools/bin/sdkmanager"
	c.Run(str_sdkMgrCheck, sh)
}

func (c *Cmds) Cmd_checkSDKbuildTools(path, version string) {
	c.status[str_sdkBuildToolsCheck] = BuildToolsInfo{path, version}
	c.BindCheck(str_sdkBuildToolsCheck, c.checkAndroidSdkBuildTools)
	sh := "ls " + path + "/build-tools/" + version
	c.Run(str_sdkBuildToolsCheck, sh)
}
func (c *Cmds) cmd_installSDKbuildTools() {
	info := c.status[str_sdkBuildToolsCheck].(BuildToolsInfo)
	path := info.Path
	version := info.Version
	sh := "export ANDROID_HOME=" + path +
		"\nexport PATH=$PATH:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools" +
		"\nsdkmanager \"build-tools;" + version + "\""
	c.Run(str_sdkBuildTools, sh)
}

func (c *Cmds) Cmd_checkGradle(path string) {
	c.status[str_gradleCheck] = path
	c.BindCheck(str_gradleCheck, c.checkGradle)
	sh := "export GRADLE_HOME=" + path +
		"\nexport PATH=$PATH:$GRADLE_HOME/bin" +
		"\ngradle -v"
	c.Run(str_gradleCheck, sh)
}
func (c *Cmds) Cmd_gradleBuild(pathGradle, pathAndroid, pathPro string) {
	c.BindCheck(str_gradleBuild, c.checkBuild)
	sh := "export GRADLE_HOME=" + pathGradle +
		"\nexport ANDROID_HOME=" + pathAndroid +
		"\nexport PATH=$PATH:$GRADLE_HOME/bin:$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools" +
		"\ncd " + pathPro +
		"\ngradle assembleRelease"
	fmt.Println("正在编译。。。")
	c.Run(str_gradleBuild, sh)
}

//--------check
func (c *Cmds) checkBuild(msg string) {
	if strings.Contains(msg, "BUILD SUCCESSFUL") {
		fmt.Println("gradle", "BUILD SUCCESSFUL")
	} else {
		fmt.Println("gradle", "BUILD 失败")
	}
}

func (c *Cmds) checkJDK(msg string) {
	if strings.Contains(msg, str_err) {
		fmt.Println("正在安装jdk")
		c.cmd_installJDK()
	} else {
		fmt.Println("已安装jdk")
	}
}

func (c *Cmds) checkJDKinstall(msg string) {
	if strings.Contains(msg, "Setting up default-jdk") {
		fmt.Println("安装jdk完成")
	} else {
		fmt.Println("安装jdk失败")
	}
}

func (c *Cmds) checkGradle(msg string) {
	if strings.Contains(msg, str_err) {
		fmt.Println("正在安装gradle")
		dir, fileName := filepath.Split(c.status[str_gradleCheck].(string))
		url := "https://services.gradle.org/distributions/" + fileName + "-bin.zip"
		path := dir + fileName + "-bin.zip"
		c.Cmd_downFile(url, path, true, Zip_unzip)
	} else {
		fmt.Println("已安装gradle")
	}
}

func (c *Cmds) checkAndroidSdk(msg string) {
	if !strings.Contains(msg, " list ") {
		fmt.Println("正在安装sdk")
		dir, fileName := filepath.Split(c.status[str_sdkCheck].(string))
		url := "https://dl.google.com/android/android-sdk_r24.4.1-linux.tgz"
		path := dir + fileName + ".tgz"
		c.Cmd_downFile(url, path, true, Zip_tar)
	} else {
		fmt.Println("已安装sdk")
	}
	c.cmd_checkSDKtools(c.status[str_sdkCheck].(string))
}

func (c *Cmds) checkAndroidSdkMgr(msg string) {
	if strings.Contains(msg, "No such file") {
		fmt.Println("正在安装sdk-tools")
		dir := c.status[str_sdkMgrCheck].(string)
		url := "https://dl.google.com/android/repository/sdk-tools-linux-4333796.zip"
		path := dir + "/sdk-tools.zip"
		c.Cmd_downFile(url, path, true, Zip_unzip)
	} else {
		fmt.Println("已安装sdk-tools")
	}
}
func (c *Cmds) checkAndroidSdkBuildTools(msg string) {
	if strings.Contains(msg, "No such file") {
		fmt.Println("正在安装build-tools")
		c.cmd_installSDKbuildTools()
	} else {
		fmt.Println("已安装build-tools")
	}
}
