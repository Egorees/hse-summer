package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"server/configs"
	"server/internal/repository"
	"server/internal/server"
	"server/internal/service"
	"server/proto"
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
	Server := server.BuildServer(Service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountsServer(s, Server)
	slog.Info("Server stating listening at 0.0.0.0:8081")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
