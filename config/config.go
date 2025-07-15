package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-blog/util"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type ServerConfig struct {
	Name        string
	Host        string
	Port        int
	LocalIp     string
	MsgSize     int
	CheckSig    bool
	CheckSecret string
	LocalMac    string
}

var svrCfg ServerConfig
var appViper *viper.Viper

func InitServerConfig() {
	var err error
	//v, errs := initConfigFromRemote()
	appViper, err = initConfigFromLocal()
	if err != nil {
		panic(err)
		return
	}
	svrCfg.Host = appViper.GetString("server.host")
	svrCfg.Port = appViper.GetInt("server.port")
	svrCfg.Name = appViper.GetString("server.name")
	svrCfg.MsgSize = appViper.GetInt("server.msgSize")
	svrCfg.CheckSig = appViper.GetBool("server.checkSig")
	svrCfg.CheckSecret = appViper.GetString("server.checkSecret")
	//AppCfg.EtcdAddr = appViper.GetString("server.etcdAddr")
	localIp, err := util.GetOutBoundIP()
	if err != nil {
		log.Printf("InitServerConfig err:%v", err)
	}
	svrCfg.LocalIp = localIp
	if len(svrCfg.Host) > 0 {
		svrCfg.LocalIp = svrCfg.Host
	}
	svrCfg.LocalMac = util.GetMacByIp(svrCfg.LocalIp)
	_ = os.Setenv("server", svrCfg.Name)
}

func initConfigFromRemote() (*viper.Viper, error) {
	v := viper.New()
	// 远程配置 etcd3而不是etcd
	v.AddRemoteProvider("etcd3", "http://localhost:2379", "/main.yaml")
	//v.SetConfigType("json")
	v.SetConfigFile("main.yaml")
	v.SetConfigType("yaml")

	if err := v.ReadRemoteConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		log.Printf("InitConfigFromRemote -> %+v\n", err)
		return nil, err
	}
	return v, nil
}

var configPath string

func SetConfigPath(path string) {
	configPath = path
}

func dynamicName() string {
	env := strings.ToLower(os.Getenv("ENV"))
	configName := "main.yaml"
	if env != "" {
		configName = "main_" + env + ".yaml"
	}
	return configName
}

func initConfigFromLocal() (*viper.Viper, error) {
	v := viper.New()
	if configPath == "" {
		v.SetConfigName(dynamicName())
		v.SetConfigType("yaml")
		filePath, _ := os.Getwd()
		log.Printf("initConfigFromLocal filePath1:%s\n", filePath)
		filePath = filepath.ToSlash(filePath) + "/config"
		log.Printf("initConfigFromLocal filePath2:%s\n", filePath)
		v.AddConfigPath(filePath)
	} else {
		v.SetConfigFile(configPath)
	}
	//_, conf.filePath, _, _ = runtime.Caller(0)
	//conf.viper.SetConfigType("yml")
	//conf.viper.AddConfigPath(path.Dir(conf.filePath))

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("initConfigFromLocal 读取配置文件出错:", err)
		return nil, err
	} else {
		log.Printf("initConfigFromLocal %+v\n", v.AllSettings())
	}

	return v, nil
}

func GetViper() *viper.Viper {
	return appViper
}

func GetServerConfig() ServerConfig {
	return svrCfg
}

func Initialize() {
	InitServerConfig()
}

func UnmarshalKey(key string, rawVal interface{}) error {
	return appViper.UnmarshalKey(key, rawVal)
}

func Unmarshal(rawVal interface{}) error {
	return appViper.Unmarshal(rawVal)
}

func UpdateLocalIp(ip string) {
	svrCfg.LocalIp = ip
}
