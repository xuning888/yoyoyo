package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var RuntimeViper *viper.Viper

func Init() {
	RuntimeViper = viper.New()
	RuntimeViper.SetConfigType("toml")
	RuntimeViper.SetConfigName("cfg")
	RuntimeViper.AddConfigPath("/etc/yoyoyo/")
	RuntimeViper.AddConfigPath("./config/")
	if err := RuntimeViper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	RuntimeViper.WatchConfig()
	RuntimeViper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed:%s", e.Name)
	})

	// ollama server
	DefaultConfig.LlmModel = RuntimeViper.GetString("ollama.model")
	DefaultConfig.OllmServerUrl = RuntimeViper.GetString("ollama.serverUrl")

	// http server config
	DefaultConfig.ServerPort = RuntimeViper.GetString("server.port")

	// redis config
	DefaultConfig.Redis.Model = RuntimeViper.GetString("redis.model")
	DefaultConfig.Redis.Addr = RuntimeViper.GetString("redis.addr")
	DefaultConfig.Redis.Addrs = RuntimeViper.GetStringSlice("redis.addrs")
	DefaultConfig.Redis.Addrs = RuntimeViper.GetStringSlice("redis.sentinel-addrs")
	DefaultConfig.Redis.MasterName = RuntimeViper.GetString("redis.master-name")
	DefaultConfig.Redis.Password = RuntimeViper.GetString("redis.password")
	DefaultConfig.Redis.DB = RuntimeViper.GetInt("redis.DB")

	// mysql config
	DefaultConfig.MySQL.DSN = RuntimeViper.GetString("mysql.dsn")
	log.Printf("init config success: %v\n", DefaultConfig)
}
