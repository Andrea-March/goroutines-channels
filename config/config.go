package config

import (
	"errors"
	"fmt"
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
		v.AddConfigPath(".")
	}

	v.SetDefault("port", 8000)
	v.SetDefault("host", "")

	//env variables
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.SetEnvPrefix("PBS")
	v.AutomaticEnv()
	v.ReadInConfig()
}

func New(v *viper.Viper)  (*Configuration, error) {
	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal app config: %v", err)
	}
	if err := c.validate(); err != nil{
		return &c, errors.New(err.Error())
	}
	
	return &c, nil
}