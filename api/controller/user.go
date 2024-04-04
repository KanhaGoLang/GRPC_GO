package controller

import (
	"encoding/json"
	"net/http"

	"github.com/KanhaGoLang/grpc_go/api/service"
)

type UserController struct {
	userService service.UserService
}

// NewUserController initializes a new instance of UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// users := make(map[string]string)
	// users["sanjeev"] = "hello from user controller"

	users, err := uc.userService.GetUsers()
	if err != nil {
		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
