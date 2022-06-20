package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler punya ropository
type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// Tangkat input dari user
	// Map input dari user ke struct RegisterUserIput
	// struct diatas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err) // diambil dari helper
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()
	formater := user.FormatUser(newUser, "tokentokentoken")

	// ambil dari helper.go
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//user memasukkan input (email & password)
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct passing service
	// diservice mencari dg bantuan repositiry user dengan email x
	//mencocokkan password
}
