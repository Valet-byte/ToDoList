package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"todoApp/internal/config"
)

func NewPostgresDB(conf *config.Config) *sqlx.DB {

	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		conf.Server.Database.Username,
		conf.Server.Database.Password,
		conf.Server.Database.Host,
		conf.Server.Database.Port,
		conf.Server.Database.DbName,
		conf.Server.Database.SslMode,
	))

	if err != nil {
		logrus.Fatal("Failed create Postgres DB connection! err :", err)
	}

	err = db.Ping()

	if err != nil {
		logrus.Fatal("Failed connect Postgres DB! err :", err)
	}

	return db
}
