package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var config = new(Config)

type Config struct {
	 App struct {
	 	Name string `toml:"name"`
	 	Debug bool  `toml:"debug"`
	 	Addr string `toml:"addr"`
	 	Port int `toml:"port"`
	 	Env  string `toml:"env"`
	 } `toml:"app"`
	MySQL struct {
		Db struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"db"`
		Gorm struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"gorm"`
	} `toml:"mysql"`
	Redis struct {
		Cache struct{
			Addr         string `toml:"addr"`
			Pass         string `toml:"pass"`
			Db           int    `toml:"db"`
			MaxRetries   int    `toml:"maxRetries"`
			PoolSize     int    `toml:"poolSize"`
			MinIdleConns int    `toml:"minIdleConns"`
		} `toml:"cache"`
	} `toml:"redis"`
}


func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) { //fsnotify 监控配置文件变化，如果变化重新赋值
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}
