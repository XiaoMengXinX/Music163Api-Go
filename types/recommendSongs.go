package types

// RecommendSongsData 获取日推歌曲 API 返回数据类型
type RecommendSongsData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Data    struct {
		DailySongs []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
			Pst  int    `json:"pst"`
			T    int    `json:"t"`
			Ar   []struct {
				Id    int           `json:"id"`
				Name  string        `json:"name"`
				Tns   []interface{} `json:"tns"`
				Alias []interface{} `json:"alias"`
			} `json:"ar"`
			Alia []string    `json:"alia"`
			Pop  float64     `json:"pop"`
			St   int         `json:"st"`
			Rt   *string     `json:"rt"`
			Fee  int         `json:"fee"`
			V    int         `json:"v"`
			Crbt interface{} `json:"crbt"`
			Cf   string      `json:"cf"`
			Al   struct {
				Id     int           `json:"id"`
				Name   string        `json:"name"`
				PicUrl string        `json:"picUrl"`
				Tns    []interface{} `json:"tns"`
				PicStr string        `json:"pic_str"`
				Pic    int64         `json:"pic"`
			} `json:"al"`
			Dt int `json:"dt"`
			H  *struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"h"`
			M struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"m"`
			L struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"l"`
			A                    interface{}   `json:"a"`
			Cd                   string        `json:"cd"`
			No                   int           `json:"no"`
			RtUrl                interface{}   `json:"rtUrl"`
			Ftype                int           `json:"ftype"`
			RtUrls               []interface{} `json:"rtUrls"`
			DjId                 int           `json:"djId"`
			Copyright            int           `json:"copyright"`
			SId                  int           `json:"s_id"`
			Mark                 int           `json:"mark"`
			OriginCoverType      int           `json:"originCoverType"`
			OriginSongSimpleData *struct {
				SongId  int    `json:"songId"`
				Name    string `json:"name"`
				Artists []struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				AlbumMeta struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"albumMeta"`
			} `json:"originSongSimpleData"`
			Single          int         `json:"single"`
			NoCopyrightRcmd interface{} `json:"noCopyrightRcmd"`
			Rurl            interface{} `json:"rurl"`
			Rtype           int         `json:"rtype"`
			Mst             int         `json:"mst"`
			Cp              int         `json:"cp"`
			Mv              int         `json:"mv"`
			PublishTime     int64       `json:"publishTime"`
			Reason          string      `json:"reason"`
			VideoInfo       struct {
				MoreThanOne bool `json:"moreThanOne"`
				Video       *struct {
					Vid         string      `json:"vid"`
					Type        int         `json:"type"`
					Title       *string     `json:"title"`
					PlayTime    int         `json:"playTime"`
					CoverUrl    *string     `json:"coverUrl"`
					PublishTime int64       `json:"publishTime"`
					Artists     interface{} `json:"artists"`
				} `json:"video"`
			} `json:"videoInfo"`
			Privilege struct {
				Id                 int         `json:"id"`
				Fee                int         `json:"fee"`
				Paied              int         `json:"payed"`
				St                 int         `json:"st"`
				Pl                 int         `json:"pl"`
				Dl                 int         `json:"dl"`
				Sp                 int         `json:"sp"`
				Cp                 int         `json:"cp"`
				Subp               int         `json:"subp"`
				Cs                 bool        `json:"cs"`
				Maxbr              int         `json:"maxbr"`
				Fl                 int         `json:"fl"`
				Toast              bool        `json:"toast"`
				Flag               int         `json:"flag"`
				PreSell            bool        `json:"preSell"`
				PlayMaxbr          int         `json:"playMaxbr"`
				DownloadMaxbr      int         `json:"downloadMaxbr"`
				Rscl               interface{} `json:"rscl"`
				FreeTrialPrivilege struct {
					ResConsumable  bool `json:"resConsumable"`
					UserConsumable bool `json:"userConsumable"`
				} `json:"freeTrialPrivilege"`
				ChargeInfoList []struct {
					Rate          int         `json:"rate"`
					ChargeUrl     interface{} `json:"chargeUrl"`
					ChargeMessage interface{} `json:"chargeMessage"`
					ChargeType    int         `json:"chargeType"`
				} `json:"chargeInfoList"`
			} `json:"privilege"`
			Alg string   `json:"alg"`
			Tns []string `json:"tns,omitempty"`
		} `json:"dailySongs"`
		OrderSongs       []interface{} `json:"orderSongs"`
		RecommendReasons []struct {
			SongId int    `json:"songId"`
			Reason string `json:"reason"`
		} `json:"recommendReasons"`
	} `json:"data"`
}
