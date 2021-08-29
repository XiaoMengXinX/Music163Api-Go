package types

// BatchUserSettingData 批处理 API 返回的 UserSetting API 数据
type BatchUserSettingData struct {
	Api  UserSettingData `json:"/api/user/setting"`
	Code int             `json:"code"`
}

// BatchUserInfoData 批处理 UserInfo 数据
type BatchUserInfoData struct {
	Api  UserInfoData `json:"/api/v1/user/info"`
	Code int          `json:"code"`
}

// BatchSongURLData 批处理 SongURL 数据
type BatchSongURLData struct {
	Api  SongURLData `json:"/api/song/enhance/player/url/v1"`
	Code int         `json:"code"`
}

// BatchSongDetailData 批处理 SongDetail 数据
type BatchSongDetailData struct {
	Api  SongDetailData `json:"/api/v3/song/detail"`
	Code int            `json:"code"`
}

// BatchUserDetailData 批处理 UserDetailData 数据
type BatchUserDetailData struct {
	Api  UserDetailData `json:"/api/v1/user/detail"`
	Code int            `json:"code"`
}
