package controller

import (
	"encoding/json"
	"fmt"
	"log"
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
	w.Header().Set("Content-Type", "application/json")
	common.MyLogger.Println(color.YellowString("UC Create User"))

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		common.MyLogger.Println(color.RedString(err.Error()))

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// err = validator.New().Struct(user)

	validate := validator.New()
	validate.RegisterValidation("strength", validatePasswordStrength)
	validate.RegisterValidation("validateRole", validateRole)
	err = validate.Struct(user)

	if err != nil {
		var validationErrors = make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[getJSONTag(err)] = getValidationErrorMsg(err)
			// fmt.Sprintf("%s is %s", getJSONTag(err), err.Tag())
			// fmt.Sprintf("%s", err.Tag())
		}

		errorResponse := ErrorResponse{Message: "Validation failed", Details: validationErrors}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	newUser := mapUserToProto(user)

	createdUser, err := uc.userService.CreateUser(newUser)
	if err != nil {
		common.MyLogger.Println(color.RedString(err.Error()))

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)

}

func mapUserToProto(user models.User) *proto.User {
	return &proto.User{
		// Id:        12,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		IsActive: user.IsActive,
		// CreatedAt: "123",
		// UpdatedAt: "456",
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

func validateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	for _, validRole := range models.ValidRoles {
		if validRole == role {
			return true
		}
	}

	return false
}

func getValidationErrorMsg(err validator.FieldError) string {
	log.Println(err.Tag(), err.Field())
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
