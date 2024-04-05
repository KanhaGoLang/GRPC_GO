package main

import (
	"log"
	"net/http"

	"github.com/KanhaGoLang/grpc_go/api/controller"
	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	port := ":1414"
	router := mux.NewRouter()

	// initialize grpc client
	userGrpcConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial grpc service: %v", err)
	}
	defer userGrpcConn.Close()

	initUserRoutes(router, userGrpcConn)

	log.Println("Server started on port : ", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initUserRoutes(router *mux.Router, userGrpcConn *grpc.ClientConn) {

	userService := service.NewUserServiceClient(userGrpcConn)

	// create user controller with userService instance

	userController := controller.NewUserController(userService)
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")
}
