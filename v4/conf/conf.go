package conf

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ikun666/v4/iface"
)

type GlobalConfig struct {
	Server     iface.IServer
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	ServerName string `json:"serverName"`

	Version        string `json:"version"`
	MaxConn        int    `json:"maxConn"`
	MaxPackageSize int    `json:"maxPackageSize"`
}

var GConfig *GlobalConfig //全局配置
func Init(path string) error {
	GConfig = &GlobalConfig{}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("read conf err:%v", err)
		return fmt.Errorf("read conf err:%v", err)
	}
	err = json.Unmarshal(data, GConfig)
	if err != nil {
		fmt.Printf("unmarshal conf err:%v", err)
		return fmt.Errorf("unmarshal conf err:%v", err)
	}
	return nil
}
