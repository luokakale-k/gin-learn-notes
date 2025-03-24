package config

import "github.com/spf13/viper"

type App struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type MySQLInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type Logger struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

// Config 总配置结构体（对应 config.yaml）
type Config struct {
	App    App       `mapstructure:"app"`
	MySQL  MySQLInfo `mapstructure:"mysql"`
	Logger Logger    `mapstructure:"logger"`
	Redis  Redis     `mapstructure:"redis"`
}

var Conf *Config

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config") // 不带扩展名
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 当前目录

	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败: " + err.Error())
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic("解析配置文件失败: " + err.Error())
	}

	println("配置加载成功")
}
