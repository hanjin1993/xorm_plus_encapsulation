package utils

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

type Config struct {
	Module string //模块
	Filed  string //查询某一个配置
}

const CONFIG_PATH string = "./config/config.ini"

var config *goconfig.ConfigFile

func init() {
	configFile, err := goconfig.LoadConfigFile(CONFIG_PATH)
	fmt.Println(configFile)
	if err == nil {
		config = configFile
	}
}

func (search *Config) Config() (res string) {
	//list := config.GetKeyList(search.Module)
	if config == nil {
		return ""
	}
	value, err := config.GetValue(search.Module, search.Filed)
	if err != nil {
		return ""
	}
	return value
}

func (search *Config) ConfigList() (res []string) {
	res = config.GetKeyList(search.Module)
	return
}
