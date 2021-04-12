package middle

import "github.com/gin-gonic/gin"

// CORSMiddleWare CORS跨域设置
func CORSMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		_origin := context.Request.Header.Get("Origin")

		context.Header("Access-Control-Allow-Origin", _origin) //设置允许跨域的域名
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,"+
			" X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, "+
			"Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, X-Requested-With")
		if context.Request.Method == "OPTIONS" {
			return
		}
		context.Header("Connection", "close") //设置链接关闭
		context.Next()
	}
}
