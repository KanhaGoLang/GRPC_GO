package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/KanhaGoLang/grpc_go/api/models"
	"github.com/KanhaGoLang/grpc_go/api/service"
	"github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/KanhaGoLang/go_common/common"
)

type ErrorResponse struct {
	Message string            `json:"message"`
	Details map[string]string `json:"validationErrors"`
}

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
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	validate.RegisterValidation("strength", validatePasswordStrength)
	validate.RegisterValidation("validateRole", validateRole(models.ValidRoles))
	err = validate.Struct(user)

	if err != nil {
		handleValidationErrors(w, err)
		return
	}

	newUser := mapUserToProto(user)

	createdUser, err := userFunc(newUser)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func handleError(w http.ResponseWriter, err error, status int) {
	common.MyLogger.Println(color.RedString(err.Error()))
	http.Error(w, err.Error(), status)
}

func handleValidationErrors(w http.ResponseWriter, err error) {
	var validationErrors = make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors[getJSONTag(err)] = getValidationErrorMsg(err)
	}

	errorResponse := ErrorResponse{Message: "Validation failed", Details: validationErrors}

	common.MyLogger.Println(color.RedString("ErrorResponse %v", errorResponse))

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
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

func getJSONTag(err validator.FieldError) string {
	field, _ := reflect.TypeOf(models.User{}).FieldByName(err.StructField())
	return field.Tag.Get("json")
}

func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return len(password) >= 8
}

func validateRole(roles []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		for _, validRole := range roles {
			if validRole == role {
				return true
			}
		}

		return false
	}

}

func getValidationErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s is below minimum length", err.Field())
	case "max":
		return fmt.Sprintf("%s is above maximum length", err.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email address", err.Field())
	case "validateRole":
		return fmt.Sprintf("%s is not one of the allowed values %s", err.Field(), strings.Join(models.ValidRoles, ", "))
	case "strength":
		return fmt.Sprintf("%s is too weak", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
