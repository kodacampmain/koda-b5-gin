package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u UserController) GetUsers(c *gin.Context) {
	data, err := u.userService.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": data,
	})
}

func (u UserController) AddUser(c *gin.Context) {
	var newUser dto.NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	data, err := u.userService.AddUser(c.Request.Context(), newUser)
	if err != nil {
		// Expected Error
		if errors.Is(err, service.ErrInvalidGender) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   err.Error(),
				Data:    []any{},
			})
			return
		}
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   "Name already in use",
				Data:    []any{},
			})
			return
		}
		// Unexpected Error
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Create User Success",
		Success: true,
		Data:    []any{data},
	})
}

func (u UserController) Register(c *gin.Context) {
	var newUser service.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	if err := u.userService.Register(newUser); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   "Failed To Register",
			Data:    []any{},
		})
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Register Success",
		Success: true,
		Data:    []any{},
	})
}

func (u UserController) Login(c *gin.Context) {
	var newUser service.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	isValid, err := u.userService.Login(&newUser)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "email/password is wrong") {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   err.Error(),
				Data:    []any{},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	// log.Println("isvalid", isValid)
	if !isValid {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   "email/password is wrong",
			Data:    []any{},
		})
		return
	}
	token, err := u.userService.GenJWTToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Login Success",
		Success: true,
		Data: []any{
			gin.H{
				"token": token,
			},
		},
	})
}
