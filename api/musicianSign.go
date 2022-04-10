package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// MusicianSignAPI 音乐人签到 API
const MusicianSignAPI = "/api/creator/user/access"

// MusicianSign 音乐人签到
func MusicianSign(data utils.RequestData) (result types.MusicianSignData, err error) {
	var options utils.EapiOption
	options.Path = MusicianSignAPI
	options.Url = "https://music.163.com/eapi/creator/user/access"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
