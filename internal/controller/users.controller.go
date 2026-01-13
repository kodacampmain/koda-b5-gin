package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/err"
	"github.com/kodacampmain/koda-b5-gin/internal/service"
	"github.com/kodacampmain/koda-b5-gin/pkg"
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
		log.Println(err.Error())
		if strings.Contains(err.Error(), "required") {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Register have to include all field: email, password, gender",
				Success: false,
				Error:   "Bad request",
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
	data, err := u.userService.Register(c.Request.Context(), newUser)
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
				Error:   "Email already registered",
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
	c.JSON(http.StatusCreated, dto.UsersResponse{
		Response: dto.Response{
			Msg:     "Create User Success",
			Success: true,
		},
		Data: []dto.User{data},
	})
}

func (u UserController) EditProfile(c *gin.Context) {
	// Data binding
	var user dto.EditUser
	if e := c.ShouldBindWith(&user, binding.FormMultipart); e != nil {
		log.Println("binding", e.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	// akses claims
	token, _ := c.Get("token")
	accessToken, _ := token.(pkg.JWTClaims)
	// validasi ekstensi (jpg, png)
	ext := path.Ext(user.Image.Filename)
	re := regexp.MustCompile("^[.](jpg|png)$")
	if !re.Match([]byte(ext)) {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     err.ErrInvalidExt.Error(),
			Error:   "Bad Request",
			Success: false,
			Data:    []any{},
		})
		return
	}
	// validasi ukuran
	// validasi banyak (multi upload)

	// Misal image akan ditaruh di /namaFile.ext
	// timestamp_function.ext
	filename := fmt.Sprintf("%d_profile_%d%s", time.Now().UnixNano(), accessToken.Id, ext)

	// save upload file
	// 1. simpan di db (sbg blob)
	// 2. simpan di storage (yg di demo kan)
	// 3. simpan di cloud (buat reusable code untuk cloud service)
	if e := c.SaveUploadedFile(user.Image, filepath.Join("public", "profile", filename)); e != nil {
		log.Println(e.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}

	if e := u.userService.UpdateImage(c.Request.Context(), fmt.Sprintf("/profile/%s", filename), accessToken.Id); e != nil {
		log.Println(e.Error())
		if errors.Is(e, err.ErrNoRowsUpdated) {
			c.JSON(http.StatusNotFound, dto.Response{
				Msg:     e.Error(),
				Success: false,
				Error:   "User not found",
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

	// logika untuk menghapus file lama

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "OK",
		Success: true,
		Data: []any{
			gin.H{
				"img": fmt.Sprintf("/profile/%s", filename),
			},
		},
	})
}

// func (u UserController) Register(c *gin.Context) {
// 	var newUser service.User
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusInternalServerError, dto.Response{
// 			Msg:     "Internal Server Error",
// 			Success: false,
// 			Error:   "internal server error",
// 			Data:    []any{},
// 		})
// 		return
// 	}
// 	if err := u.userService.Register(newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, dto.Response{
// 			Msg:     "Bad Request",
// 			Success: false,
// 			Error:   "Failed To Register",
// 			Data:    []any{},
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, dto.Response{
// 		Msg:     "Register Success",
// 		Success: true,
// 		Data:    []any{},
// 	})
// }

func (u UserController) Login(c *gin.Context) {
	var user dto.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	userInfo, err := u.userService.Login(c.Request.Context(), user.Email, user.Password)
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
	token, err := u.userService.GenJWTToken(userInfo)
	if err != nil {
		log.Println(err.Error())
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
