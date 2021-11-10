package main

import (
	"github.com/config"
	"github.com/router"
	"github.com/server"
	"github.com/spf13/viper"
)

func main()  {

	cfg, err := loadConfig()
	if err != nil {
		panic("Error in loading configuration!")
	}
	err = serve(cfg)
	if err != nil {
		panic("Error in serve!")
	}
}

const configFileName ="config"

func loadConfig() (*config.Configuration, error) {
	v := viper.New()
	config.SetupViper(v, configFileName)
	return config.New(v)
}


func serve(cfg *config.Configuration) error {
	r, err := router.New()
	if err!= nil {
		return err
	}
	corsRouter := router.SupportCORS(r)
	server.Listen(cfg, router.NoCache{Handler: corsRouter})
	return nil
}