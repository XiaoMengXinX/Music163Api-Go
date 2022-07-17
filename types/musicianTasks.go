package types

// MusicianDailyTasksData 音乐人每日任务列表
type MusicianDailyTasksData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Business        string `json:"business"`
			MissionId       int    `json:"missionId"`
			UserId          int    `json:"userId"`
			MissionEntityId int    `json:"missionEntityId"`
			RewardId        int    `json:"rewardId"`
			ProgressRate    int    `json:"progressRate"`
			Description     string `json:"description"`
			Type            int    `json:"type"`
			Tag             int    `json:"tag"`
			ActionType      int    `json:"actionType"`
			Platform        int    `json:"platform"`
			Status          int    `json:"status"`
			Button          string `json:"button"`
			SortValue       int    `json:"sortValue"`
			StartTime       int    `json:"startTime"`
			EndTime         int    `json:"endTime"`
			ExtendInfo      string `json:"extendInfo"`
			Period          int    `json:"period"`
			NeedToReceive   int    `json:"needToReceive,omitempty"`
			TargetCount     int    `json:"targetCount"`
			RewardWorth     string `json:"rewardWorth"`
			RewardType      int    `json:"rewardType"`
			UserMissionId   int    `json:"userMissionId,omitempty"`
			UpdateTime      int    `json:"updateTime,omitempty"`
		} `json:"list"`
	} `json:"data"`
}

// MusicianWeeklyTasksData 音乐人周任务列表
type MusicianWeeklyTasksData struct {
	RawJson string
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Business            string `json:"business"`
			MissionId           int    `json:"missionId"`
			UserId              int    `json:"userId"`
			MissionEntityId     int    `json:"missionEntityId"`
			RewardId            int    `json:"rewardId"`
			ProgressRate        int    `json:"progressRate"`
			Description         string `json:"description"`
			Type                int    `json:"type"`
			Tag                 int    `json:"tag"`
			ActionType          int    `json:"actionType"`
			Platform            int    `json:"platform"`
			Status              int    `json:"status"`
			Button              string `json:"button"`
			SortValue           int    `json:"sortValue"`
			StartTime           int64  `json:"startTime"`
			EndTime             int64  `json:"endTime"`
			ExtendInfo          string `json:"extendInfo"`
			Period              int    `json:"period"`
			NeedToReceive       int    `json:"needToReceive,omitempty"`
			UserStageTargetList []struct {
				SortValue     int   `json:"sortValue"`
				StageType     int   `json:"stageType"`
				Times         int   `json:"times"`
				Value         int   `json:"value"`
				ProgressRate  int   `json:"progressRate"`
				SumTarget     int   `json:"sumTarget"`
				RewardId      int   `json:"rewardId"`
				RewardType    int   `json:"rewardType"`
				Worth         int   `json:"worth"`
				Status        int   `json:"status"`
				UserMissionId int64 `json:"userMissionId,omitempty"`
			} `json:"userStageTargetList"`
			HasSendWorth  int   `json:"hasSendWorth"`
			UserMissionId int64 `json:"userMissionId,omitempty"`
			UpdateTime    int64 `json:"updateTime,omitempty"`
		} `json:"list"`
	} `json:"data"`
}
