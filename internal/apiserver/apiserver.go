package apiserver

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
	"todoApp/internal/config"
)

type ApiServer struct {
	server *http.Server
	log    *logrus.Logger
}

func (s ApiServer) Run(conf *config.Config, handler http.Handler) error {
	s.server = &http.Server{
		Handler:        handler,
		Addr:           conf.Server.Host + ":" + conf.Server.Port,
		MaxHeaderBytes: 1 << conf.Server.MaxHeaderBytes,
		ReadTimeout:    conf.Server.Timeout.Read * time.Second,
		WriteTimeout:   conf.Server.Timeout.Write * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s ApiServer) shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func Run(conf *config.Config, handler http.Handler) error {
	s := ApiServer{log: logrus.New()}

	level, err := logrus.ParseLevel(conf.Log.Level)

	if err != nil {
		log.Fatal("Not correct log level! err :", err)
	}

	s.log.SetLevel(level)

	return s.Run(conf, handler)
}
