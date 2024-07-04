package main

import (
	"SummerHSE/sso/internal/handler"
	"SummerHSE/sso/internal/repository"
	"SummerHSE/sso/internal/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
)

func main() {

	var host = "psql"
	var port = "5432"
	var user = "psql"
	var password = "aboba"
	var dbname = "psql"
	var sslmode = "disable"

	var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sqlx.Open("postgres", dbInfo)
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
