package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/u22-2018/gae/service"
)

func GetHello(c *gin.Context)  {
	massage := service.GetHello()
	c.JSON(http.StatusOK,massage)
}