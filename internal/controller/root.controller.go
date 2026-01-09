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

// Get Root
// @Summary      Trying root controller
// @Description  Trying root controller with header
// @Tags         root
// @Produce      json
// @Param        x-who-am-i	header	string	false	"custom header"
// @Success      200  {object}  dto.RootResponse
// @Router       / [get]
func (r *RootController) GetRoot(c *gin.Context) {
	whoami := c.GetHeader("x-who-am-i")
	// Set CORS Header
	// c.Header("Access-Control-Allow-Origin", "http://localhost:5173")

	log.Println("CONTROLLER/HANDLER")

	c.JSON(http.StatusOK, dto.RootResponse{
		Msg:    "Selamat Datang",
		WhoAmI: whoami,
	})
}

// Post Root
// @Summary      Trying root controller
// @Description  Trying root controller with body
// @Tags         root
// @Produce      json
// @Accept		 json
// @Param		 body	body	dto.PostBody	true	"user body"
// @Success      200  {object}  dto.RootResponse
// @Failure      500  {object}  dto.Response
// @Router       / [post]
func (r *RootController) PostRoot(c *gin.Context) {
	// Set CORS Header
	// c.Header("Access-Control-Allow-Origin", "http://localhost:5500")
	// reading body
	var body dto.PostBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Data:    []any{},
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
