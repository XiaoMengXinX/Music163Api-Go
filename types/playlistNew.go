package types

// NewPlaylistData 新建歌单 API 的返回数据
type NewPlaylistData struct {
	RawJson  string `json:"-"`
	Code     int    `json:"code"`
	Playlist struct {
		Subscribers           []interface{} `json:"subscribers"`
		Subscribed            interface{}   `json:"subscribed"`
		Creator               interface{}   `json:"creator"`
		Artists               interface{}   `json:"artists"`
		Tracks                interface{}   `json:"tracks"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		BackgroundCoverId     int           `json:"backgroundCoverId"`
		BackgroundCoverUrl    interface{}   `json:"backgroundCoverUrl"`
		TitleImage            int           `json:"titleImage"`
		TitleImageUrl         interface{}   `json:"titleImageUrl"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		OpRecommend           bool          `json:"opRecommend"`
		RecommendInfo         interface{}   `json:"recommendInfo"`
		SubscribedCount       int           `json:"subscribedCount"`
		CloudTrackCount       int           `json:"cloudTrackCount"`
		UserId                int           `json:"userId"`
		TotalDuration         int           `json:"totalDuration"`
		CoverImgId            int           `json:"coverImgId"`
		Privacy               int           `json:"privacy"`
		TrackUpdateTime       int           `json:"trackUpdateTime"`
		TrackCount            int           `json:"trackCount"`
		UpdateTime            int           `json:"updateTime"`
		CommentThreadId       string        `json:"commentThreadId"`
		CoverImgUrl           string        `json:"coverImgUrl"`
		SpecialType           int           `json:"specialType"`
		Anonimous             bool          `json:"anonimous"`
		CreateTime            int           `json:"createTime"`
		HighQuality           bool          `json:"highQuality"`
		NewImported           bool          `json:"newImported"`
		TrackNumberUpdateTime int           `json:"trackNumberUpdateTime"`
		PlayCount             int           `json:"playCount"`
		AdType                int           `json:"adType"`
		Description           interface{}   `json:"description"`
		Tags                  []interface{} `json:"tags"`
		Ordered               bool          `json:"ordered"`
		Status                int           `json:"status"`
		Name                  string        `json:"name"`
		Id                    int           `json:"id"`
		CoverImgIdStr         string        `json:"coverImgId_str"`
		SharedUsers           interface{}   `json:"sharedUsers"`
		ShareStatus           interface{}   `json:"shareStatus"`
	} `json:"playlist"`
	Id int `json:"id"`
}
