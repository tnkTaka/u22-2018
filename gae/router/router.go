package router

import (
	"github.com/gin-gonic/gin"
	"github.com/u22-2018/gae/controller"
)

// GetRouter ルーターを設定してgin.Engineを返す
func GetRouter() *gin.Engine {
	r := gin.Default()
	//r.Static("/assets", "./assets")
	//r.LoadHTMLGlob("./templates/*")

	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	r.GET("/", controller.GetHello)

	return r
}
