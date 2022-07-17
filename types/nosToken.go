package types

// NosTokenData NosToken 数据
type NosTokenData struct {
	RawJson string `json:"-"`
	Type    int
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Bucket    string `json:"bucket"`
		DocId     string `json:"docId"`
		ObjectKey string `json:"objectKey"`
		Token     string `json:"token"`
	} `json:"result"`
	Data struct {
		Bucket     string `json:"bucket"`
		Token      string `json:"token"`
		OuterUrl   string `json:"outerUrl"`
		DocId      string `json:"docId"`
		ObjectKey  string `json:"objectKey"`
		ResourceId int    `json:"resourceId"`
	} `json:"data"`
}
