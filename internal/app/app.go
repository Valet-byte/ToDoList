package app

import (
	"github.com/sirupsen/logrus"
	"todoApp/internal/apiserver"
	"todoApp/internal/config"
	"todoApp/internal/handler"
	"todoApp/internal/repository"
	"todoApp/internal/repository/db"
	"todoApp/internal/service"
)

func Run(configPath string) {
	if configPath == "" {
		configPath = "configs/appConfig.yml"
	}
	conf, err := config.NewConfig(configPath)

	if err != nil {
		logrus.Fatal("Not correct configPath! err :", err)
	}

	var postgresDB = db.NewPostgresDB(conf)
	r := repository.NewRepository(postgresDB)
	s := service.NewService(r, conf.Jwt.JwtKye, conf.PasswordSalt, conf.Jwt.ExpiresAt)
	h := handler.NewHandler(s)

	if err := apiserver.Run(conf, h.InitHandler()); err != nil {
		logrus.Fatal("Can not start apiserver! err :", err)
	}
}
