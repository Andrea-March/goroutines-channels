package config

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Port int `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

func (c Configuration) validate() error {
	if c.Port == 0 {
		return errors.New("Error in port configuration")
	}
	return nil
}

func SetupViper(v *viper.Viper, filename string)  {
	if filename != "" {
		v.SetConfigName(filename)
		v.AddConfigPath("./config/")
	}

	v.SetDefault("port", 8000)
	v.SetDefault("host", "")

	//env variables
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.SetEnvPrefix("PBS")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		glog.Warning(fmt.Sprintf("Error in reading viper config.yml: %s", err.Error()))
	}
}

func New(v *viper.Viper)  (*Configuration, error) {
	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal app config.yml: %v", err)
	}
	if err := c.validate(); err != nil{
		return &c, errors.New(err.Error())
	}
	
	return &c, nil
}