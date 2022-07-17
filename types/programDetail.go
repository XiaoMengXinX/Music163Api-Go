package types

type ProgramDetailData struct {
	RawJson string `json:"-"`
	Program struct {
		MainSong                 ProgramSong   `json:"mainSong"`
		Songs                    []ProgramSong `json:"songs"`
		Dj                       ProgramDJ     `json:"dj"`
		BlurCoverURL             string        `json:"blurCoverUrl"`
		Duration                 int           `json:"duration"`
		Buyed                    bool          `json:"buyed"`
		ProgramDesc              interface{}   `json:"programDesc"`
		H5Links                  interface{}   `json:"h5Links"`
		CanReward                bool          `json:"canReward"`
		AuditStatus              int           `json:"auditStatus"`
		VideoInfo                interface{}   `json:"videoInfo"`
		Score                    int           `json:"score"`
		LiveInfo                 interface{}   `json:"liveInfo"`
		Alg                      interface{}   `json:"alg"`
		DisPlayStatus            interface{}   `json:"disPlayStatus"`
		AuditDisPlayStatus       int           `json:"auditDisPlayStatus"`
		CategoryName             string        `json:"categoryName"`
		SecondCategoryName       string        `json:"secondCategoryName"`
		ExistLyric               bool          `json:"existLyric"`
		DjPlayRecordVo           interface{}   `json:"djPlayRecordVo"`
		Recommended              bool          `json:"recommended"`
		TrackCount               int           `json:"trackCount"`
		Channels                 []interface{} `json:"channels"`
		CategoryID               int           `json:"categoryId"`
		CreateTime               int64         `json:"createTime"`
		CreateEventID            int           `json:"createEventId"`
		ListenerCount            int           `json:"listenerCount"`
		ScheduledPublishTime     int64         `json:"scheduledPublishTime"`
		SerialNum                int           `json:"serialNum"`
		CoverID                  int64         `json:"coverId"`
		CoverURL                 string        `json:"coverUrl"`
		PubStatus                int           `json:"pubStatus"`
		BdAuditStatus            int           `json:"bdAuditStatus"`
		SecondCategoryID         int           `json:"secondCategoryId"`
		SmallLanguageAuditStatus int           `json:"smallLanguageAuditStatus"`
		TitbitImages             interface{}   `json:"titbitImages"`
		IsPublish                bool          `json:"isPublish"`
		MainTrackID              int           `json:"mainTrackId"`
		ProgramFeeType           int           `json:"programFeeType"`
		Titbits                  interface{}   `json:"titbits"`
		FeeScope                 int           `json:"feeScope"`
		SubscribedCount          int           `json:"subscribedCount"`
		Reward                   bool          `json:"reward"`
		CommentThreadID          string        `json:"commentThreadId"`
		Privacy                  bool          `json:"privacy"`
		Description              string        `json:"description"`
		Name                     string        `json:"name"`
		ID                       int64         `json:"id"`
		Subscribed               bool          `json:"subscribed"`
		ShareCount               int           `json:"shareCount"`
		LikedCount               int           `json:"likedCount"`
		CommentCount             int           `json:"commentCount"`
	} `json:"program"`
	Code int `json:"code"`
}

type ProgramArtist struct {
	Name        string        `json:"name"`
	ID          int           `json:"id"`
	PicID       int           `json:"picId"`
	Img1V1ID    int           `json:"img1v1Id"`
	BriefDesc   string        `json:"briefDesc"`
	PicURL      string        `json:"picUrl"`
	Img1V1URL   string        `json:"img1v1Url"`
	AlbumSize   int           `json:"albumSize"`
	Alias       []interface{} `json:"alias"`
	Trans       string        `json:"trans"`
	MusicSize   int           `json:"musicSize"`
	TopicPerson int           `json:"topicPerson"`
}

