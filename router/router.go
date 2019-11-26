package router

import (
	"github.com/adolfheir/go-wechat/consts"
	wechat "github.com/adolfheir/go-wechat/controller/wx"
	_ "github.com/adolfheir/go-wechat/docs"
	"github.com/adolfheir/go-wechat/middleware" // 微信SDK包
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//全局路由实例

func InitRouter() *gin.Engine {
	routerInstance := gin.New()
	//doc
	routerInstance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//中间件
	routerInstance.Use(middleware.Logger())

	// wx相关接口
	wxGroup := routerInstance.Group("/wx")

	wxGroup.GET("/check", wechat.WxCheck)
	wxGroup.POST("/check", wechat.WxCallbackHandler)

	//api接口
	apiGroup := routerInstance.Group("/api/v1")

	apiGroup.GET("/", func(c *gin.Context) {
		c.JSON(consts.SUCCESS, gin.H{
			"success": true,
			"code":    200,
			"message": "This works",
			"data":    nil,
		})
	})
	return routerInstance
}
