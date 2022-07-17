package types

// VipInfoData VIP 信息数据
type VipInfoData struct {
	RawJson string `json:"-"`
	Message string `json:"message"`
	Data    struct {
		RedVipLevel       int   `json:"redVipLevel"`
		Now               int64 `json:"now"`
		RedVipAnnualCount int   `json:"redVipAnnualCount"`
		MusicPackage      struct {
			VipCode         int   `json:"vipCode"`
			ExpireTime      int64 `json:"expireTime"`
			IsSignDeduct    bool  `json:"isSignDeduct"`
			IsSignIapDeduct bool  `json:"isSignIapDeduct"`
			IsSign          bool  `json:"isSign"`
			IsSignIap       bool  `json:"isSignIap"`
		} `json:"musicPackage"`
		Associator struct {
			VipCode         int   `json:"vipCode"`
			ExpireTime      int64 `json:"expireTime"`
			IsSignDeduct    bool  `json:"isSignDeduct"`
			IsSignIapDeduct bool  `json:"isSignIapDeduct"`
			IsSign          bool  `json:"isSign"`
			IsSignIap       bool  `json:"isSignIap"`
		} `json:"associator"`
		OldCacheProtocol bool `json:"oldCacheProtocol"`
		UserId           int  `json:"userId"`
	} `json:"data"`
	Code int `json:"code"`
}
