package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/krulsaidme0w/golang_pet_project_3/internal/user-service/repository"
	proto "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/grpc/proto"
	userserver "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/grpc/server"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"user",
		"password",
		"library"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	proto.RegisterUserUseCaseServer(server, userserver.NewUserService(userRepo))

	err = server.Serve(listen)
}
