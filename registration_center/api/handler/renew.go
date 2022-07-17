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

func RenewHandler(c *gin.Context) {
	log.Println("request api/renew...")
	var req model.RequestRenew
	if e := c.ShouldBindJSON(&req); e != nil {
		log.Println("error:", e)
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}

	//registry global  discovery
	instance, err := global.Discovery.Registry.Renew(req.Env, req.AppId, req.Hostname)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}

	//replication to other server
	if !req.Replication {
		global.Discovery.Nodes.Load().(*model.Nodes).Replicate(configs.Renew, instance)
	}

	//过期
	if req.DirtyTimestamp > instance.DirtyTimestamp {
		err = errcode.NotFound
	} else if req.DirtyTimestamp < instance.DirtyTimestamp { //冲突
		err = errcode.Conflict
	}
	c.JSON(http.StatusOK, gin.H{
		"code": configs.StatusOK,
	})
}
