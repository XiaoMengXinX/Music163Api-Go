package types

// UploadFileData 上传文件返回数据
type UploadFileData struct {
	RawJson        string
	RequestId      string `json:"requestId"`
	Offset         int    `json:"offset"`
	Context        string `json:"context"`
	CallbackRetMsg string `json:"callbackRetMsg"`
	ErrCode        string `json:"errCode"`
	ErrMsg         string `json:"errMsg"`
}

// UploadNodeData 上传加速节点地址
type UploadNodeData struct {
	RawJson string
	Lbs     string   `json:"lbs"`
	Upload  []string `json:"upload"`
}

// UploadEventImgData 用于发送动态的图片信息
type UploadEventImgData struct {
	RawJson    string
	PicSubtype string
	PicInfo    struct {
		OriginId         int64  `json:"originId"`
		SquareId         int64  `json:"squareId"`
		RectangleId      int64  `json:"rectangleId"`
		PcSquareId       int64  `json:"pcSquareId"`
		PcRectangleId    int64  `json:"pcRectangleId"`
		OriginJpgId      int64  `json:"originJpgId"`
		Format           string `json:"format"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		OriginIdStr      string `json:"originIdStr"`
		SquareIdStr      string `json:"squareIdStr"`
		RectangleIdStr   string `json:"rectangleIdStr"`
		PcSquareIdStr    string `json:"pcSquareIdStr"`
		PcRectangleIdStr string `json:"pcRectangleIdStr"`
		PcSquareUrl      string `json:"pcSquareUrl"`
		PcRectangleUrl   string `json:"pcRectangleUrl"`
	} `json:"picInfo"`
	Code int `json:"code"`
}
