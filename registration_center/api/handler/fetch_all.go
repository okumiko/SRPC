package handler

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"registration_center/global"
)

func FetchAllHandler(c *gin.Context) {
	log.Println("request api/fetchall...")

	data := global.Discovery.Registry.FetchAll()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data":    data,
	})
}
