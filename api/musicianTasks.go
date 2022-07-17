package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// MusicianTasksDailyAPI 获取音乐人每日任务 API
const MusicianTasksDailyAPI = "/api/nmusician/workbench/mission/cycle/list"

// MusicianTasksWeeklyAPI 获取音乐人周任务 API
const MusicianTasksWeeklyAPI = "/api/nmusician/workbench/mission/stage/list"

// GetMusicianDailyTasks 获取音乐人每日任务
func GetMusicianDailyTasks(data utils.RequestData) (result types.MusicianDailyTasksData, err error) {
	var options utils.EapiOption
	options.Path = MusicianTasksDailyAPI
	options.Url = "https://music.163.com/eapi/nmusician/workbench/mission/cycle/list"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// GetMusicianWeeklyTasks 获取音乐人周任务
func GetMusicianWeeklyTasks(data utils.RequestData) (result types.MusicianWeeklyTasksData, err error) {
	var options utils.EapiOption
	options.Path = MusicianTasksWeeklyAPI
	options.Url = "https://music.163.com/eapi/nmusician/workbench/mission/stage/list"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
