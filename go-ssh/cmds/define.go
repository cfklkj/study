package cmds

const (
	Zip_unzip = 1
	Zip_tar   = 2
)
const (
	str_nginx              = "nginx"
	str_nginxCheck         = "nginxCheck"
	str_unzip              = "unzip"
	str_unzipCheck         = "unzipCheck"
	str_wget               = "wget"
	str_unzipFile          = "unzipFile"
	str_nginxReStart       = "nginxReStart"
	str_nginxReload        = "nginxReload"
	str_certbot            = "certbot"
	str_certbotCheck       = "certbotCheck"
	str_certNew            = "certNew"
	str_certRenew          = "certRenew"
	str_certls             = "certLs"
	str_ufwCheck           = "ufwCheck"
	str_ufw                = "ufw"
	str_ufwAllow           = "ufwAllow"
	str_ufwDelAllow        = "ufwDelAllow"
	str_ufwEnable          = "ufwEnable"
	str_environment        = "environment"
	str_jdkCheck           = "jdkCheck"
	str_jdk                = "jdk"
	str_sdkCheck           = "sdkCheck"
	str_sdk                = "sdk"
	str_sdkMgrCheck        = "sdkMgrCheck"
	str_sdkBuildToolsCheck = "sdkBuildToolsCheck"
	str_sdkBuildTools      = "sdkBuildTools"
	str_gradleCheck        = "gradleCheck"
	str_gradleBuild        = "gradleBuild"
	str_fail2banCheck      = "fail2banCheck"
	str_fail2ban           = "fail2ban"
	str_fail2banRestart    = "fail2banRestart"
	str_lsofPort           = "lsofPort"
	str_lsofPid            = "lsofPid"
	str_kill               = "kill"
)

const (
	str_null         = "status 0"
	str_err          = "exited with status"
	str_nofidCmd     = "status 127"
	str_nginxConfErr = "status 1"
)

type DownInfo struct {
	Url      string
	KeepPath string
	IsZip    bool
	ZipType  int
}

type BuildToolsInfo struct {
	Path    string
	Version string
}
