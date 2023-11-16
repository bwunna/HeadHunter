package delivery

import (
	"HeadHunter/internal/controller"
	v1 "HeadHunter/internal/delivery/v1"
	"HeadHunter/internal/providers/db"
	"HeadHunter/internal/services/userService"
	"HeadHunter/pkg/employmentService"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

// configuration for server

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "9340fk3__132AA@"
	dbName     = "company"
	driverName = "postgres"
)

func RunGRPCServer() error {
	time.Sleep(time.Second * 0)
	dataBase, err := db.NewDB(host, user, password, dbName, driverName, port)
	if err != nil {
		return err
	}

	client, err := userService.Init("localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	newController := controller.NewController(dataBase)

	server := v1.NewGrpcServer(newController, client)

	fmt.Println("server is working")

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	employmentService.RegisterEmploymentCenterServer(grpcSrv, server)

	err = grpcSrv.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
