package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (u UserController) AddUser(c *gin.Context) {}
