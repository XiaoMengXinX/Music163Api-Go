package types

// SendEventData 发送动态 API 的返回数据
type SendEventData struct {
	RawJson string
	Code    int    `json:"code"`
	Message string `json:"message"`
	UserId  int    `json:"userId"`
	Id      int    `json:"id"`
	Event   struct {
		ActName          interface{} `json:"actName"`
		PendantData      interface{} `json:"pendantData"`
		ForwardCount     int         `json:"forwardCount"`
		LotteryEventData interface{} `json:"lotteryEventData"`
		Json             string      `json:"json"`
		User             struct {
			DefaultAvatar      bool        `json:"defaultAvatar"`
			Province           int         `json:"province"`
			AuthStatus         int         `json:"authStatus"`
			Followed           bool        `json:"followed"`
			AvatarUrl          string      `json:"avatarUrl"`
			AccountStatus      int         `json:"accountStatus"`
			Gender             int         `json:"gender"`
			City               int         `json:"city"`
			Birthday           int         `json:"birthday"`
			UserId             int         `json:"userId"`
			UserType           int         `json:"userType"`
			Nickname           string      `json:"nickname"`
			Signature          string      `json:"signature"`
			Description        string      `json:"description"`
			DetailDescription  string      `json:"detailDescription"`
			AvatarImgId        int         `json:"avatarImgId"`
			BackgroundImgId    int         `json:"backgroundImgId"`
			BackgroundUrl      string      `json:"backgroundUrl"`
			Authority          int         `json:"authority"`
			Mutual             bool        `json:"mutual"`
			ExpertTags         interface{} `json:"expertTags"`
			Experts            interface{} `json:"experts"`
			DjStatus           int         `json:"djStatus"`
			VipType            int         `json:"vipType"`
			RemarkName         interface{} `json:"remarkName"`
			UrlAnalyze         bool        `json:"urlAnalyze"`
			Followeds          int         `json:"followeds"`
			AvatarImgIdStr     string      `json:"avatarImgId_str"`
			AvatarImgIdStr1    string      `json:"avatarImgIdStr"`
			BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
			VipRights          interface{} `json:"vipRights"`
			AvatarDetail       struct {
				UserType        int    `json:"userType"`
				IdentityLevel   int    `json:"identityLevel"`
				IdentityIconUrl string `json:"identityIconUrl"`
			} `json:"avatarDetail"`
			CommonIdentity interface{} `json:"commonIdentity"`
		} `json:"user"`
		Uuid       string      `json:"uuid"`
		ExpireTime int         `json:"expireTime"`
		RcmdInfo   interface{} `json:"rcmdInfo"`
		EventTime  int         `json:"eventTime"`
		ActId      int         `json:"actId"`
		Pics       []struct {
			OriginUrl      string `json:"originUrl"`
			SquareUrl      string `json:"squareUrl"`
			RectangleUrl   string `json:"rectangleUrl"`
			PcSquareUrl    string `json:"pcSquareUrl"`
			PcRectangleUrl string `json:"pcRectangleUrl"`
			Format         string `json:"format"`
			Width          int    `json:"width"`
			Height         int    `json:"height"`
			PicInfo        struct {
				OriginId         int         `json:"originId"`
				SquareId         int         `json:"squareId"`
				RectangleId      int         `json:"rectangleId"`
				PcSquareId       int         `json:"pcSquareId"`
				PcRectangleId    int         `json:"pcRectangleId"`
				OriginJpgId      int         `json:"originJpgId"`
				Format           string      `json:"format"`
				Width            int         `json:"width"`
				Height           int         `json:"height"`
				OriginIdStr      string      `json:"originIdStr"`
				PcSquareIdStr    string      `json:"pcSquareIdStr"`
				PcRectangleIdStr string      `json:"pcRectangleIdStr"`
				PcSquareUrl      interface{} `json:"pcSquareUrl"`
				PcRectangleUrl   interface{} `json:"pcRectangleUrl"`
				SquareIdStr      string      `json:"squareIdStr"`
				RectangleIdStr   string      `json:"rectangleIdStr"`
			} `json:"picInfo"`
		} `json:"pics"`
		TmplId             int  `json:"tmplId"`
		ShowTime           int  `json:"showTime"`
		InsertTime         int  `json:"insertTime"`
		Id                 int  `json:"id"`
		Type               int  `json:"type"`
		TopEvent           bool `json:"topEvent"`
		InsiteForwardCount int  `json:"insiteForwardCount"`
		Info               struct {
			CommentThread struct {
				Id               string        `json:"id"`
				ResourceInfo     interface{}   `json:"resourceInfo"`
				ResourceType     int           `json:"resourceType"`
				CommentCount     int           `json:"commentCount"`
				LikedCount       int           `json:"likedCount"`
				ShareCount       int           `json:"shareCount"`
				HotCount         int           `json:"hotCount"`
				LatestLikedUsers []interface{} `json:"latestLikedUsers"`
				ResourceOwnerId  int           `json:"resourceOwnerId"`
				ResourceId       int           `json:"resourceId"`
			} `json:"commentThread"`
			LatestLikedUsers []interface{} `json:"latestLikedUsers"`
			Liked            bool          `json:"liked"`
			Comments         []interface{} `json:"comments"`
			ResourceType     int           `json:"resourceType"`
			ResourceId       int           `json:"resourceId"`
			ThreadId         string        `json:"threadId"`
			ShareCount       int           `json:"shareCount"`
			CommentCount     int           `json:"commentCount"`
			LikedCount       int           `json:"likedCount"`
		} `json:"info"`
		TailMark    interface{} `json:"tailMark"`
		ExtJsonInfo struct {
			ActId          int           `json:"actId"`
			ActIds         []interface{} `json:"actIds"`
			Uuid           string        `json:"uuid"`
			ExtType        string        `json:"extType"`
			ExtId          string        `json:"extId"`
			CircleId       interface{}   `json:"circleId"`
			CirclePubType  interface{}   `json:"circlePubType"`
			TailMark       interface{}   `json:"tailMark"`
			TypeDesc       interface{}   `json:"typeDesc"`
			PrivacySetting int           `json:"privacySetting"`
			ExtParams      struct {
			} `json:"extParams"`
		} `json:"extJsonInfo"`
		PrivacySetting int         `json:"privacySetting"`
		ExtPageParam   interface{} `json:"extPageParam"`
		LogInfo        interface{} `json:"logInfo"`
	} `json:"event"`
	Sns struct {
	} `json:"sns"`
	ResUrl string `json:"resUrl"`
}
