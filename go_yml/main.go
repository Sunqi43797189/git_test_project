package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)
type Conf struct {
	RunMode string
	App App
	Server Server
	Database Database
}

type App struct {
	PageSize int
	JwtSecret string
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func main(){
	conf := Conf{}
	data, err := ioutil.ReadFile("app.yml")
	if err != nil {
		fmt.Println("文件读取失败")
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("yaml文件转换失败")
	}
	fmt.Printf("%T", conf.App.ImageAllowExts)
}