package api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// RegisterAnonAPI 匿名登录签发 MUSIC_A API
const RegisterAnonAPI = "/api/register/anonimous"

const anonIDXorKey = "3go8&$8*3*3h0k(2)2"

type registerAnonReq struct {
	Username string `json:"username"`
}

func encodeAnonDeviceID(deviceID string) string {
	xored := make([]byte, len(deviceID))
	for i := 0; i < len(deviceID); i++ {
		xored[i] = deviceID[i] ^ anonIDXorKey[i%len(anonIDXorKey)]
	}
	sum := md5.Sum(xored)
	digest := base64.StdEncoding.EncodeToString(sum[:])
	return base64.StdEncoding.EncodeToString([]byte(deviceID + " " + digest))
}

// CreateRegisterAnonReqJson 创建匿名登录请求 json，deviceID 为空时使用 utils.DeviceId
func CreateRegisterAnonReqJson(deviceID string) string {
	if deviceID == "" {
		deviceID = utils.DeviceId
	}
	reqBody, _ := json.Marshal(registerAnonReq{
		Username: encodeAnonDeviceID(deviceID),
	})
	return string(reqBody)
}

// RegisterAnonimous 匿名登录并从响应 Set-Cookie 中提取签发的 MUSIC_A
func RegisterAnonimous(data utils.RequestData, deviceID ...string) (result types.RegisterAnonData, header http.Header, err error) {
	id := utils.DeviceId
	if len(deviceID) > 0 && deviceID[0] != "" {
		id = deviceID[0]
		data.Cookies = append(append([]*http.Cookie(nil), data.Cookies...), &http.Cookie{Name: "deviceId", Value: id})
	}
	options := utils.EapiOption{
		Path: RegisterAnonAPI,
		Url:  "https://music.163.com/eapi/register/anonimous",
		Json: CreateRegisterAnonReqJson(id),
	}
	resBody, header, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, header, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	result.MusicA = extractCookieValue(header, "MUSIC_A")
	return result, header, err
}

func extractCookieValue(header http.Header, name string) string {
	for _, cookie := range header.Values("Set-Cookie") {
		first := strings.SplitN(cookie, ";", 2)[0]
		if k, v, ok := strings.Cut(strings.TrimSpace(first), "="); ok && k == name {
			return v
		}
	}
	return ""
}
