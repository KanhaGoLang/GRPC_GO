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
	// initialize grpc client
	grpcConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial grpc service: %v", err)
	}
	defer grpcConn.Close()

	// create user service client with grpc client
	userService := service.NewUserServiceClient(grpcConn)

	// create user controller with userService instance

	userController := controller.NewUserController(userService)

	// create a new router
	router := mux.NewRouter()
	port := ":1414"
	//register router handler
	router.HandleFunc("/", userController.GetUsers).Methods("GET")

	log.Println("Server started on port : ", port)
	log.Fatal(http.ListenAndServe(port, router))

}
