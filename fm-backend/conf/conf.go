package conf

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"testing"
)

var GConfig *viper.Viper

func init()  {
	confString := flag.String("c", "./config", "file path to read from")
	testing.Init()
	flag.Parse()
	Init(*confString)
}

func Init(path string) {
	path = path + "/"
	fmt.Println("Loading configuration logics...")
	GConfig = initConfig(path)
	go dynamicConfig()
}

func initConfig(path string) *viper.Viper {
	GlobalConfig := viper.New()
	GlobalConfig.SetConfigName("config")
	GlobalConfig.AddConfigPath(path)
	GlobalConfig.SetConfigType("toml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to get the configuration.")
	}
	return GlobalConfig
}

func dynamicConfig() {
	GConfig.WatchConfig()
	GConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})
}