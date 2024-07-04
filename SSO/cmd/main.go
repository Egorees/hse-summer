package main

import (
	"SSO/configs"
	"SSO/internal/handler"
	"SSO/internal/repository"
	"SSO/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

func main() {

	DBCfg := configs.ParseDBConfig("configs/DbConfig.yaml")
	DBInfo := DBCfg.GetDBInfo()

	db, err := sqlx.Open("postgres", DBInfo)
	if err != nil {
		panic(err)
	}

	Repo := repository.BuildRepository(db)
	Service := service.BuildService(Repo)
	Handler := handler.BuildHandler(Service)

	err = fasthttp.ListenAndServe(":8080", Handler)
	if err != nil {
		panic(err)
	}
}
