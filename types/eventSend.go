package types

// SendEventData 发送动态 API 的返回数据
type SendEventData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	UserId  int    `json:"userId"`
	Id      int64  `json:"id"`
	Event   struct {
		DiscussId        string      `json:"discussId"`
		ActName          interface{} `json:"actName"`
		PendantData      interface{} `json:"pendantData"`
		ForwardCount     int         `json:"forwardCount"`
		LotteryEventData interface{} `json:"lotteryEventData"`
		Json             string      `json:"json"`
		TitleAlias       interface{} `json:"titleAlias"`
		User             struct {
			DefaultAvatar      bool        `json:"defaultAvatar"`
			Province           int         `json:"province"`
			AuthStatus         int         `json:"authStatus"`
			Followed           bool        `json:"followed"`
			AvatarUrl          string      `json:"avatarUrl"`
			AccountStatus      int         `json:"accountStatus"`
			Gender             int         `json:"gender"`
			City               int         `json:"city"`
			Birthday           int64       `json:"birthday"`
			UserId             int         `json:"userId"`
			UserType           int         `json:"userType"`
			Nickname           string      `json:"nickname"`
			Signature          string      `json:"signature"`
			Description        string      `json:"description"`
			DetailDescription  string      `json:"detailDescription"`
			AvatarImgId        int64       `json:"avatarImgId"`
			BackgroundImgId    int64       `json:"backgroundImgId"`
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
			CommonIdentity      interface{} `json:"commonIdentity"`
			RelationTag         interface{} `json:"relationTag"`
			AuthenticationTypes int         `json:"authenticationTypes"`
		} `json:"user"`
		Uuid       string      `json:"uuid"`
		ExpireTime int         `json:"expireTime"`
		RcmdInfo   interface{} `json:"rcmdInfo"`
		EventTime  int         `json:"eventTime"`
		ActId      int         `json:"actId"`
		ThreadId   string      `json:"threadId"`
		ExtType    string      `json:"extType"`
		ExtSource  interface{} `json:"extSource"`
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
			ResourceId       int64         `json:"resourceId"`
			ThreadId         string        `json:"threadId"`
			ShareCount       int           `json:"shareCount"`
			CommentCount     int           `json:"commentCount"`
			LikedCount       int           `json:"likedCount"`
		} `json:"info"`
		TailMark struct {
			MarkTitle      string      `json:"markTitle"`
			MarkType       string      `json:"markType"`
			MarkResourceId string      `json:"markResourceId"`
			MarkOrpheusUrl string      `json:"markOrpheusUrl"`
			ExtInfo        interface{} `json:"extInfo"`
			Circle         struct {
				ImageUrl  string `json:"imageUrl"`
				PostCount string `json:"postCount"`
				Member    string `json:"member"`
			} `json:"circle"`
		} `json:"tailMark"`
		TypeDesc            interface{} `json:"typeDesc"`
		AlterLinkUrl        interface{} `json:"alterLinkUrl"`
		AlterLinkWebviewUrl interface{} `json:"alterLinkWebviewUrl"`
		ExtJsonInfo         struct {
			ActId          int           `json:"actId"`
			ActIds         []interface{} `json:"actIds"`
			Uuid           string        `json:"uuid"`
			ExtType        string        `json:"extType"`
			ExtSource      interface{}   `json:"extSource"`
			ExtId          string        `json:"extId"`
			CircleId       interface{}   `json:"circleId"`
			CirclePubType  interface{}   `json:"circlePubType"`
			TailMark       interface{}   `json:"tailMark"`
			TypeDesc       interface{}   `json:"typeDesc"`
			PrivacySetting int           `json:"privacySetting"`
			QuestionId     interface{}   `json:"questionId"`
			ExtParams      struct {
			} `json:"extParams"`
			VoiceInfo      interface{} `json:"voiceInfo"`
			PointTopicInfo struct {
				Id          interface{} `json:"id"`
				Type        interface{} `json:"type"`
				SubType     interface{} `json:"subType"`
				Name        interface{} `json:"name"`
				Icon        interface{} `json:"icon"`
				Desc        interface{} `json:"desc"`
				Target      interface{} `json:"target"`
				ThroughInfo interface{} `json:"throughInfo"`
				Ext         interface{} `json:"ext"`
				Parent      interface{} `json:"parent"`
			} `json:"pointTopicInfo"`
			ActivityInfos interface{} `json:"activityInfos"`
			AnonymityInfo struct {
				Anonymous int         `json:"anonymous"`
				Name      interface{} `json:"name"`
				AvatarUrl interface{} `json:"avatarUrl"`
				Me        interface{} `json:"me"`
			} `json:"anonymityInfo"`
			TitleAlias interface{} `json:"titleAlias"`
		} `json:"extJsonInfo"`
		PrivacySetting      int         `json:"privacySetting"`
		ExtPageParam        interface{} `json:"extPageParam"`
		LogInfo             interface{} `json:"logInfo"`
		BottomActivityInfos []struct {
			Id          string      `json:"id"`
			Type        int         `json:"type"`
			SubType     interface{} `json:"subType"`
			Name        string      `json:"name"`
			Icon        interface{} `json:"icon"`
			Desc        interface{} `json:"desc"`
			Target      string      `json:"target"`
			ThroughInfo interface{} `json:"throughInfo"`
			Ext         interface{} `json:"ext"`
			Parent      interface{} `json:"parent"`
		} `json:"bottomActivityInfos"`
		PointTopicInfo struct {
			Id          interface{} `json:"id"`
			Type        interface{} `json:"type"`
			SubType     interface{} `json:"subType"`
			Name        interface{} `json:"name"`
			Icon        interface{} `json:"icon"`
			Desc        interface{} `json:"desc"`
			Target      interface{} `json:"target"`
			ThroughInfo interface{} `json:"throughInfo"`
			Ext         interface{} `json:"ext"`
			Parent      interface{} `json:"parent"`
		} `json:"pointTopicInfo"`
		Voice            interface{} `json:"voice"`
		TimingInfo       interface{} `json:"timingInfo"`
		EventActionToast interface{} `json:"eventActionToast"`
		RelationTopic    interface{} `json:"relationTopic"`
		AnonymityInfo    struct {
			Anonymous int         `json:"anonymous"`
			Name      interface{} `json:"name"`
			AvatarUrl interface{} `json:"avatarUrl"`
			Me        int         `json:"me"`
		} `json:"anonymityInfo"`
		Ctrp interface{} `json:"ctrp"`
	} `json:"event"`
	Sns struct {
	} `json:"sns"`
	ResUrl      string      `json:"resUrl"`
	AfterAction interface{} `json:"afterAction"`
	JustReturn  bool        `json:"justReturn"`
}
