package apiserver

import (
	"context"
	"net/http"
	"time"
	"todoApp/internal/config"
)

type ApiServer struct {
	server *http.Server
}

func (s ApiServer) Run(conf *config.Config, handler http.Handler) error {
	s.server = &http.Server{
		Handler:        handler,
		Addr:           conf.Server.Host + ":" + conf.Server.Port,
		MaxHeaderBytes: 1 << conf.Server.MaxHeaderBytes,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s ApiServer) shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func Run(conf *config.Config, handler http.Handler) error {
	s := ApiServer{}

	return s.Run(conf, handler)
}
