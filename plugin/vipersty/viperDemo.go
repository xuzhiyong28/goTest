package vipersty

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)
var gloalConfig *viper.Viper
var databaseConfig *viper.Viper
var applicationConfig *viper.Viper

type DataBaseConfigBean struct {
	Dbtype string
	Host string
	Name string
	Password string
	Port int
	Username string
}

type ApplicationConfigBean struct {
	Domain string
	Host string
	Ishttps bool
	Mode string
	Name string
	Port int
	Readtimeout int

}

var DatabaseConfig = new(DataBaseConfigBean)

var ApplicationConfig = new(ApplicationConfigBean)

func init(){
	gloalConfig = viper.New()
	path, _ := os.Getwd()
	gloalConfig.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	gloalConfig.AddConfigPath(path)
	gloalConfig.SetConfigType("yml") //如果是json的话就换成json就行
	if err := gloalConfig.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	databaseConfig = gloalConfig.Sub("settings.database")
	applicationConfig = gloalConfig.Sub("settings.application")

	DatabaseConfig = InitDatabase(databaseConfig) //加载数据库配置到对象
	InitApplication() //加载application配置

}


func InitDatabase(cfg *viper.Viper) *DataBaseConfigBean{
	return&DataBaseConfigBean{
		Dbtype: cfg.GetString("dbtype"),
		Host: cfg.GetString("host"),
		Name: cfg.GetString("name"),
		Password: cfg.GetString("password"),
		Port: cfg.GetInt("port"),
		Username : cfg.GetString("username"),
	}
}

func InitApplication(){
	err := applicationConfig.Unmarshal(&ApplicationConfig)
	if err != nil{
		fmt.Println(err)
	}
}


func Demo1(){
	fmt.Println(DatabaseConfig.Port)
	fmt.Println(ApplicationConfig.Port)
}


func Demo2(){
	var ostype = runtime.GOOS //获取系统
	filePath := ""
	path, _ := os.Getwd()
	if ostype == "windows"{
		filePath = path + "\\" + "config.json"
	}else if ostype == "linux"{
		filePath = path + "/" + "config.json"
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	//支持从file直接读取
	testViper := viper.New()
	err = testViper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(testViper.GetString("otherset.company"))
	}
}

//配置写入配置文件中
//下面的配置会被写入到config.toml中去 我们可以用这个功能来做配置文件的生成
func Demo3(){
	path, _ := os.Getwd()
	testViper := viper.New()
	testViper.SetConfigName("config")
	testViper.SetConfigType("toml")
	testViper.AddConfigPath(path)
	testViper.Set("app_name", "awesome web")
	testViper.Set("log_level", "DEBUG")
	testViper.Set("mysql.ip", "127.0.0.1")
	testViper.Set("mysql.port", 3306)
	testViper.Set("mysql.user", "root")
	testViper.Set("mysql.password", "123456")
	testViper.Set("mysql.database", "awesome")

	testViper.Set("redis.ip", "127.0.0.1")
	testViper.Set("redis.port", 6381)

	err := testViper.SafeWriteConfig()
	if err != nil {
		fmt.Println(err)
	}
}