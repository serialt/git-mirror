package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/serialt/git-mirror/config"
	"github.com/serialt/git-mirror/service"
	"github.com/serialt/sugar"
)

func env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}

var (
	APPName    = "git-audit"
	Maintainer = "tserialt@gmail.com"
	APPVersion = "v0.2"
	BuildTime  = "200601021504"
	GitCommit  = "ccccccccccccccc"
	appVersion bool
)

func init() {
	// 初始化app信息
	config.APPName = APPName
	config.Maintainer = Maintainer
	config.APPVersion = APPVersion
	config.BuildTime = BuildTime
	config.GitCommit = GitCommit

	flag.BoolVarP(&appVersion, "version", "v", false, "Display build and version msg.")
	flag.StringVarP(&config.ConfigPath, "cfgFile", "c", env("CONFIG", config.ConfigPath), "Path to config yaml file.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")

	flag.CommandLine.SortFlags = false
	flag.Parse()
	// register global var
}

func GitInit() {
	config.LoadConfig(config.ConfigPath)
	mylg := &sugar.Logger{
		LogLevel:      config.Config.GitLog.LogLevel,
		LogFile:       config.Config.GitLog.LogFile,
		LogType:       config.Config.GitLog.LogType,
		LogMaxSize:    50,
		LogMaxBackups: 3,
		LogMaxAge:     365,
		LogCompress:   true,
	}
	config.Logger = mylg.NewMyLogger()
	config.LogSugar = config.Logger.Sugar()
	service.LogSugar = config.Logger.Sugar()

	// mydb := &sugar.Database{
	// 	Type:     config.Config.Database.Type,
	// 	Addr:     config.Config.Database.Addr,
	// 	Port:     config.Config.Database.Port,
	// 	DBName:   config.Config.Database.DBName,
	// 	Username: config.Config.Database.Username,
	// 	Password: config.Config.Database.Password,
	// }
	// config.DB = mydb.NewDBConnect(config.Logger)
}

func main() {

	if appVersion {
		fmt.Printf("APPName: %v\n Maintainer: %v\n Version: %v\n BuildTime: %v\n GitCommit: %v\n GoVersion: %v\n OS/Arch: %v\n",
			config.APPName,
			config.Maintainer,
			config.APPVersion,
			config.BuildTime,
			config.GitCommit,
			config.GOVERSION,
			config.GOOSARCH)
		return
	}
	GitInit()

	// pkg.Sugar.Info(config.LogFile)
	client := &service.GiteeClient{
		AccessToken: "4909e77545e6d5f891ce9a8e9b6697ea",
	}
	_, _ = client.GiteeCreateRepo("go-gitee5", false)
	// fmt.Printf("resp: %v\n", resp)
	// fmt.Printf("err: %v\n", err)
	service.CloneRepo("https://github.com/serialt/sugar", "/Users/serialt/Desktop/flkj/tmp/gitee")
}
