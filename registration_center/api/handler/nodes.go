package handler

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"registration_center/configs"
	"registration_center/global"
	"registration_center/model"
	"registration_center/pkg/errcode"
)

func NodesHandler(c *gin.Context) {
	log.Println("request api/nodes...")
	var req model.RequestNodes
	if e := c.ShouldBindJSON(&req); e != nil {
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}

	fetchData, err := global.Discovery.Registry.Fetch(req.Env, configs.DiscoveryAppId, configs.NodeStatusUp, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"data":    "",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data":    fetchData,
	})
}
