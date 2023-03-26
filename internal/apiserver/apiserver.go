package apiserver

import (
	"awesomeProject/internal/config"
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type ApiServer struct {
	server *http.Server
	log    *logrus.Logger
}

func (s ApiServer) Run(conf *config.Config) error {
	s.server = &http.Server{
		Addr:           conf.Server.Host + ":" + conf.Server.Port,
		MaxHeaderBytes: 1 << conf.Server.MaxHeaderBytes,
		ReadTimeout:    conf.Server.Timeout.Read * time.Second,
		WriteTimeout:   conf.Server.Timeout.Write * time.Second,
	}

	s.log.Info("ApiServer start!")

	return s.server.ListenAndServe()
}

func (s ApiServer) shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func Run(conf *config.Config) error {
	s := ApiServer{log: logrus.New()}

	level, err := logrus.ParseLevel(conf.Log.Level)

	if err != nil {
		log.Fatal("Not correct log level! err :", err)
	}

	s.log.SetLevel(level)

	return s.Run(conf)
}
