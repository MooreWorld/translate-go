package main

import (
	"translate-go/routers"
	"translate-go/tool"
)

// 服务地址: localhost + 端口号
const PORT = ":8081"

func main() {
	tool.LogAPI().Info("start main...")
	router := routers.InitRouter()

	if err := router.Run(PORT); err != nil {
		tool.LogAPI().Info("error main...: " + err.Error())
	}
	tool.LogAPI().Info("end main...")
}
