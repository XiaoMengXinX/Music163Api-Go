package types

// UserDetailData 用户详细信息数据
type UserDetailData struct {
	RawJson  string `json:"-"`
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
	CurrentProduct interface{} `json:"currentProduct"`
	Products       []struct {
		ProductionTypeName string `json:"productionTypeName"`
		ProductionTypeId   int    `json:"productionTypeId"`
	} `json:"products"`
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
		PrivacyItemUnlimit struct {
			Area       bool `json:"area"`
			College    bool `json:"college"`
			Age        bool `json:"age"`
			VillageAge bool `json:"villageAge"`
		} `json:"privacyItemUnlimit"`
		AvatarDetail struct {
			UserType        interface{} `json:"userType"`
			IdentityLevel   int         `json:"identityLevel"`
			IdentityIconUrl string      `json:"identityIconUrl"`
		} `json:"avatarDetail"`
		AvatarImgIdStr     string      `json:"avatarImgIdStr"`
		BackgroundImgIdStr string      `json:"backgroundImgIdStr"`
		Mutual             bool        `json:"mutual"`
		RemarkName         interface{} `json:"remarkName"`
		CreateTime         int64       `json:"createTime"`
		Birthday           int64       `json:"birthday"`
		AvatarImgId        int64       `json:"avatarImgId"`
		Province           int         `json:"province"`
		City               int         `json:"city"`
		Gender             int         `json:"gender"`
		Followed           bool        `json:"followed"`
		Nickname           string      `json:"nickname"`
		AvatarUrl          string      `json:"avatarUrl"`
		BackgroundImgId    int64       `json:"backgroundImgId"`
		BackgroundUrl      string      `json:"backgroundUrl"`
		UserType           int         `json:"userType"`
		DefaultAvatar      bool        `json:"defaultAvatar"`
		DjStatus           int         `json:"djStatus"`
		AccountStatus      int         `json:"accountStatus"`
		VipType            int         `json:"vipType"`
		AuthStatus         int         `json:"authStatus"`
		DetailDescription  string      `json:"detailDescription"`
		Experts            struct {
		} `json:"experts"`
		ExpertTags   interface{} `json:"expertTags"`
		Description  string      `json:"description"`
		UserId       int         `json:"userId"`
		Signature    string      `json:"signature"`
		Authority    int         `json:"authority"`
		AllAuthTypes []struct {
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
		InBlacklist     bool        `json:"inBlacklist"`
		SDJPCount       int         `json:"sDJPCount"`
		ArtistName      string      `json:"artistName"`
		PlaylistCount   int         `json:"playlistCount"`
		SCount          int         `json:"sCount"`
		NewFollows      int         `json:"newFollows"`
	} `json:"profile"`
	PeopleCanSeeMyPlayRecord bool `json:"peopleCanSeeMyPlayRecord"`
	Bindings                 []struct {
		ExpiresIn    int         `json:"expiresIn"`
		RefreshTime  int         `json:"refreshTime"`
		BindingTime  int64       `json:"bindingTime"`
		TokenJsonStr interface{} `json:"tokenJsonStr"`
		Url          string      `json:"url"`
		Expired      bool        `json:"expired"`
		UserId       int         `json:"userId"`
		Id           int64       `json:"id"`
		Type         int         `json:"type"`
	} `json:"bindings"`
	AdValid            bool  `json:"adValid"`
	Code               int   `json:"code"`
	CreateTime         int64 `json:"createTime"`
	CreateDays         int   `json:"createDays"`
	ProfileVillageInfo struct {
		Title     string `json:"title"`
		ImageUrl  string `json:"imageUrl"`
		TargetUrl string `json:"targetUrl"`
	} `json:"profileVillageInfo"`
}
