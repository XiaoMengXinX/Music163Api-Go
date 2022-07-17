package types

// LoginStatusData 登录状态数据'
type LoginStatusData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Account struct {
		Id                 int    `json:"id"`
		UserName           string `json:"userName"`
		Type               int    `json:"type"`
		Status             int    `json:"status"`
		WhitelistAuthority int    `json:"whitelistAuthority"`
		CreateTime         int    `json:"createTime"`
		TokenVersion       int    `json:"tokenVersion"`
		Ban                int    `json:"ban"`
		BaoyueVersion      int    `json:"baoyueVersion"`
		DonateVersion      int    `json:"donateVersion"`
		VipType            int    `json:"vipType"`
		AnonimousUser      bool   `json:"anonimousUser"`
		PaidFee            bool   `json:"paidFee"`
	} `json:"account"`
	Profile struct {
		UserId              int         `json:"userId"`
		UserType            int         `json:"userType"`
		Nickname            string      `json:"nickname"`
		AvatarImgId         int         `json:"avatarImgId"`
		AvatarUrl           string      `json:"avatarUrl"`
		BackgroundImgId     int         `json:"backgroundImgId"`
		BackgroundUrl       string      `json:"backgroundUrl"`
		Signature           string      `json:"signature"`
		CreateTime          int         `json:"createTime"`
		UserName            string      `json:"userName"`
		AccountType         int         `json:"accountType"`
		ShortUserName       string      `json:"shortUserName"`
		Birthday            int         `json:"birthday"`
		Authority           int         `json:"authority"`
		Gender              int         `json:"gender"`
		AccountStatus       int         `json:"accountStatus"`
		Province            int         `json:"province"`
		City                int         `json:"city"`
		AuthStatus          int         `json:"authStatus"`
		Description         string      `json:"description"`
		DetailDescription   string      `json:"detailDescription"`
		DefaultAvatar       bool        `json:"defaultAvatar"`
		ExpertTags          interface{} `json:"expertTags"`
		Experts             interface{} `json:"experts"`
		DjStatus            int         `json:"djStatus"`
		LocationStatus      int         `json:"locationStatus"`
		VipType             int         `json:"vipType"`
		Followed            bool        `json:"followed"`
		Mutual              bool        `json:"mutual"`
		Authenticated       bool        `json:"authenticated"`
		LastLoginTime       int         `json:"lastLoginTime"`
		LastLoginIP         string      `json:"lastLoginIP"`
		RemarkName          interface{} `json:"remarkName"`
		ViptypeVersion      int         `json:"viptypeVersion"`
		AuthenticationTypes int         `json:"authenticationTypes"`
		AvatarDetail        interface{} `json:"avatarDetail"`
		Anchor              bool        `json:"anchor"`
	} `json:"profile"`
}
