package types

// VipTaskRewardData 领取所有任务成长值返回数据
type VipTaskRewardData struct {
	RawJson string
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result bool `json:"result"`
	} `json:"data"`
}
