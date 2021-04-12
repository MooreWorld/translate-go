package network_common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

//DefaultResponse 默认网络返回数据格式
type DefaultResponse struct {
	Status  string `json:"status"`
	ErrCode int32  `json:"errcode"` //错误码，成功默认0，失败默认10001
	Msg     string `json:"msg"`
}

//DefaultDataResponse 默认带参数返回数据格式
type DefaultDataResponse struct {
	DefaultResponse
	Data interface{} `json:"data"`
}

// ErrorResponse 通用返回错误
func ErrorResponse(ctx *gin.Context, err interface{}, depends ...string) {
	errStr := fmt.Sprintf("%v", err)
	errReq := DefaultResponse{"fail", 10001, errStr}
	response, _ := json.Marshal(errReq)
	ctx.Status(200)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Report-Msg", errStr)
	_, _ = ctx.Writer.Write(response)
}

// SuccessResponse 通用返回成功
func SuccessResponse(ctx *gin.Context, data interface{}, depends ...string) {
	successReq := DefaultResponse{"success", 200, ""}
	var response []byte
	if data != nil {
		jsonData := DefaultDataResponse{successReq, data}
		response, _ = json.Marshal(jsonData)
	} else {
		response, _ = json.Marshal(successReq)
	}
	ctx.Status(200)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	_, _ = ctx.Writer.Write(response)
}
