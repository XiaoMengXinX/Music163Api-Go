package types

// UserDetailData 用户详细信息数据
type UserDetailData struct {
	RawJson string

	Identify struct {
		ImageUrl  string      `json:"imageUrl"`
		ImageDesc string      `json:"imageDesc"`
		ActionUrl interface{} `json:"actionUrl"`
	} `json:"identify"`
	CurrentExpert struct {
		RoleId   int         `json:"roleId"`
		RoleName string      `json:"roleName"`
		Level    interface{} `json:"level"`
	} `json:"currentExpert"`
	ExpertArray []struct {
		RoleId   int         `json:"roleId"`
		RoleName string      `json:"roleName"`
		Level    interface{} `json:"level"`
	} `json:"expertArray"`
	CurrentProduct struct {
		ProductionTypeName string `json:"productionTypeName"`
		ProductionTypeId   int    `json:"productionTypeId"`
	} `json:"currentProduct"`
	Products []struct {
		ProductionTypeName string `json:"productionTypeName"`
		ProductionTypeId   int    `json:"productionTypeId"`
	} `json:"products"`
	Level       int `json:"level"`
	ListenSongs int `json:"listenSongs"`
	UserPoint   struct {
		UserId       int `json:"userId"`
		Balance      int `json:"balance"`
		UpdateTime   int `json:"updateTime"`
		Version      int `json:"version"`
		Status       int `json:"status"`
		BlockBalance int `json:"blockBalance"`
	} `json:"userPoint"`
	MobileSign bool `json:"mobileSign"`
	PcSign     bool `json:"pcSign"`
	Profile    struct {
		AvatarDetail struct {
			UserType        interface{} `json:"userType"`
			IdentityLevel   int         `json:"identityLevel"`
			IdentityIconUrl string      `json:"identityIconUrl"`
		} `json:"avatarDetail"`
		BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
		UserId             int         `json:"userId"`
		DjStatus           int         `json:"djStatus"`
		AccountStatus      int         `json:"accountStatus"`
		Province           int         `json:"province"`
		VipType            int         `json:"vipType"`
		Followed           bool        `json:"followed"`
		CreateTime         int         `json:"createTime"`
		AvatarImgId        int         `json:"avatarImgId"`
		Birthday           int         `json:"birthday"`
		Gender             int         `json:"gender"`
		Nickname           string      `json:"nickname"`
		AvatarImgIdStr     string      `json:"avatarImgIdStr"`
		Description        string      `json:"description"`
		Mutual             bool        `json:"mutual"`
		RemarkName         interface{} `json:"remarkName"`
		UserType           int         `json:"userType"`
		AuthStatus         int         `json:"authStatus"`
		DetailDescription  string      `json:"detailDescription"`
		Experts            struct {
		} `json:"experts"`
		ExpertTags      interface{} `json:"expertTags"`
		City            int         `json:"city"`
		DefaultAvatar   bool        `json:"defaultAvatar"`
		BackgroundImgId int         `json:"backgroundImgId"`
		BackgroundUrl   string      `json:"backgroundUrl"`
		AvatarUrl       string      `json:"avatarUrl"`
		Signature       string      `json:"signature"`
		Authority       int         `json:"authority"`
		AllAuthTypes    []struct {
			Type int      `json:"type"`
			Desc string   `json:"desc"`
			Tags []string `json:"tags"`
		} `json:"allAuthTypes"`
		Followeds                 int  `json:"followeds"`
		Follows                   int  `json:"follows"`
		Blacklist                 bool `json:"blacklist"`
		ArtistId                  int  `json:"artistId"`
		EventCount                int  `json:"eventCount"`
		AllSubscribedCount        int  `json:"allSubscribedCount"`
		PlaylistBeSubscribedCount int  `json:"playlistBeSubscribedCount"`
		MainAuthType              struct {
			Type int      `json:"type"`
			Desc string   `json:"desc"`
			Tags []string `json:"tags"`
		} `json:"mainAuthType"`
		AvatarImgIdStr1 string      `json:"avatarImgId_str"`
		FollowTime      interface{} `json:"followTime"`
		FollowMe        bool        `json:"followMe"`
		ArtistIdentity  []int       `json:"artistIdentity"`
		CCount          int         `json:"cCount"`
		SDJPCount       int         `json:"sDJPCount"`
		ArtistName      string      `json:"artistName"`
		PlaylistCount   int         `json:"playlistCount"`
		SCount          int         `json:"sCount"`
		NewFollows      int         `json:"newFollows"`
	} `json:"profile"`
	PeopleCanSeeMyPlayRecord bool `json:"peopleCanSeeMyPlayRecord"`
	Bindings                 []struct {
		UserId       int         `json:"userId"`
		Url          string      `json:"url"`
		ExpiresIn    int         `json:"expiresIn"`
		RefreshTime  int         `json:"refreshTime"`
		BindingTime  int         `json:"bindingTime"`
		TokenJsonStr interface{} `json:"tokenJsonStr"`
		Expired      bool        `json:"expired"`
		Id           int         `json:"id"`
		Type         int         `json:"type"`
	} `json:"bindings"`
	AdValid            bool `json:"adValid"`
	Code               int  `json:"code"`
	CreateTime         int  `json:"createTime"`
	CreateDays         int  `json:"createDays"`
	ProfileVillageInfo struct {
		Title     string      `json:"title"`
		ImageUrl  interface{} `json:"imageUrl"`
		TargetUrl string      `json:"targetUrl"`
	} `json:"profileVillageInfo"`
}
