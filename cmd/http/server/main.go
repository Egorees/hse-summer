package main

import (
	"accounts/configs"
	"accounts/internal/handler"
	"accounts/internal/repository"
	"accounts/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log/slog"
)

func main() {

	DBCfg := configs.ParseDBConfig("configs/DbConfig.yaml")
	DBInfo := DBCfg.GetDBInfo()

	db, err := sqlx.Open("postgres", DBInfo)
	if err != nil {
		panic(err)
	}

	Repo := repository.BuildRepository(db)
	slog.Info("Repo built")
	Service := service.BuildService(Repo)
	slog.Info("Service built")
	Handler := handler.BuildHandler(Service)
	slog.Info("Handler built")

	Server := fasthttp.Server{
		Handler: Handler,
	}

	slog.Info("Server started and listening at 0.0.0.0:8080")
	err = Server.ListenAndServe(":8080")
	if err != nil {
		panic(err)
	}
}
