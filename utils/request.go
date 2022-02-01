package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
	"time"
)

// DEBUG 是否开启打印调试信息
var DEBUG bool

// EapiRequest 返回内容加密的 API 请求
func EapiRequest(eapiOption EapiOption, options RequestData) (result string, header http.Header, err error) {
	data := SpliceStr(eapiOption.Path, eapiOption.Json)
	answer, header, err := CreateNewRequest(Format2Params(data), eapiOption.Url, options)
	if err == nil {
		decrypted := EapiDecrypt([]byte(answer))
		if DEBUG {
			log.Printf("[RespBodyJson]: %s", string(decrypted))
			log.Printf("[RespHeader]: %s", header)
		}
		return string(decrypted), header, nil
	}
	return "", header, err
}

// ApiRequest 返回内容未加密的 API 请求
func ApiRequest(eapiOption EapiOption, options RequestData) (result string, header http.Header, err error) {
	data := SpliceStr(eapiOption.Path, eapiOption.Json)
	answer, header, err := CreateNewRequest(Format2Params(data), eapiOption.Url, options)
	if err == nil {
		if DEBUG {
			log.Printf("[RespBodyJson]: %s", answer)
			log.Printf("[RespHeader]: %s", header)
		}
		return answer, header, nil
	}
	return "", header, err
}

// RawRequest 用于上传文件
func RawRequest(url string, options RequestData) (result string, err error) {
	answer, _, err := CreateNewRequest(options.Body, url, options)
	if err == nil {
		if DEBUG {
			log.Printf("[RespBody]: %s", answer)
		}
		return answer, nil
	}
	return "", err
}

// SpliceStr 拼接字符串
func SpliceStr(path string, data string) (result string) {
	nobodyKnowThis := "36cd479b6b5"
	text := fmt.Sprintf("nobody%suse%smd5forencrypt", path, data)
	MD5 := md5.Sum([]byte(text))
	md5str := fmt.Sprintf("%x", MD5)
	result = fmt.Sprintf("%s-%s-%s-%s-%s", path, nobodyKnowThis, data, nobodyKnowThis, md5str)
	return result
}

// Format2Params 拼接字符串
func Format2Params(str string) (data string) {
	data = fmt.Sprintf("params=%X", EapiEncrypt(str))
	return data
}

// ChooseUserAgent 随机 UserAgent
func ChooseUserAgent() string {
	userAgentList := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"NeteaseMusic/6.5.0.1575377963(164);Dalvik/2.1.0 (Linux; U; Android 9; MIX 2 MIUI/V12.0.1.0.PDECNXM)",
	}
	rand.Seed(time.Now().UnixNano())
	var index int
	index = rand.Intn(len(userAgentList))
	return userAgentList[index]
}

func encodeURIComponent(str string) string {
	r := netUrl.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

// CreateNewRequest 创建 eapi 请求
func CreateNewRequest(data string, url string, options RequestData) (answer string, resHeader http.Header, err error) {
	client := &http.Client{}
	reqBody := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", resHeader, err
	}

	cookie := map[string]interface{}{}
	for _, v := range options.Cookies {
		cookie[v.Name] = v.Value
	}

	for _, v := range options.Headers {
		req.Header.Set(v.Name, v.Value)
	}

	cookie["appver"] = "6.5.0"
	cookie["versioncode"] = "164"
	cookie["buildver"] = strconv.FormatInt(time.Now().Unix(), 10)[0:10]
	cookie["resolution"] = "1920x1080"
	cookie["os"] = "android"

	_, ok := cookie["MUSIC_U"]
	if !ok {
		_, ok := cookie["MUSIC_A"]
		if !ok {
			cookie["MUSIC_A"] = "4ee5f776c9ed1e4d5f031b09e084c6cb333e43ee4a841afeebbef9bbf4b7e4152b51ff20ecb9e8ee9e89ab23044cf50d1609e4781e805e73a138419e5583bc7fd1e5933c52368d9127ba9ce4e2f233bf5a77ba40ea6045ae1fc612ead95d7b0e0edf70a74334194e1a190979f5fc12e9968c3666a981495b33a649814e309366"
		}
	}

	var cookies string
	for key, val := range cookie {
		cookies += encodeURIComponent(key) + "=" + encodeURIComponent(fmt.Sprintf("%v", val)) + "; "
	}
	req.Header.Set("Cookie", strings.TrimRight(cookies, "; "))

	if len(req.Header["Content-Type"]) == 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Set("User-Agent", ChooseUserAgent())

	if DEBUG {
		log.Printf("[Request]: %+v", req)
		if len([]byte(data)) < 51200 {
			log.Printf("[ReqBody]: %+v", data)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", resHeader, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resHeader, err
	}

	return string(body), resp.Header, nil
}
