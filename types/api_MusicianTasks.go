package types

// MusicianTasksData 音乐人任务列表数据
type MusicianTasksData struct {
	RawJson string
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
