package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"github.com/KanhaGoLang/go_common/common"
)

type UserController struct {
	userService service.UserService
}

// NewUserController initializes a new instance of UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("API USER get all users")

	users, err := uc.userService.GetUsers()
	log.Println("error >>> ", err)
	if err != nil {
		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	common.MyLogger.Println(color.YellowString("UC Get user by id"))
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the "id" parameter to an integer
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid User id", http.StatusBadRequest)
	}

	user, err := uc.userService.GetUser(int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
