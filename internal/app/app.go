package app

import (
	"awesomeProject/internal/apiserver"
	"awesomeProject/internal/config"
	"log"
)

func Run(configPath string) {
	if configPath == "" {
		configPath = "configs/appConfig.yml"
	}
	conf, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatal("Not correct configPath! err :", err)
	}

	if err := apiserver.Run(conf); err != nil {
		log.Fatal("Can not start apiserver! err :", err)
	}

}
