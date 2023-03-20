package handlers

import (
	"net/http"
	"strconv"
	dto "waysbeans_be/dto/result"
	usersdto "waysbeans_be/dto/users"
	"waysbeans_be/models"
	"waysbeans_be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(c echo.Context) error {

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := dto.SuccessResult{Status: "Success", Data: users}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := dto.SuccessResult{Status: "Success", Data: convertResponse(user)}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) CreateUser(c echo.Context) error {

	request := new(usersdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest}
		return c.JSON(http.StatusBadRequest, response)
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Status:   "customer",
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateUser(c echo.Context) error {

	request := new(usersdto.UpdateUserRequest)
	if err := c.Bind(request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	data, err := h.UserRepository.UpdateUSer(user)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponse(data)})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
