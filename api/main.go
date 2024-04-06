package main

import (
	"log"
	"net/http"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/KanhaGoLang/grpc_go/api/controller"
	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	port := ":1414"
	router := mux.NewRouter()

	initUserRoutes(router, "localhost:50052")

	common.MyLogger.Println(color.GreenString("Server started on port %s", port))
	log.Fatal(http.ListenAndServe(port, router))
}

func initUserRoutes(router *mux.Router, grpcAddress string) {
	// Initialize gRPC client
	userGrpcConn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		common.MyLogger.Fatalf(color.RedString("failed to dial grpc service: %v", err))
	}
	// defer userGrpcConn.Close()

	userService := service.NewUserServiceClient(userGrpcConn)

	// Create user controller with userService instance
	userController := controller.NewUserController(userService)

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")
}
