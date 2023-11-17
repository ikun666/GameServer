package conf

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ikun666/v5/iface"
)

type GlobalConfig struct {
	Server     iface.IServer `json:"server,omitempty"`
	IP         string        `json:"ip,omitempty"`
	Port       int           `json:"port,omitempty"`
	ServerName string        `json:"server_name,omitempty"`

	Version        string `json:"version,omitempty"`
	MaxConn        int    `json:"max_conn,omitempty"`
	MaxPackageSize int    `json:"max_package_size,omitempty"`
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
