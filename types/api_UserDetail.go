package types

// UserDetailData 用户详细信息数据
type UserDetailData struct {
	Level       int `json:"level"`
	ListenSongs int `json:"listenSongs"`
	UserPoint   struct {
		UserId       int   `json:"userId"`
		Balance      int   `json:"balance"`
		UpdateTime   int64 `json:"updateTime"`
		Version      int   `json:"version"`
		Status       int   `json:"status"`
		BlockBalance int   `json:"blockBalance"`
	} `json:"userPoint"`
	MobileSign bool `json:"mobileSign"`
	PcSign     bool `json:"pcSign"`
	Profile    struct {
		AvatarDetail       interface{} `json:"avatarDetail"`
		AvatarImgIdStr     string      `json:"avatarImgIdStr"`
		BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
		UserId             int         `json:"userId"`
		AvatarImgId        int64       `json:"avatarImgId"`
		Birthday           int64       `json:"birthday"`
		Gender             int         `json:"gender"`
		Nickname           string      `json:"nickname"`
		AccountStatus      int         `json:"accountStatus"`
		Province           int         `json:"province"`
		VipType            int         `json:"vipType"`
		Followed           bool        `json:"followed"`
		UserType           int         `json:"userType"`
		CreateTime         int64       `json:"createTime"`
		AvatarUrl          string      `json:"avatarUrl"`
		AuthStatus         int         `json:"authStatus"`
		DetailDescription  string      `json:"detailDescription"`
		Experts            struct {
		} `json:"experts"`
		ExpertTags      interface{} `json:"expertTags"`
		City            int         `json:"city"`
		DefaultAvatar   bool        `json:"defaultAvatar"`
		BackgroundImgId int64       `json:"backgroundImgId"`
		BackgroundUrl   string      `json:"backgroundUrl"`
		Mutual          bool        `json:"mutual"`
		RemarkName      interface{} `json:"remarkName"`
		Description     string      `json:"description"`
		DjStatus        int         `json:"djStatus"`
		Signature       string      `json:"signature"`
		Authority       int         `json:"authority"`
		AllAuthTypes    []struct {
			Type int         `json:"type"`
			Desc string      `json:"desc"`
			Tags interface{} `json:"tags"`
		} `json:"allAuthTypes"`
		Followeds                 int  `json:"followeds"`
		Follows                   int  `json:"follows"`
		Blacklist                 bool `json:"blacklist"`
		EventCount                int  `json:"eventCount"`
		AllSubscribedCount        int  `json:"allSubscribedCount"`
		PlaylistBeSubscribedCount int  `json:"playlistBeSubscribedCount"`
		MainAuthType              struct {
			Type int         `json:"type"`
			Desc string      `json:"desc"`
			Tags interface{} `json:"tags"`
		} `json:"mainAuthType"`
		FollowTime     string        `json:"followTime"`
		FollowMe       bool          `json:"followMe"`
		ArtistIdentity []interface{} `json:"artistIdentity"`
		CCount         int           `json:"cCount"`
		SDJPCount      int           `json:"sDJPCount"`
		PlaylistCount  int           `json:"playlistCount"`
		SCount         int           `json:"sCount"`
		NewFollows     int           `json:"newFollows"`
	} `json:"profile"`
	PeopleCanSeeMyPlayRecord bool `json:"peopleCanSeeMyPlayRecord"`
	Bindings                 []struct {
		UserId       int         `json:"userId"`
		ExpiresIn    int         `json:"expiresIn"`
		RefreshTime  int         `json:"refreshTime"`
		BindingTime  int64       `json:"bindingTime"`
		TokenJsonStr interface{} `json:"tokenJsonStr"`
		Expired      bool        `json:"expired"`
		Url          string      `json:"url"`
		Id           int64       `json:"id"`
		Type         int         `json:"type"`
	} `json:"bindings"`
	AdValid    bool  `json:"adValid"`
	Code       int   `json:"code"`
	CreateTime int64 `json:"createTime"`
	CreateDays int   `json:"createDays"`
}
