package tests

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T)  {

	c := viper.New()
	//设置配置文件的名字
	c.SetConfigName("../configs")
	c.AddConfigPath("conf")
	//设置配置文件类型
	c.SetConfigType("yaml")
	c.ReadInConfig()

	a := c.AllSettings()

	fmt.Println(a)

}