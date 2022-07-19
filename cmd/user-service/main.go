package main

import (
	"database/sql"
	"fmt"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/httputil"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/krulsaidme0w/golang_pet_project_3/internal/user-service/delivery/http"
	"github.com/krulsaidme0w/golang_pet_project_3/internal/user-service/repository"
	"github.com/krulsaidme0w/golang_pet_project_3/internal/user-service/usecase"
	proto "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/grpc/proto"
	userserver "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/grpc/server"
)

func registerHttpEndpoints(router *fiber.App, handler *http.UserHandler) {
	user := router.Group("/user")

	user.Post("/", handler.SaveUser)
	user.Get("/:id", handler.GetUser)
	user.Post("/update", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}

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

	adapter := httputil.NewFiberFrameworkAdapter()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase, adapter)

	router := fiber.New()
	router.Use(cors.New())

	registerHttpEndpoints(router, userHandler)

	go func() {
		if err := router.Listen(":8000"); err != nil {
			log.Fatal(err)
		}
	}()

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	proto.RegisterUserUseCaseServer(server, userserver.NewUserService(userRepo))

	if err = server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
