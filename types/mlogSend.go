package types

// SendMlogData 发送 Mlog 返回数据
type SendMlogData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Feed struct {
			MlogBaseData struct {
				Id              string      `json:"id"`
				Type            int         `json:"type"`
				Text            string      `json:"text"`
				Desc            string      `json:"desc"`
				InterveneText   interface{} `json:"interveneText"`
				PubTime         int         `json:"pubTime"`
				CoverUrl        string      `json:"coverUrl"`
				GreatCover      bool        `json:"greatCover"`
				CoverPicKey     string      `json:"coverPicKey"`
				CoverHeight     int         `json:"coverHeight"`
				CoverWidth      int         `json:"coverWidth"`
				CoverColor      int         `json:"coverColor"`
				CoverDynamicUrl interface{} `json:"coverDynamicUrl"`
				AudioDTO        interface{} `json:"audioDTO"`
				Talk            interface{} `json:"talk"`
				ThreadId        string      `json:"threadId"`
				Duration        int         `json:"duration"`
			} `json:"mlogBaseData"`
			MlogExtVO struct {
				LikedCount         int           `json:"likedCount"`
				CommentCount       int           `json:"commentCount"`
				PlayCount          int           `json:"playCount"`
				ShareCount         int           `json:"shareCount"`
				Song               interface{}   `json:"song"`
				AlgSong            interface{}   `json:"algSong"`
				VideoStartPlayTime interface{}   `json:"videoStartPlayTime"`
				CanCollect         interface{}   `json:"canCollect"`
				ArtistName         interface{}   `json:"artistName"`
				Artists            []interface{} `json:"artists"`
				SpecialTag         interface{}   `json:"specialTag"`
				ChannelTag         string        `json:"channelTag"`
			} `json:"mlogExtVO"`
			UserProfile struct {
				UserId       int    `json:"userId"`
				Nickname     string `json:"nickname"`
				AvatarUrl    string `json:"avatarUrl"`
				Followed     bool   `json:"followed"`
				UserType     int    `json:"userType"`
				IsAnchor     bool   `json:"isAnchor"`
				AvatarDetail struct {
					UserType        int    `json:"userType"`
					IdentityLevel   int    `json:"identityLevel"`
					IdentityIconUrl string `json:"identityIconUrl"`
				} `json:"avatarDetail"`
			} `json:"userProfile"`
			Status   int    `json:"status"`
			ShareUrl string `json:"shareUrl"`
		} `json:"feed"`
		Event struct {
			ActName          interface{} `json:"actName"`
			PendantData      interface{} `json:"pendantData"`
			ForwardCount     int         `json:"forwardCount"`
			LotteryEventData interface{} `json:"lotteryEventData"`
			Json             string      `json:"json"`
			User             struct {
				DefaultAvatar     bool        `json:"defaultAvatar"`
				Province          int         `json:"province"`
				AuthStatus        int         `json:"authStatus"`
				Followed          bool        `json:"followed"`
				AvatarUrl         string      `json:"avatarUrl"`
				AccountStatus     int         `json:"accountStatus"`
				Gender            int         `json:"gender"`
				City              int         `json:"city"`
				Birthday          int         `json:"birthday"`
				UserId            int         `json:"userId"`
				UserType          int         `json:"userType"`
				Nickname          string      `json:"nickname"`
				Signature         string      `json:"signature"`
				Description       string      `json:"description"`
				DetailDescription string      `json:"detailDescription"`
				AvatarImgId       int         `json:"avatarImgId"`
				BackgroundImgId   int         `json:"backgroundImgId"`
				BackgroundUrl     string      `json:"backgroundUrl"`
				Authority         int         `json:"authority"`
				Mutual            bool        `json:"mutual"`
				ExpertTags        interface{} `json:"expertTags"`
				Experts           interface{} `json:"experts"`
				DjStatus          int         `json:"djStatus"`
				VipType           int         `json:"vipType"`
				RemarkName        interface{} `json:"remarkName"`
				UrlAnalyze        bool        `json:"urlAnalyze"`
				Followeds         int         `json:"followeds"`
				AvatarImgIdStr    string      `json:"avatarImgId_str"`
			} `json:"user"`
			Uuid               interface{}   `json:"uuid"`
			ExpireTime         int           `json:"expireTime"`
			RcmdInfo           interface{}   `json:"rcmdInfo"`
			EventTime          int           `json:"eventTime"`
			ActId              int           `json:"actId"`
			Pics               []interface{} `json:"pics"`
			TmplId             int           `json:"tmplId"`
			ShowTime           int           `json:"showTime"`
			Id                 int           `json:"id"`
			Type               int           `json:"type"`
			TopEvent           bool          `json:"topEvent"`
			InsiteForwardCount int           `json:"insiteForwardCount"`
			Info               struct {
				CommentThread struct {
					Id               string      `json:"id"`
					ResourceInfo     interface{} `json:"resourceInfo"`
					ResourceType     int         `json:"resourceType"`
					CommentCount     int         `json:"commentCount"`
					LikedCount       int         `json:"likedCount"`
					ShareCount       int         `json:"shareCount"`
					HotCount         int         `json:"hotCount"`
					LatestLikedUsers interface{} `json:"latestLikedUsers"`
					ResourceTitle    interface{} `json:"resourceTitle"`
					ResourceId       int         `json:"resourceId"`
					ResourceOwnerId  int         `json:"resourceOwnerId"`
					ExtProperties    interface{} `json:"extProperties"`
					XInfo            interface{} `json:"xInfo"`
				} `json:"commentThread"`
				LatestLikedUsers interface{} `json:"latestLikedUsers"`
				Liked            bool        `json:"liked"`
				Comments         interface{} `json:"comments"`
				ResourceType     int         `json:"resourceType"`
				ResourceId       int         `json:"resourceId"`
				LikedCount       int         `json:"likedCount"`
				CommentCount     int         `json:"commentCount"`
				ShareCount       int         `json:"shareCount"`
				ThreadId         string      `json:"threadId"`
			} `json:"info"`
			TailMark interface{} `json:"tailMark"`
		} `json:"event"`
	} `json:"data"`
}
