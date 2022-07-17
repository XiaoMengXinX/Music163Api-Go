package types

// VipGrowthpointData VIP 成长值信息数据
type VipGrowthpointData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		UserLevel struct {
			UserId         int         `json:"userId"`
			Level          int         `json:"level"`
			GrowthPoint    int         `json:"growthPoint"`
			LevelName      string      `json:"levelName"`
			YesterdayPoint int         `json:"yesterdayPoint"`
			VipType        int         `json:"vipType"`
			ExtJson        string      `json:"extJson"`
			ExpireTime     int64       `json:"expireTime"`
			AvatarUrl      interface{} `json:"avatarUrl"`
			Normal         bool        `json:"normal"`
			MaxLevel       bool        `json:"maxLevel"`
		} `json:"userLevel"`
		LevelCard struct {
			RightId                           int         `json:"rightId"`
			Level                             int         `json:"level"`
			PrivilegeName                     string      `json:"privilegeName"`
			PrivilegeSubTitle                 string      `json:"privilegeSubTitle"`
			PrivilegeIconUrl                  string      `json:"privilegeIconUrl"`
			ResourceId                        int         `json:"resourceId"`
			LevelBackgroundCardImageUrl       string      `json:"levelBackgroundCardImageUrl"`
			LevelBackgroundCardExpireImageUrl string      `json:"levelBackgroundCardExpireImageUrl"`
			LevelName                         string      `json:"levelName"`
			LevelMarkImageUrl                 string      `json:"levelMarkImageUrl"`
			LevelMarkExpireImageUrl           string      `json:"levelMarkExpireImageUrl"`
			BackgroundImageUrl                string      `json:"backgroundImageUrl"`
			UpgradeFireworksImageUrl          string      `json:"upgradeFireworksImageUrl"`
			BlurryBackgroundImageUrl          string      `json:"blurryBackgroundImageUrl"`
			RedVipImageUrl                    string      `json:"redVipImageUrl"`
			RedVipExpireImageUrl              string      `json:"redVipExpireImageUrl"`
			RedVipWholeImageUrl               string      `json:"redVipWholeImageUrl"`
			RedVipExpireWholeImageUrl         string      `json:"redVipExpireWholeImageUrl"`
			RedVipBuckleImageUrl              string      `json:"redVipBuckleImageUrl"`
			RedVipExpireBuckleImageUrl        string      `json:"redVipExpireBuckleImageUrl"`
			VipGiftRightBarImageUrl           string      `json:"vipGiftRightBarImageUrl"`
			VipGiftExpireRightBarImageUrl     interface{} `json:"vipGiftExpireRightBarImageUrl"`
			VipLevelPageCardImgUrl            string      `json:"vipLevelPageCardImgUrl"`
			VipLevelPageExpireCardImgUrl      string      `json:"vipLevelPageExpireCardImgUrl"`
		} `json:"levelCard"`
	} `json:"data"`
}
