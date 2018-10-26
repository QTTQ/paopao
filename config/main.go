package config
import (
	 "gopkg.in/yaml.v2"
	"io/ioutil"
)
const (
	EncryptKey= "LKHlhb899Y09olUi321431ew"
	TOKEN_EXPIRE_TIME = 7 * 86400 // 七天有效
)
var Conf ConfStruct
type ConfStruct struct{
	Port int `yaml:"port"`
	EncryptKey string `yaml:"encryptKey"`
	DB struct {
		Host string `yaml:"host"`
		Port uint `yaml:"port"`
		User string `yaml:"user"`
		Name string `yaml:"name"`
		Password string `yaml:"password"`
		Prefix string `yaml:"prefix"`
	} `yaml:"db"`
}

func IniConfig() error{
	source,err:=ioutil.ReadFile("config.yaml")
	if err !=nil{
		panic(err)
	}
	err=yaml.Unmarshal(source,&Conf)
	if err !=nil{
		panic(err)
	}
	return nil
}