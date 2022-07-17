# 网易云音乐 API-Go

网易云音乐 Golang API 实现

## 灵感来自

[Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)

## 说明

与 [Binaryify/NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) 不同的是，本项目将全部采用
Eapi(即网易云音乐客户端使用的API)

本项目可能会很咕，欢迎各位的 issue 与 pr

## API 列表

目前仅实现了少部分API：

- 批处理
- 获取音乐 URL
- 获取音乐歌词
- 获取音乐详细信息
- 获取专辑详细信息
- 获取歌单详细信息
- 获取日推歌曲
- 获取资源评论
- 获取用户歌单
- 新建歌单
- 发送私信（支持图片）
- 发送/删除动态（支持图片）
- 发送/删除 Mlog（仅支持图片）
- 发送/删除评论
- 添加/删除歌单歌曲
- 红心/取消红心歌曲
- 获取 NosToken 以及 上传文件
- QR 登录 (获取 unikey, 检测登录状态)
- 搜索 API (复杂搜索, 多重匹配搜索, 获取搜索建议, 歌曲搜索)
- 短链接生成
- 用户签到
- 获取账号登录状态
- 获取音乐人云豆数量
- 音乐人领取云豆
- 音乐人签到
- 获取音乐人任务列表/周任务列表
- 获取用户详细
- 获取用户信息
- 获取用户设置
- 获取 VIP 信息
- 获取 VIP 成长值
- 领取 VIP 成长值奖励
- 获取/设置用户状态
- 获取歌手详细
- 获取云圈动态
- 获取电台节目
- 分享歌曲回调
- 动态资源分享

## API 调用文档

#### GetSongDetail 获取歌曲详细信息

[在 replit 上试试](https://replit.com/@xibaole2333/Music163Api-Go-demo1)

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

[在 replit 上试试](https://replit.com/@xibaole2333/Music163Api-Go-demo2)

```go
package main

import (
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"net/http"
)

func main() {
	data := utils.RequestData{
		Cookies: []*http.Cookie{
			{
				Name:  "MUSIC_U", // 获取无损音质需填写 Cookies 中的 MUSIC_U
				Value: "8bc6da9b916433c064c742a9ca02f1405cccf02062aac31cdfea4eb6024e0d2248eecaa9668dfe7f43124f3fcebe94e446b14e3f0c3f8af929f5e126cc9926cbc3061cd18d77b7a0",
			},
		},
	}

	result, _ := api.GetSongURL(data, api.SongURLConfig{Ids: []int{1295601353}}) // 获取 ID:1295601353 的歌曲 URL

	fmt.Println(result.Data[0].Url) // 打印歌曲 URL
}
```

返回

```
http://m8.music.126.net/20211031014702/3ace83970e99ef576e4fc350f095382f/ymusic/0edd/e4e3/4eaf/d2db5cbbef195ff34812eb8c82c83d67.flac
```

#### Batch 批处理

示例：获取当前 Cookie 的用户 ID、获取歌曲ID:1416956209 的详细信息

[在 replit 上试试](https://replit.com/@xibaole2333/Music163Api-Go-demo3)

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/api"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"net/http"
)

func main() {
	data := utils.RequestData{
		Cookies: []*http.Cookie{
			{
				Name:  "MUSIC_U",
				Value: "8bc6da9b916433c064c742a9ca02f1405cccf02062aac31cdfea4eb6024e0d2248eecaa9668dfe7f43124f3fcebe94e446b14e3f0c3f8af929f5e126cc9926cbc3061cd18d77b7a0",
			},
		},
	}

	batch := api.NewBatch(api.BatchAPI{Key: api.UserSettingAPI})                                          // 创建初始化 Batch 对象并添加 API
	batch.Add(api.BatchAPI{Key: api.SongDetailAPI, Json: api.CreateSongDetailReqJson([]int{1416956209})}) // 继续添加要批处理的 API

	_, result := batch.Do(data).Parse() // 请求 Batch API 并解析返回的 Json

	var userData types.UserSettingData
	_ = json.Unmarshal([]byte(result[api.UserSettingAPI]), &userData) // 解析 Json 到 struct

	var songDetail types.SongsDetailData
	_ = json.Unmarshal([]byte(result[api.SongDetailAPI]), &songDetail) // 解析 Json 到 struct

	fmt.Println(songDetail.Songs[0].Al.Name) // 打印歌曲专辑名称
	fmt.Println(userData.Setting.UserId)     // 打印 UserID
}
```

Batch 可用的 API 列表详见 https://pkg.go.dev/github.com/XiaoMengXinX/Music163Api-Go/api#pkg-constants

## 已知问题

#### 已解决：

- API 的返回内容解密是多余的，理论上可以移除所有解密相关代码
