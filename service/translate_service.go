package service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"translate-go/constants"
	"translate-go/dto"
	"translate-go/tool"
)

type translateService struct {
}

var _translateService *translateService

// TranslateService 翻译服务
func TranslateService() *translateService {
	_translateService = &translateService{}
	return _translateService
}

// YouDaoAPI 有道翻译
func (service *translateService) YouDaoAPI(req *dto.TranslateReq) (rsp *dto.TranslateRsp, err error) {
	tool.LogAPI().Info("Get YouDaoAPI...")
	if err := tool.CheckParams(req.Msg); err != nil {
		return rsp, err
	}
	lts := getLTS()
	slat := getSlat(lts)
	initial := constants.WebOrigin + req.Msg + slat + constants.WebExtraFlag
	sign := tool.GetMD5(initial)
	fmt.Printf("\nmsg: %s\nlts: %s\nslat: %s\ninitail: %s\nsign: %s\n", req.Msg, lts, slat, initial, sign)

	data := &dto.YouDaoData{
		I:           req.Msg,
		From:        "AUTO",
		To:          "AUTO",
		SmartResult: "dict",
		Client:      constants.WebOrigin,
		Salt:        slat,
		Sign:        sign,
		Lts:         lts,
		Bv:          constants.Bv,
		Doctype:     "json",
		Version:     "2.1",
		KeyFrom:     "fanyi.web",
		Action:      "FY_BY_CLICKBUTTION",
	}

	headers := map[string]interface{}{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Referer":    "http://fanyi.youdao.com/",
		"Cookie":     "OUTFOX_SEARCH_USER_ID=-610384760@10.169.0.83",
	}

	reqBody, err := tool.StructToMap(data, "json")
	if err != nil {
		return nil, err
	}
	response, err := tool.HTTPPost(constants.YouDaoURL, reqBody, nil, headers)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	print("bodyStr: " + bodyStr)
	tool.LogAPI().Info("End YouDaoAPI...")
	return rsp, nil
}

func getLTS() string {
	return tool.GetMillTimeStampStr()
}

func getSlat(time string) string {
	n := strconv.Itoa(rand.Intn(10))
	return time + n
}
