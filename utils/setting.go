package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode    string
	HttpPort   string
	JwyKey     string
	Db         string
	DbHost     string
	DbPort     string
	DuUser     string
	DbPassWord string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)

	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {

	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwyKey = file.Section("server").Key("JwyKey").MustString("8989898")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DuUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPasserWord").MustString("1234abc")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("wIV0i2KyKbNqlC7051ZiNZ4_ARTg3i5pIQd9DHAo")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("soRxx-l3_Mq2H5cIQXQkM1UhBw4EIZET5KLh6EQA")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("ginblogimg")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("http://rp52kl8q0.hn-bkt.clouddn.com/")
}
