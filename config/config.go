package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"

	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// cli里的配置参数，使用类型类似firewalld
var (
	// 版本信息
	APPName    = ""
	Maintainer = ""
	APPVersion = ""
	BuildTime  = ""
	GitCommit  = ""
	GOVERSION  = runtime.Version()
	GOOSARCH   = runtime.GOOS + "/" + runtime.GOARCH
	// 其他配置文件
	ConfigPath = ""

	Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	DB       *gorm.DB

	WG sync.WaitGroup
)

type GitLog struct {
	LogLevel string `yaml:"logLevel"` // 日志级别
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
	LogType  string `yaml:"logType"`  // 日志类型，支持 txt 和 json ，默认txt
	// LogMaxSize    int    //单位M
	// LogMaxBackups int    // 日志文件保留个数
	// LogMaxAge     int    // 单位天
	// LogCompress   bool   // 压缩轮转的日志
}

type MyConfig struct {
	GitLog     GitLog   `yaml:"gitLog"`
	GiteeToken string   `yaml:"giteeToken"`
	GithubRepo []string `yaml:"githubRepo"`
}

var Config *MyConfig

// 判断文件目录否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func LoadConfig(filepath string) {
	if filepath == "" {
		dir, _ := homedir.Dir()
		filepath = fmt.Sprintf("%v/%v", dir, ".git-mirror.yaml")
	}
	filepath, err := homedir.Expand(filepath)
	if err != nil {
		fmt.Printf("get config file failed: %v\n", err)
	}
	if !Exists(filepath) {
		fmt.Printf("file not exist, please check it: %v\n", filepath)
		os.Exit(8)
	}
	config, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("pread config failed, please check the path: %v , err: %v\n", filepath, err)
	}
	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		fmt.Printf("Unmarshal to struct, err: %v", err)
	}
	// fmt.Printf("LoadConfig: %v\n", Config)
	fmt.Printf("config path: %v\n", filepath)
}
