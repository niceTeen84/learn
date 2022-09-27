package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	ApiVersion *string
	Server     *ServerInfo
}

type ServerInfo struct {
	Bind uint16
	Addr *string
}

func main() {
	viper.AddRemoteProvider("consul", "localhost:8500", "renbw")
	viper.SetConfigType("yaml")
	// viper.SetConfigName("renbw")
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal("read remote consul failed ", err.Error())
	}
	// viper 在读取 yaml 文件时可以使用 json path 的方式
	fmt.Println(viper.GetInt32("server.bind"))
	// viper 可以将结构化的 yaml 文件序列化为 struct
	ac := AppConfig{}
	viper.Unmarshal(&ac)
}
