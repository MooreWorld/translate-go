package service

import (
	"encoding/json"
	"io/ioutil"
	"translate-go/constants"
	"translate-go/dto"
	"translate-go/tool"
)

type pictureService struct {
}

var _pictureService *pictureService

// PictureService 图片服务
func PictureService() *pictureService {
	_pictureService = &pictureService{}
	return _pictureService
}

// GetBingPic 获取必应图片资源
func (service *pictureService) GetBingPic() (picData *dto.PicData, err error) {
	tool.LogAPI().Info("start GetBingPic...")
	response, err := tool.HTTPGet(constants.BingPicURL, nil, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	bingData := &dto.BingDataRsp{}
	err = json.Unmarshal([]byte(bodyStr), bingData)
	if err != nil {
		return nil, err
	}
	picData = &dto.PicData{
		ImageURL:    constants.BingBaseURL + bingData.Image[0].Url,
		HashSha:     bingData.Image[0].Hsh,
		CopyRight:   bingData.Image[0].CopyRight,
		CreateDate:  bingData.Image[0].FullStartDate,
		Tips:        bingData.ToolTips.Walls,
		ImageOrigin: bingData.Image[0].Quiz,
	}
	tool.LogAPI().Info("end GetBingPic...")
	return picData, nil
}
