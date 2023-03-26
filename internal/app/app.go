package app

import (
	"log"
	"todoApp/internal/apiserver"
	"todoApp/internal/config"
	"todoApp/internal/handler"
	"todoApp/internal/repository"
	"todoApp/internal/service"
)

func Run(configPath string) {
	if configPath == "" {
		configPath = "configs/appConfig.yml"
	}
	conf, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatal("Not correct configPath! err :", err)
	}
	r := repository.NewRepository()
	s := service.NewService(r)
	h := handler.NewHandler(s)

	if err := apiserver.Run(conf, h.InitHandler()); err != nil {
		log.Fatal("Can not start apiserver! err :", err)
	}

}