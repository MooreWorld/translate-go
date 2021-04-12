package service

import (
	"fmt"
	"translate-go/dto"
	"translate-go/tool"
)

type testService struct {
}

var _testService *testService

// TestService 测试服务
func TestService() *testService {
	_testService = &testService{}
	return _testService
}

// TestPost 测试POST
func (api *testService) TestPost(id, text string) (result *dto.TestStruct, err error) {
	tool.LogAPI().Info("start TestPost...")
	fmt.Printf("id: %s, text: %s\n", id, text)
	tool.LogAPI().Info("end TestPost...")
	result = &dto.TestStruct{
		ID:   id,
		Text: text,
	}
	return result, nil
}

// TestGet 测试Get
func (api *testService) TestGet(id, text string) (result *dto.TestStruct, err error) {
	tool.LogAPI().Info("start TestPost...")
	fmt.Printf("id: %s, text: %s\n", id, text)
	tool.LogAPI().Info("end TestPost...")
	result = &dto.TestStruct{
		ID:   id,
		Text: text,
	}
	return result, nil
}
