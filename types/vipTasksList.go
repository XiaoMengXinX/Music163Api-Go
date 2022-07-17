package types

// VipTaskslistDetailData VIP 任务列表数据
type VipTaskslistDetailData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TaskList []struct {
			Seq       int    `json:"seq"`
			SeqName   string `json:"seqName"`
			TaskItems []struct {
				Type            int         `json:"type"`
				TypeCode        interface{} `json:"typeCode"`
				SeqTag          int         `json:"seqTag"`
				Action          string      `json:"action"`
				Description     string      `json:"description"`
				GrowthPoint     int         `json:"growthPoint"`
				Status          int         `json:"status"`
				BasicTaskId     int         `json:"basicTaskId"`
				TaskId          string      `json:"taskId"`
				TaskTag         string      `json:"taskTag"`
				TagId           int         `json:"tagId"`
				Url             string      `json:"url"`
				IconUrl         string      `json:"iconUrl"`
				SortValue       int         `json:"sortValue"`
				TotalUngetScore int         `json:"totalUngetScore"`
				UnGetIds        []string    `json:"unGetIds"`
				RuleId          int         `json:"ruleId"`
				ActionType      int         `json:"actionType"`
				ProgressType    int         `json:"progressType"`
				TargetWorth     int         `json:"targetWorth"`
				Period          int         `json:"period"`
				CurrentProgress int         `json:"currentProgress"`
				ShowProgress    bool        `json:"showProgress"`
				NeedReceive     bool        `json:"needReceive"`
				Targets         interface{} `json:"targets"`
				Name            string      `json:"name"`
			} `json:"taskItems"`
			TaskType int `json:"taskType"`
		} `json:"taskList"`
		TaskScore int `json:"taskScore"`
	} `json:"data"`
}
