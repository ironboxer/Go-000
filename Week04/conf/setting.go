package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func LoadAddress(path string) string {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Load ini file %s failed\n", path)
	}
	address := cfg.Section("server").Key("address").String()
	return address
}

func getpwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}

var iniFilePath = getpwd() + "/conf/setting.ini"
var ServerAddr = LoadAddress(iniFilePath)
