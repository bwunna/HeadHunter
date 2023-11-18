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
)

// service starting

const (
	host           = "localhost"
	port           = 5432
	user           = "postgres"
	password       = "9340fk3__132AA@"
	dbName         = "company"
	driverName     = "postgres"
	servicePort    = ":8082"
	userServiceURL = "localhost:8080"
)

func RunGRPCServer() error {
	dataBase, err := db.New(host, user, password, dbName, driverName, port)
	if err != nil {
		return err
	}

	client, err := userService.Init(userServiceURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	newController := controller.New(dataBase)

	server := v1.NewGrpcServer(newController, client)

	fmt.Println("server is working")

	lis, err := net.Listen("tcp", servicePort)
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
