package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// MusicianTasksAPI 获取音乐人任务列表 API
const MusicianTasksAPI = "/api/nmusician/workbench/mission/cycle/list"

// GetMusicianTasks 获取音乐人任务列表
func GetMusicianTasks(data utils.RequestData) (result types.MusicianTasksData, err error) {
	var options utils.EapiOption
	options.Path = MusicianTasksAPI
	options.Url = "https://music.163.com/eapi/nmusician/workbench/mission/cycle/list"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
