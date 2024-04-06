package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KanhaGoLang/grpc_go/api/models"
	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
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
	common.MyLogger.Println(color.YellowString("UC get all users"))

	users, err := uc.userService.GetUsers()
	if err != nil {
		common.MyLogger.Println(color.RedString(err.Error()))

		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the "id" parameter to an integer
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid User id", http.StatusBadRequest)
	}
	idInt32 := int32(id)
	common.MyLogger.Println(color.YellowString("UC Get user by id %d", idInt32))

	user, err := uc.userService.GetUser((idInt32))

	if err != nil {
		common.MyLogger.Println(color.RedString(err.Error()))

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	handleUserRequest(w, r, uc.userService.CreateUser, "Create User")
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handleUserRequest(w, r, uc.userService.UpdateUser, "Update User")
}

func handleUserRequest(w http.ResponseWriter, r *http.Request, userFunc func(*proto.User) (*proto.User, error), operation string) {
	common.MyLogger.Println(color.YellowString("UC %s", operation))

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	// validate.RegisterValidation("strength", validatePasswordStrength)
	validate.RegisterValidation("strength", common.ValidatePasswordStrength)
	validate.RegisterValidation("validateRole", common.ValidateRole(models.ValidRoles))
	err = validate.Struct(user)

	if err != nil {
		common.HandleValidationErrors(w, err)
		return
	}

	newUser := mapUserToProto(user)

	createdUser, err := userFunc(newUser)
	if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func mapUserToProto(user models.User) *proto.User {
	common.MyLogger.Println(color.CyanString("user map user to proto  %v", user))

	return &proto.User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		IsActive: user.IsActive,
	}
}
