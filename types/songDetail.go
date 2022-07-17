package types

// SongsDetailData 获取歌曲详细信息 API 的返回数据
type SongsDetailData struct {
	RawJson    string           `json:"-"`
	Songs      []SongDetailData `json:"songs"`
	Privileges []struct {
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
	} `json:"privileges"`
	Code int `json:"code"`
}

// SongDetailData 获取歌曲详细信息 API 的返回数据
type SongDetailData struct {
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
	Alia []interface{} `json:"alia"`
	Pop  float64       `json:"pop"`
	St   int           `json:"st"`
	Rt   interface{}   `json:"rt"`
	Fee  int           `json:"fee"`
	V    int           `json:"v"`
	Crbt interface{}   `json:"crbt"`
	Cf   string        `json:"cf"`
	Al   struct {
		Id     int           `json:"id"`
		Name   string        `json:"name"`
		PicUrl string        `json:"picUrl"`
		Tns    []interface{} `json:"tns"`
		PicStr string        `json:"pic_str"`
		Pic    int           `json:"pic"`
	} `json:"al"`
	Dt int `json:"dt"`
	H  struct {
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
	OriginSongSimpleData interface{}   `json:"originSongSimpleData"`
	ResourceState        bool          `json:"resourceState"`
	Version              int           `json:"version"`
	Single               int           `json:"single"`
	NoCopyrightRcmd      interface{}   `json:"noCopyrightRcmd"`
	Mv                   int           `json:"mv"`
	Rtype                int           `json:"rtype"`
	Rurl                 interface{}   `json:"rurl"`
	Mst                  int           `json:"mst"`
	Cp                   int           `json:"cp"`
	PublishTime          int           `json:"publishTime"`
}