type ProgramAlbum struct {
	Name            string        `json:"name"`
	ID              int           `json:"id"`
	Type            interface{}   `json:"type"`
	Size            int           `json:"size"`
	PicID           int64         `json:"picId"`
	BlurPicURL      string        `json:"blurPicUrl"`
	CompanyID       int           `json:"companyId"`
	Pic             int64         `json:"pic"`
	PicURL          string        `json:"picUrl"`
	PublishTime     int           `json:"publishTime"`
	Description     string        `json:"description"`
	Tags            string        `json:"tags"`
	Company         interface{}   `json:"company"`
	BriefDesc       string        `json:"briefDesc"`
	Artist          ProgramArtist `json:"artist"`
	Songs           []interface{} `json:"songs"`
	Alias           []interface{} `json:"alias"`
	Status          int           `json:"status"`
	CopyrightID     int           `json:"copyrightId"`
	CommentThreadID string        `json:"commentThreadId"`
	SubType         interface{}   `json:"subType"`
	TransName       interface{}   `json:"transName"`
	Mark            int           `json:"mark"`
	PicIDStr        string        `json:"picId_str"`
}

type ProgramSong struct {
	Name            string          `json:"name"`
	ID              int             `json:"id"`
	Position        int             `json:"position"`
	Alias           []interface{}   `json:"alias"`
	Status          int             `json:"status"`
	Fee             int             `json:"fee"`
	CopyrightID     int             `json:"copyrightId"`
	Disc            string          `json:"disc"`
	No              int             `json:"no"`
	Artists         []ProgramArtist `json:"artists"`
	Album           ProgramAlbum    `json:"album"`
	Starred         bool            `json:"starred"`
	Popularity      float64         `json:"popularity"`
	Score           int             `json:"score"`
	StarredNum      int             `json:"starredNum"`
	Duration        int             `json:"duration"`
	PlayedNum       int             `json:"playedNum"`
	DayPlays        int             `json:"dayPlays"`
	HearTime        int             `json:"hearTime"`
	Ringtone        string          `json:"ringtone"`
	Crbt            interface{}     `json:"crbt"`
	Audition        interface{}     `json:"audition"`
	CopyFrom        string          `json:"copyFrom"`
	CommentThreadID string          `json:"commentThreadId"`
	RtURL           interface{}     `json:"rtUrl"`
	Ftype           int             `json:"ftype"`
	RtUrls          []interface{}   `json:"rtUrls"`
	Copyright       int             `json:"copyright"`
	TransName       string          `json:"transName"`
	Sign            interface{}     `json:"sign"`
	Mark            int             `json:"mark"`
	NoCopyrightRcmd interface{}     `json:"noCopyrightRcmd"`
	Mvid            int             `json:"mvid"`
	Rtype           int             `json:"rtype"`
	Rurl            interface{}     `json:"rurl"`
	TransNames      []string        `json:"transNames"`
}

type ProgramDJ struct {
	DefaultAvatar       bool        `json:"defaultAvatar"`
	Province            int         `json:"province"`
	AuthStatus          int         `json:"authStatus"`
	Followed            bool        `json:"followed"`
	AvatarURL           string      `json:"avatarUrl"`
	AccountStatus       int         `json:"accountStatus"`
	Gender              int         `json:"gender"`
	City                int         `json:"city"`
	Birthday            int64       `json:"birthday"`
	UserID              int         `json:"userId"`
	UserType            int         `json:"userType"`
	Nickname            string      `json:"nickname"`
	Signature           string      `json:"signature"`
	Description         string      `json:"description"`
	DetailDescription   string      `json:"detailDescription"`
	AvatarImgID         int64       `json:"avatarImgId"`
	BackgroundImgID     int64       `json:"backgroundImgId"`
	BackgroundURL       string      `json:"backgroundUrl"`
	Authority           int         `json:"authority"`
	Mutual              bool        `json:"mutual"`
	ExpertTags          interface{} `json:"expertTags"`
	Experts             interface{} `json:"experts"`
	DjStatus            int         `json:"djStatus"`
	VipType             int         `json:"vipType"`
	RemarkName          interface{} `json:"remarkName"`
	AuthenticationTypes int         `json:"authenticationTypes"`
	AvatarDetail        interface{} `json:"avatarDetail"`
	AvatarImgIDStr      string      `json:"avatarImgIdStr,avatarImgId_str"`
	BackgroundImgIDStr  string      `json:"backgroundImgIdStr"`
	Anchor              bool        `json:"anchor"`
	Brand               string      `json:"brand"`
}
