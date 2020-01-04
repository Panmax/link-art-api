package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret  string
	JwtTimeout time.Duration
}

var AppConfig = &App{}

type Server struct {
	Mode         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseConfig = &Database{}

type Qiniu struct {
	Bucket    string
	Prefix    string
	AccessKey string
	SecretKey string
}

var QiniuConfig = &Qiniu{}

type Redis struct {
	Host     string
	Password string
	DB       int
}

var RedisConfig = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("application.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
	mapTo("qiniu", QiniuConfig)
	mapTo("redis", RedisConfig)

	// FIXME
	AppConfig.JwtTimeout = AppConfig.JwtTimeout * time.Second
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
