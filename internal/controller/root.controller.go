package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
)

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (r *RootController) GetRoot(c *gin.Context) {
	whoami := c.GetHeader("x-who-am-i")
	c.JSON(http.StatusOK, gin.H{
		"msg":    "Selamat Datang",
		"whoami": whoami,
	})
}

func (r *RootController) PostRoot(c *gin.Context) {
	// reading body
	var body dto.PostBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Internal Server Error",
			// "err": err.Error(),
		})
		return
	}
	// validasi
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": body,
	})
}
