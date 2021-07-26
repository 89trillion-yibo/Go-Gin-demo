package utils

import (
	"github.com/tietang/props/ini"
	"strconv"
)

//解析app.ini文件，并调出端口参数
func RandIni() string {
	conf := ini.NewIniFileConfigSource("./config/app.ini")
	tmp := conf.GetIntDefault("server.HttpPort",3306)
	port := ":"+strconv.Itoa(tmp)
	return port
}
