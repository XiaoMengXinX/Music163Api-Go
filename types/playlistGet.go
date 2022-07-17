package types

// UserPlaylistData 获取用户歌单 API 的返回数据
type UserPlaylistData struct {
	RawJson  string `json:"-"`
	Version  string `json:"version"`
	More     bool   `json:"more"`
	Playlist []struct {
		Subscribers []interface{} `json:"subscribers"`
		Subscribed  bool          `json:"subscribed"`
		Creator     struct {
			DefaultAvatar     bool     `json:"defaultAvatar"`
			Province          int      `json:"province"`
			AuthStatus        int      `json:"authStatus"`
			Followed          bool     `json:"followed"`
			AvatarUrl         string   `json:"avatarUrl"`
			AccountStatus     int      `json:"accountStatus"`
			Gender            int      `json:"gender"`
			City              int      `json:"city"`
			Birthday          int      `json:"birthday"`
			UserId            int      `json:"userId"`
			UserType          int      `json:"userType"`
			Nickname          string   `json:"nickname"`
			Signature         string   `json:"signature"`
			Description       string   `json:"description"`
			DetailDescription string   `json:"detailDescription"`
			AvatarImgId       int      `json:"avatarImgId"`
			BackgroundImgId   int      `json:"backgroundImgId"`
			BackgroundUrl     string   `json:"backgroundUrl"`
			Authority         int      `json:"authority"`
			Mutual            bool     `json:"mutual"`
			ExpertTags        []string `json:"expertTags"`
			Experts           *struct {
				Field1 string `json:"2,omitempty"`
				Field2 string `json:"1,omitempty"`
			} `json:"experts"`
			DjStatus            int         `json:"djStatus"`
			VipType             int         `json:"vipType"`
			RemarkName          interface{} `json:"remarkName"`
			AuthenticationTypes int         `json:"authenticationTypes"`
			AvatarDetail        interface{} `json:"avatarDetail"`
			AvatarImgIdStr      string      `json:"avatarImgIdStr"`
			BackgroundImgIdStr  string      `json:"backgroundImgIdStr"`
			Anchor              bool        `json:"anchor"`
			AvatarImgIdStr1     string      `json:"avatarImgId_str,omitempty"`
		} `json:"creator"`
		Artists            interface{} `json:"artists"`
		Tracks             interface{} `json:"tracks"`
		UpdateFrequency    *string     `json:"updateFrequency"`
		BackgroundCoverId  int         `json:"backgroundCoverId"`
		BackgroundCoverUrl *string     `json:"backgroundCoverUrl"`
		TitleImage         int         `json:"titleImage"`
		TitleImageUrl      *string     `json:"titleImageUrl"`
		EnglishTitle       *string     `json:"englishTitle"`
		OpRecommend        bool        `json:"opRecommend"`
		RecommendInfo      *struct {
			Alg     string `json:"alg"`
			LogInfo string `json:"logInfo"`
		} `json:"recommendInfo"`
		SubscribedCount       int         `json:"subscribedCount"`
		CloudTrackCount       int         `json:"cloudTrackCount"`
		UserId                int         `json:"userId"`
		TotalDuration         int         `json:"totalDuration"`
		CoverImgId            int         `json:"coverImgId"`
		Privacy               int         `json:"privacy"`
		TrackUpdateTime       int         `json:"trackUpdateTime"`
		TrackCount            int         `json:"trackCount"`
		UpdateTime            int         `json:"updateTime"`
		CommentThreadId       string      `json:"commentThreadId"`
		CoverImgUrl           string      `json:"coverImgUrl"`
		SpecialType           int         `json:"specialType"`
		Anonimous             bool        `json:"anonimous"`
		CreateTime            int         `json:"createTime"`
		HighQuality           bool        `json:"highQuality"`
		NewImported           bool        `json:"newImported"`
		TrackNumberUpdateTime int         `json:"trackNumberUpdateTime"`
		PlayCount             int         `json:"playCount"`
		AdType                int         `json:"adType"`
		Description           *string     `json:"description"`
		Tags                  []string    `json:"tags"`
		Ordered               bool        `json:"ordered"`
		Status                int         `json:"status"`
		Name                  string      `json:"name"`
		Id                    int         `json:"id"`
		CoverImgIdStr         *string     `json:"coverImgId_str"`
		SharedUsers           interface{} `json:"sharedUsers"`
		ShareStatus           interface{} `json:"shareStatus"`
	} `json:"playlist"`
	Code int `json:"code"`
}
