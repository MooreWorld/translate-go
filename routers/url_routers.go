package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"translate-go/constants"
	"translate-go/handle"
	"translate-go/middle"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(middle.CORSMiddleWare())
	router.Handle("OPTIONS", "/*any", func(c *gin.Context) {
		return
	})

	// 测试服务API
	testGroup := router.Group(fmt.Sprintf("/%s", constants.ApiVersion))
	{
		testGroup.POST("/add/text", handle.TestPost)
		testGroup.GET("/get/text", handle.TestGet)
	}

	// 翻译服务API
	translateGroup := router.Group(fmt.Sprintf("/%s", constants.ApiVersion))
	{
		translateGroup.POST("/fanyi/youdao", handle.YouDao)
	}

	// 文件服务API
	fileGroup := router.Group(fmt.Sprintf("/%s", constants.ApiVersion))
	{
		fileGroup.GET("/picture/bing", handle.BingPic)
	}

	return router
}
