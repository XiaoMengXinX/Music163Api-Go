package types

// SendMsgData 发送私信 API 的返回数据
type SendMsgData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	NewMsgs []struct {
		FromUser struct {
			Description        string      `json:"description"`
			BackgroundImgId    int         `json:"backgroundImgId"`
			Birthday           int         `json:"birthday"`
			AccountStatus      int         `json:"accountStatus"`
			City               int         `json:"city"`
			DetailDescription  string      `json:"detailDescription"`
			DefaultAvatar      bool        `json:"defaultAvatar"`
			DjStatus           int         `json:"djStatus"`
			Followed           bool        `json:"followed"`
			BackgroundUrl      string      `json:"backgroundUrl"`
			Gender             int         `json:"gender"`
			AvatarImgId        int         `json:"avatarImgId"`
			AvatarDetail       interface{} `json:"avatarDetail"`
			UserType           int         `json:"userType"`
			UserId             int         `json:"userId"`
			Nickname           string      `json:"nickname"`
			Province           int         `json:"province"`
			Mutual             bool        `json:"mutual"`
			AvatarUrl          string      `json:"avatarUrl"`
			AuthStatus         int         `json:"authStatus"`
			ExpertTags         interface{} `json:"expertTags"`
			RemarkName         interface{} `json:"remarkName"`
			VipType            int         `json:"vipType"`
			Experts            interface{} `json:"experts"`
			AvatarImgIdStr     string      `json:"avatarImgIdStr"`
			BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
			Signature          string      `json:"signature"`
			Authority          int         `json:"authority"`
		} `json:"fromUser"`
		ToUser struct {
			Description        string      `json:"description"`
			BackgroundImgId    int         `json:"backgroundImgId"`
			Birthday           int         `json:"birthday"`
			AccountStatus      int         `json:"accountStatus"`
			City               int         `json:"city"`
			DetailDescription  string      `json:"detailDescription"`
			DefaultAvatar      bool        `json:"defaultAvatar"`
			DjStatus           int         `json:"djStatus"`
			Followed           bool        `json:"followed"`
			BackgroundUrl      string      `json:"backgroundUrl"`
			Gender             int         `json:"gender"`
			AvatarImgId        int         `json:"avatarImgId"`
			AvatarDetail       interface{} `json:"avatarDetail"`
			UserType           int         `json:"userType"`
			UserId             int         `json:"userId"`
			Nickname           string      `json:"nickname"`
			Province           int         `json:"province"`
			Mutual             bool        `json:"mutual"`
			AvatarUrl          string      `json:"avatarUrl"`
			AuthStatus         int         `json:"authStatus"`
			ExpertTags         interface{} `json:"expertTags"`
			RemarkName         interface{} `json:"remarkName"`
			VipType            int         `json:"vipType"`
			Experts            interface{} `json:"experts"`
			AvatarImgIdStr     string      `json:"avatarImgIdStr"`
			BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
			Signature          string      `json:"signature"`
			Authority          int         `json:"authority"`
		} `json:"toUser"`
		RealFromUser interface{} `json:"realFromUser"`
		Msg          string      `json:"msg"`
		Time         int         `json:"time"`
		BatchId      int         `json:"batchId"`
		Id           int         `json:"id"`
	} `json:"newMsgs"`
	Id            int   `json:"id"`
	Sendblacklist []int `json:"sendblacklist"`
	Blacklist     []int `json:"blacklist"`
}
