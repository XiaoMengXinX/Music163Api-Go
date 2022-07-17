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
	Api  SongsURLData `json:"/api/song/enhance/player/url/v1"`
	Code int          `json:"code"`
}

// BatchSongDetailData 批处理 SongDetail 数据
type BatchSongDetailData struct {
	Api  SongsDetailData `json:"/api/v3/song/detail"`
	Code int             `json:"code"`
}

// BatchUserDetailData 批处理 UserDetail 数据
type BatchUserDetailData struct {
	Api  UserDetailData `json:"/api/v1/user/detail"`
	Code int            `json:"code"`
}

// BatchCloudbeanNumData 批处理 CloudbeanNum 数据
type BatchCloudbeanNumData struct {
	Api  CloudBeanNumData `json:"/api/cloudbean/get"`
	Code int              `json:"code"`
}

// BatchMusicianDailyTasksData 批处理 MusicianTasks 数据
type BatchMusicianDailyTasksData struct {
	Api  MusicianDailyTasksData `json:"/api/nmusician/workbench/mission/cycle/list"`
	Code int                    `json:"code"`
}

// BatchMusicianWeeklyTasksData 批处理 MusicianTasks 数据
type BatchMusicianWeeklyTasksData struct {
	Api  MusicianWeeklyTasksData `json:"/api/nmusician/workbench/mission/stage/list"`
	Code int                     `json:"code"`
}

// BatchLoginStatusData 批处理 LoginStatus 数据
type BatchLoginStatusData struct {
	Api  LoginStatusData `json:"/api/w/nuser/account/get"`
	Code int             `json:"code"`
}

// BatchMusicianSignData 批处理 MusicianSign 数据
type BatchMusicianSignData struct {
	Api  MusicianSignData `json:"/api/creator/user/access"`
	Code int              `json:"code"`
}

// BatchObtainCLoudbeanData 批处理 ObtainCLoudbean 数据
type BatchObtainCLoudbeanData struct {
	Api  ObtainCloudebeanData `json:"/api/nmusician/workbench/mission/reward/obtain/new"`
	Code int                  `json:"code"`
}

// BatchUserSignData 批处理 UserSign 数据
type BatchUserSignData struct {
	Api  UserSignData `json:"/api/point/dailyTask"`
	Code int          `json:"code"`
}
