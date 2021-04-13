package dto

// TestStruct 测试专用
type TestStruct struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// TranslateReq 翻译对象 入参
type TranslateReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

// TranslateRsp 翻译对象 出参
type TranslateRsp struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Before string `json:"before"`
	After  string `json:"after"`
}

// YouDaoData 有道API data
type YouDaoData struct {
	I           string `json:"i"`
	From        string `json:"from"`
	To          string `json:"to"`
	SmartResult string `json:"smartresult"`
	Client      string `json:"client"`
	Salt        string `json:"salt"`
	Sign        string `json:"sign"`
	Lts         string `json:"lts"`
	Bv          string `json:"bv"`
	Doctype     string `json:"doctype"`
	Version     string `json:"version"`
	KeyFrom     string `json:"keyfrom"`
	Action      string `json:"action"`
}

// Images bing图片对象
type Images struct {
	StartDate     string `json:"startdate"`
	FullStartDate string `json:"fullstartdate"`
	EndDate       string `json:"enddate"`
	Url           string `json:"url"`
	UrlBase       string `json:"urlbase"`
	CopyRight     string `json:"copyright"`
	CopyRightLink string `json:"copyrightlink"`
	Title         string `json:"title"`
	Quiz          string `json:"quiz"`
	Wp            bool   `json:"wp"`
	Hsh           string `json:"hsh"`
	Drk           int    `json:"drk"`
	Top           int    `json:"top"`
	Bot           int    `json:"bot"`
}

// ToolTips bing工具对象
type ToolTips struct {
	Loading  string `json:"loading"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Walle    string `json:"walle"`
	Walls    string `json:"walls"`
}

// BingDataRsp bing返回对象
type BingDataRsp struct {
	Image    []*Images `json:"images"`
	ToolTips ToolTips  `json:"tooltips"`
}

// PicData 图片数据
type PicData struct {
	ImageURL    string `json:"image_url"`
	HashSha     string `json:"hash_sha"`
	CopyRight   string `json:"copy_right"`
	CreateDate  string `json:"create_date"`
	ImageOrigin string `json:"image_origin"`
	Tips        string `json:"tips"`
}
