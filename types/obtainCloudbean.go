package types

// ObtainCloudebeanData 领取云豆返回数据
type ObtainCloudebeanData struct {
	RawJson string
	Code    int    `json:"code"`
	Message string `json:"message"`
}
