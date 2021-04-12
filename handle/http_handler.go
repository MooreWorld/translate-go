package handle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"translate-go/dto"
	"translate-go/handle/network_common"
	"translate-go/service"
)

// TestPost
func TestPost(ctx *gin.Context) {
	var testStruct dto.TestStruct
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		network_common.ErrorResponse(ctx, err)
		return
	}
	if err := json.Unmarshal(body, &testStruct); err != nil {
		network_common.ErrorResponse(ctx, err)
		return
	}

	result, err := service.TestService().TestPost(testStruct.ID, testStruct.Text)
	if err != nil {
		network_common.ErrorResponse(ctx, err)
		return
	}
	network_common.SuccessResponse(ctx, result)
}

// TestGet
func TestGet(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	text := ctx.DefaultQuery("text", "")
	if id == "" {
		network_common.ErrorResponse(ctx, "empty params")
		return
	}
	result, err := service.TestService().TestPost(id, text)
	if err != nil {
		network_common.ErrorResponse(ctx, err)
		return
	}
	network_common.SuccessResponse(ctx, result)
}
