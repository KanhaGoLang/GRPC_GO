package routes

import (
	"net/http"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/KanhaGoLang/grpc_go/api/controller"
	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitUserRoutes(router *mux.Router) {
	common.MyLogger.Println(color.MagentaString("Init User Routes"))
	common.MyLogger.Println(color.MagentaString("Establishing connection to User GRPC service on port %s", common.UserServiceAddress))

	// Initialize gRPC client
	userGrpcConn, err := grpc.Dial(common.UserServiceAddress, grpc.WithInsecure())
	if err != nil {
		common.MyLogger.Fatalf(color.RedString("failed to dial grpc service: %v", err))
	}
	// defer userGrpcConn.Close()

	userService := service.NewUserServiceClient(userGrpcConn)

	// Create user controller with userService instance
	userController := controller.NewUserController(userService)

	// Middleware to set Content-Type to application/json for all routes
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/user", userController.UpdateUser).Methods("PUT")
}
