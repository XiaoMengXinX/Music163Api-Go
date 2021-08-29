# 网易云音乐 API-Go

网易云音乐 Golang API 实现

## 灵感来自

[Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)

## 说明

与 [Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) 不同的是，本项目将全部采用 Eapi（即网易云音乐客户端使用的API）

本项目可能会很咕，欢迎各位的 issue 与 pr

## API 调用文档

#### GetSongDetail 获取歌曲详细信息

```go
package main

import (
   "fmt"
   "github.com/XiaoMengXinX/Music163Api-Go/api"
   "github.com/XiaoMengXinX/Music163Api-Go/utils"
)

func main() {
   data := utils.RequestData{}

    result, _ := api.GetSongDetail(data, []int{1295601353}) // 获取 ID:1295601353 的详细信息

   fmt.Println(result.Songs[0].Name) // 打印歌曲名称
}
```

返回

```
生而为人，我很抱歉（VOCALOID）
```

#### GetSongURL 获取歌曲 URL

```go
data := utils.RequestData{
    Cookies: utils.Cookies{
        {
            Key:   "MUSIC_U", // 获取无损音质需填写 Cookies 中的 MUSIC_U
            Value: "984e8c072dc9c670f40d019a3699f326a6d9565cd22e7bcb944dfa4deb7124ae33a649814e309366",
        },
    },
}

result, _ := api.GetSongURL(data, api.SongURLConfig{Ids: []int{1295601353}}) // 获取 ID:1295601353 的歌曲 URL

fmt.Println(result.Data[0].Url) // 打印歌曲 URL
```

返回

```
http://m8.music.126.net/20210829141922/e77fa9b153aaadd41bf5e344ce7c847e/ymusic/0edd/e4e3/4eaf/d2db5cbbef195ff34812eb8c82c83d67.flac
```

#### Batch 批处理

示例：获取当前 Cookie 的用户 ID

```go
data := utils.RequestData{
    Cookies: utils.Cookies{
        {
            Key:   "MUSIC_U",
            Value: "984e8c072dc9c670f40d019a3699f326a6d9565cd22e7bcb944dfa4deb7124ae33a649814e309366",
        },
    },
}

batch := api.Batch{} // 创建 Batch 对象
batch.Init() // Batch 初始化
batch.Add(api.BatchAPI{Key: api.BatchUserSetting, Json: ""}) // 添加要 Batch 的 API (API 列表详见 api/Batch.go 中的 constants)

result, _ := batch.Do(data) // 请求 Batch API

var userData types.BatchUserSettingData
_ = json.Unmarshal([]byte(result), &userData) // Batch 需额外解析 json 数据

fmt.Println(userData.ApiUserSetting.Setting.UserId) // 打印用户 ID
```