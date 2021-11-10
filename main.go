package main

import(
	"github.com/server"
	"github.com/config"
	"github.com/spf13/viper"
)

func main()  {

	cfg, err := loadConfig()
}

const configFileName ="config"

func loadConfig() (*config.Configuration, error) {
	v := viper.New()
	config.SetupViper(v, configFileName)
	return config.New(v)
}

func serve(cfg *config.Configuration)  {
	corsRouter := router.SupportCORS(r)
	server.Listen(cfg, router.NoCache{Handler: corsRouter})
}