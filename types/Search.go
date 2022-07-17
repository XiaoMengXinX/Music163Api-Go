package types

// SearchSuggestData 搜索建议数据
type SearchSuggestData struct {
	RawJson string `json:"-"`
	Result  struct {
		AllMatch []struct {
			Keyword     string `json:"keyword"`
			Type        int    `json:"type"`
			Alg         string `json:"alg"`
			LastKeyword string `json:"lastKeyword"`
			Feature     string `json:"feature"`
		} `json:"allMatch"`
	} `json:"result"`
	Code int `json:"code"`
}

// SearchMultiMatchData 搜索多重匹配返回数据
type SearchMultiMatchData struct {
	RawJson string `json:"-"`
	Result  struct {
		Artist []struct {
			Name         string        `json:"name"`
			Id           int           `json:"id"`
			PicId        int64         `json:"picId"`
			Img1V1Id     int64         `json:"img1v1Id"`
			BriefDesc    string        `json:"briefDesc"`
			PicUrl       string        `json:"picUrl"`
			Img1V1Url    string        `json:"img1v1Url"`
			AlbumSize    int           `json:"albumSize"`
			Alias        []interface{} `json:"alias"`
			Trans        string        `json:"trans"`
			MusicSize    int           `json:"musicSize"`
			PicIdStr     string        `json:"picId_str"`
			Img1V1IdStr  string        `json:"img1v1Id_str"`
			TransNames   interface{}   `json:"transNames"`
			MvSize       int           `json:"mvSize"`
			Alg          string        `json:"alg"`
			FansSize     int           `json:"fansSize"`
			VideoSize    int           `json:"videoSize"`
			OfficialTags []interface{} `json:"officialTags"`
			SearchCircle struct {
				Id           string `json:"id"`
				EntranceText string `json:"entranceText"`
				EntranceUrl  string `json:"entranceUrl"`
			} `json:"searchCircle"`
			Occupation string `json:"occupation"`
		} `json:"artist"`
		Orders []string `json:"orders"`
	} `json:"result"`
	Code int `json:"code"`
}

// SearchComplexData 复杂搜索返回数据
type SearchComplexData struct {
	RawJson string
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ShowType string `json:"showType"`
		Complete struct {
			Song struct {
				MoreText string `json:"moreText"`
				HighText string `json:"highText"`
				Songs    []struct {
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
					Rt   *string       `json:"rt"`
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
						Pic    int64         `json:"pic"`
					} `json:"al"`
					Dt int `json:"dt"`
					H  *struct {
						Br   int     `json:"br"`
						Fid  int     `json:"fid"`
						Size int     `json:"size"`
						Vd   float64 `json:"vd"`
					} `json:"h"`
					M *struct {
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
					Rtype                int           `json:"rtype"`
					Rurl                 interface{}   `json:"rurl"`
					Mst                  int           `json:"mst"`
					Cp                   int           `json:"cp"`
					Mv                   int           `json:"mv"`
					PublishTime          int64         `json:"publishTime"`
					ShowRecommend        bool          `json:"showRecommend"`
					RecommendText        string        `json:"recommendText"`
					VideoInfo            struct {
						MoreThanOne bool `json:"moreThanOne"`
						Video       *struct {
							Vid         string      `json:"vid"`
							Type        int         `json:"type"`
							Title       string      `json:"title"`
							PlayTime    int         `json:"playTime"`
							CoverUrl    string      `json:"coverUrl"`
							PublishTime int64       `json:"publishTime"`
							Artists     interface{} `json:"artists"`
						} `json:"video"`
					} `json:"videoInfo"`
					OfficialTags []interface{} `json:"officialTags"`
					Privilege    struct {
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
					Lyrics      string        `json:"lyrics"`
					Alg         string        `json:"alg"`
					SpecialTags []interface{} `json:"specialTags"`
				} `json:"songs"`
				More       bool `json:"more"`
				KsongInfos struct {
				} `json:"ksongInfos"`
				ResourceIds []int `json:"resourceIds"`
			} `json:"song"`
			Voice struct {
			} `json:"voice"`
			NewMlog struct {
				MoreText  string        `json:"moreText"`
				HighText  string        `json:"highText"`
				More      bool          `json:"more"`
				Resources []interface{} `json:"resources"`
			} `json:"new_mlog"`
			Code     int `json:"code"`
			PlayList struct {
				MoreText  string `json:"moreText"`
				HighText  string `json:"highText"`
				More      bool   `json:"more"`
				PlayLists []struct {
					Id          int64  `json:"id"`
					Name        string `json:"name"`
					CoverImgUrl string `json:"coverImgUrl"`
					Creator     struct {
						Nickname   string      `json:"nickname"`
						UserId     int         `json:"userId"`
						UserType   int         `json:"userType"`
						AvatarUrl  string      `json:"avatarUrl"`
						AuthStatus int         `json:"authStatus"`
						ExpertTags interface{} `json:"expertTags"`
						Experts    interface{} `json:"experts"`
					} `json:"creator"`
					Subscribed   bool     `json:"subscribed"`
					TrackCount   int      `json:"trackCount"`
					UserId       int      `json:"userId"`
					PlayCount    int      `json:"playCount"`
					BookCount    int      `json:"bookCount"`
					SpecialType  int      `json:"specialType"`
					OfficialTags []string `json:"officialTags"`
					Description  *string  `json:"description"`
					HighQuality  bool     `json:"highQuality"`
					Track        struct {
						Name        string        `json:"name"`
						Id          int           `json:"id"`
						Position    int           `json:"position"`
						Alias       []interface{} `json:"alias"`
						Status      int           `json:"status"`
						Fee         int           `json:"fee"`
						CopyrightId int           `json:"copyrightId"`
						Disc        string        `json:"disc"`
						No          int           `json:"no"`
						Artists     []struct {
							Name      string        `json:"name"`
							Id        int           `json:"id"`
							PicId     int           `json:"picId"`
							Img1V1Id  int           `json:"img1v1Id"`
							BriefDesc string        `json:"briefDesc"`
							PicUrl    string        `json:"picUrl"`
							Img1V1Url string        `json:"img1v1Url"`
							AlbumSize int           `json:"albumSize"`
							Alias     []interface{} `json:"alias"`
							Trans     string        `json:"trans"`
							MusicSize int           `json:"musicSize"`
						} `json:"artists"`
						Album struct {
							Name        string  `json:"name"`
							Id          int     `json:"id"`
							Type        string  `json:"type"`
							Size        int     `json:"size"`
							PicId       int64   `json:"picId"`
							BlurPicUrl  string  `json:"blurPicUrl"`
							CompanyId   int     `json:"companyId"`
							Pic         int64   `json:"pic"`
							PicUrl      string  `json:"picUrl"`
							PublishTime int64   `json:"publishTime"`
							Description string  `json:"description"`
							Tags        string  `json:"tags"`
							Company     *string `json:"company"`
							BriefDesc   string  `json:"briefDesc"`
							Artist      struct {
								Name      string        `json:"name"`
								Id        int           `json:"id"`
								PicId     int           `json:"picId"`
								Img1V1Id  int           `json:"img1v1Id"`
								BriefDesc string        `json:"briefDesc"`
								PicUrl    string        `json:"picUrl"`
								Img1V1Url string        `json:"img1v1Url"`
								AlbumSize int           `json:"albumSize"`
								Alias     []interface{} `json:"alias"`
								Trans     string        `json:"trans"`
								MusicSize int           `json:"musicSize"`
							} `json:"artist"`
							Songs           []interface{} `json:"songs"`
							Alias           []interface{} `json:"alias"`
							Status          int           `json:"status"`
							CopyrightId     int           `json:"copyrightId"`
							CommentThreadId string        `json:"commentThreadId"`
							Artists         []struct {
								Name      string        `json:"name"`
								Id        int           `json:"id"`
								PicId     int           `json:"picId"`
								Img1V1Id  int           `json:"img1v1Id"`
								BriefDesc string        `json:"briefDesc"`
								PicUrl    string        `json:"picUrl"`
								Img1V1Url string        `json:"img1v1Url"`
								AlbumSize int           `json:"albumSize"`
								Alias     []interface{} `json:"alias"`
								Trans     string        `json:"trans"`
								MusicSize int           `json:"musicSize"`
							} `json:"artists"`
							PicIdStr string `json:"picId_str,omitempty"`
						} `json:"album"`
						Starred         bool          `json:"starred"`
						Popularity      float64       `json:"popularity"`
						Score           int           `json:"score"`
						StarredNum      int           `json:"starredNum"`
						Duration        int           `json:"duration"`
						PlayedNum       int           `json:"playedNum"`
						DayPlays        int           `json:"dayPlays"`
						HearTime        int           `json:"hearTime"`
						Ringtone        *string       `json:"ringtone"`
						Crbt            interface{}   `json:"crbt"`
						Audition        interface{}   `json:"audition"`
						CopyFrom        string        `json:"copyFrom"`
						CommentThreadId string        `json:"commentThreadId"`
						RtUrl           interface{}   `json:"rtUrl"`
						Ftype           int           `json:"ftype"`
						RtUrls          []interface{} `json:"rtUrls"`
						Copyright       int           `json:"copyright"`
						Rtype           int           `json:"rtype"`
						Rurl            interface{}   `json:"rurl"`
						Mvid            int           `json:"mvid"`
						BMusic          struct {
							Name        interface{} `json:"name"`
							Id          int64       `json:"id"`
							Size        int         `json:"size"`
							Extension   string      `json:"extension"`
							Sr          int         `json:"sr"`
							DfsId       int         `json:"dfsId"`
							Bitrate     int         `json:"bitrate"`
							PlayTime    int         `json:"playTime"`
							VolumeDelta float64     `json:"volumeDelta"`
						} `json:"bMusic"`
						Mp3Url interface{} `json:"mp3Url"`
						HMusic struct {
							Name        interface{} `json:"name"`
							Id          int64       `json:"id"`
							Size        int         `json:"size"`
							Extension   string      `json:"extension"`
							Sr          int         `json:"sr"`
							DfsId       int         `json:"dfsId"`
							Bitrate     int         `json:"bitrate"`
							PlayTime    int         `json:"playTime"`
							VolumeDelta float64     `json:"volumeDelta"`
						} `json:"hMusic"`
						MMusic struct {
							Name        interface{} `json:"name"`
							Id          int64       `json:"id"`
							Size        int         `json:"size"`
							Extension   string      `json:"extension"`
							Sr          int         `json:"sr"`
							DfsId       int         `json:"dfsId"`
							Bitrate     int         `json:"bitrate"`
							PlayTime    int         `json:"playTime"`
							VolumeDelta float64     `json:"volumeDelta"`
						} `json:"mMusic"`
						LMusic struct {
							Name        interface{} `json:"name"`
							Id          int64       `json:"id"`
							Size        int         `json:"size"`
							Extension   string      `json:"extension"`
							Sr          int         `json:"sr"`
							DfsId       int         `json:"dfsId"`
							Bitrate     int         `json:"bitrate"`
							PlayTime    int         `json:"playTime"`
							VolumeDelta float64     `json:"volumeDelta"`
						} `json:"lMusic"`
					} `json:"track"`
					Alg string `json:"alg"`
				} `json:"playLists"`
				ResourceIds []int64 `json:"resourceIds"`
			} `json:"playList"`
			Artist struct {
				Artists []struct {
					Id              int           `json:"id"`
					Name            string        `json:"name"`
					PicUrl          string        `json:"picUrl"`
					Alias           []interface{} `json:"alias"`
					AlbumSize       int           `json:"albumSize"`
					PicId           int64         `json:"picId"`
					Img1V1Url       string        `json:"img1v1Url"`
					AccountId       int           `json:"accountId"`
					Img1V1          int64         `json:"img1v1"`
					IdentityIconUrl string        `json:"identityIconUrl"`
					MvSize          int           `json:"mvSize"`
					Followed        bool          `json:"followed"`
					Alg             string        `json:"alg"`
					Trans           interface{}   `json:"trans"`
				} `json:"artists"`
				More        bool  `json:"more"`
				ResourceIds []int `json:"resourceIds"`
			} `json:"artist"`
			Album struct {
				MoreText string `json:"moreText"`
				Albums   []struct {
					Name        string `json:"name"`
					Id          int    `json:"id"`
					Type        string `json:"type"`
					Size        int    `json:"size"`
					PicId       int64  `json:"picId"`
					BlurPicUrl  string `json:"blurPicUrl"`
					CompanyId   int    `json:"companyId"`
					Pic         int64  `json:"pic"`
					PicUrl      string `json:"picUrl"`
					PublishTime int64  `json:"publishTime"`
					Description string `json:"description"`
					Tags        string `json:"tags"`
					Company     string `json:"company"`
					BriefDesc   string `json:"briefDesc"`
					Artist      struct {
						Name        string        `json:"name"`
						Id          int           `json:"id"`
						PicId       int64         `json:"picId"`
						Img1V1Id    int64         `json:"img1v1Id"`
						BriefDesc   string        `json:"briefDesc"`
						PicUrl      string        `json:"picUrl"`
						Img1V1Url   string        `json:"img1v1Url"`
						AlbumSize   int           `json:"albumSize"`
						Alias       []interface{} `json:"alias"`
						Trans       string        `json:"trans"`
						MusicSize   int           `json:"musicSize"`
						TopicPerson int           `json:"topicPerson"`
						PicIdStr    string        `json:"picId_str"`
						Img1V1IdStr string        `json:"img1v1Id_str"`
						Alia        []interface{} `json:"alia"`
					} `json:"artist"`
					Songs           interface{}   `json:"songs"`
					Alias           []interface{} `json:"alias"`
					Status          int           `json:"status"`
					CopyrightId     int           `json:"copyrightId"`
					CommentThreadId string        `json:"commentThreadId"`
					Artists         []struct {
						Name        string        `json:"name"`
						Id          int           `json:"id"`
						PicId       int           `json:"picId"`
						Img1V1Id    int64         `json:"img1v1Id"`
						BriefDesc   string        `json:"briefDesc"`
						PicUrl      string        `json:"picUrl"`
						Img1V1Url   string        `json:"img1v1Url"`
						AlbumSize   int           `json:"albumSize"`
						Alias       []interface{} `json:"alias"`
						Trans       string        `json:"trans"`
						MusicSize   int           `json:"musicSize"`
						TopicPerson int           `json:"topicPerson"`
						Img1V1IdStr string        `json:"img1v1Id_str"`
					} `json:"artists"`
					Paid     bool   `json:"paid"`
					OnSale   bool   `json:"onSale"`
					PicIdStr string `json:"picId_str,omitempty"`
					Alg      string `json:"alg"`
				} `json:"albums"`
				More        bool  `json:"more"`
				ResourceIds []int `json:"resourceIds"`
			} `json:"album"`
			SimQuery struct {
				SimQuerys []struct {
					Keyword string `json:"keyword"`
					Alg     string `json:"alg"`
				} `json:"sim_querys"`
				More bool `json:"more"`
			} `json:"sim_query"`
			RecType interface{} `json:"rec_type"`
			Circle  struct {
				More      bool `json:"more"`
				Resources []struct {
					ResourceId   string      `json:"resourceId"`
					ResourceType string      `json:"resourceType"`
					InternalType interface{} `json:"internalType"`
					ResourceName interface{} `json:"resourceName"`
					BaseInfo     struct {
						Circle struct {
							Id             string `json:"id"`
							Name           string `json:"name"`
							Introduction   string `json:"introduction"`
							Image          string `json:"image"`
							Type           int    `json:"type"`
							Authentication int    `json:"authentication"`
							MemberName     string `json:"memberName"`
							Artist         struct {
								Id   int    `json:"id"`
								Name string `json:"name"`
							} `json:"artist"`
							CreateTime string `json:"createTime"`
							Count      struct {
								Post      string `json:"post"`
								Read      string `json:"read"`
								Member    string `json:"member"`
								ToDayPost string `json:"toDayPost"`
							} `json:"count"`
						} `json:"circle"`
						Airborne            interface{} `json:"airborne"`
						Interactions        interface{} `json:"interactions"`
						OrpheusUrl          string      `json:"orpheusUrl"`
						OrpheusText         interface{} `json:"orpheusText"`
						IconUrl             interface{} `json:"iconUrl"`
						CircleExtensionInfo interface{} `json:"circleExtensionInfo"`
					} `json:"baseInfo"`
					Alg string `json:"alg"`
				} `json:"resources"`
			} `json:"circle"`
			RecQuery  []interface{} `json:"rec_query"`
			Voicelist struct {
			} `json:"voicelist"`
			User struct {
				MoreText string `json:"moreText"`
				More     bool   `json:"more"`
				Users    []struct {
					DefaultAvatar       bool        `json:"defaultAvatar"`
					Province            int         `json:"province"`
					AuthStatus          int         `json:"authStatus"`
					Followed            bool        `json:"followed"`
					AvatarUrl           string      `json:"avatarUrl"`
					AccountStatus       int         `json:"accountStatus"`
					Gender              int         `json:"gender"`
					City                int         `json:"city"`
					Birthday            int64       `json:"birthday"`
					UserId              int         `json:"userId"`
					UserType            int         `json:"userType"`
					Nickname            string      `json:"nickname"`
					Signature           string      `json:"signature"`
					Description         string      `json:"description"`
					DetailDescription   string      `json:"detailDescription"`
					AvatarImgId         int64       `json:"avatarImgId"`
					BackgroundImgId     int64       `json:"backgroundImgId"`
					BackgroundUrl       string      `json:"backgroundUrl"`
					Authority           int         `json:"authority"`
					Mutual              bool        `json:"mutual"`
					ExpertTags          interface{} `json:"expertTags"`
					Experts             interface{} `json:"experts"`
					DjStatus            int         `json:"djStatus"`
					VipType             int         `json:"vipType"`
					RemarkName          interface{} `json:"remarkName"`
					AuthenticationTypes int         `json:"authenticationTypes"`
					AvatarDetail        struct {
						UserType        int    `json:"userType"`
						IdentityLevel   int    `json:"identityLevel"`
						IdentityIconUrl string `json:"identityIconUrl"`
					} `json:"avatarDetail"`
					AvatarImgIdStr     string `json:"avatarImgIdStr"`
					BackgroundImgIdStr string `json:"backgroundImgIdStr"`
					Anchor             bool   `json:"anchor"`
					AvatarImgIdStr1    string `json:"avatarImgId_str"`
					Alg                string `json:"alg"`
				} `json:"users"`
				ResourceIds []int `json:"resourceIds"`
			} `json:"user"`
			Order []string `json:"order"`
		} `json:"complete"`
		Accurate interface{} `json:"accurate"`
	} `json:"data"`
}

// SearchSongData 音乐搜索数据
type SearchSongData struct {
	RawJson string
	Result  struct {
		Songs []struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				Id        int           `json:"id"`
				Name      string        `json:"name"`
				PicUrl    interface{}   `json:"picUrl"`
				Alias     []interface{} `json:"alias"`
				AlbumSize int           `json:"albumSize"`
				PicId     int           `json:"picId"`
				Img1V1Url string        `json:"img1v1Url"`
				Img1V1    int           `json:"img1v1"`
				Trans     interface{}   `json:"trans"`
			} `json:"artists"`
			Album struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Artist struct {
					Id        int           `json:"id"`
					Name      string        `json:"name"`
					PicUrl    interface{}   `json:"picUrl"`
					Alias     []interface{} `json:"alias"`
					AlbumSize int           `json:"albumSize"`
					PicId     int           `json:"picId"`
					Img1V1Url string        `json:"img1v1Url"`
					Img1V1    int           `json:"img1v1"`
					Trans     interface{}   `json:"trans"`
				} `json:"artist"`
				PublishTime int64 `json:"publishTime"`
				Size        int   `json:"size"`
				CopyrightId int   `json:"copyrightId"`
				Status      int   `json:"status"`
				PicId       int64 `json:"picId"`
				Mark        int   `json:"mark"`
			} `json:"album"`
			Duration    int           `json:"duration"`
			CopyrightId int           `json:"copyrightId"`
			Status      int           `json:"status"`
			Alias       []interface{} `json:"alias"`
			Rtype       int           `json:"rtype"`
			Ftype       int           `json:"ftype"`
			Mvid        int           `json:"mvid"`
			Fee         int           `json:"fee"`
			RUrl        interface{}   `json:"rUrl"`
			Mark        int           `json:"mark"`
		} `json:"songs"`
		HasMore   bool `json:"hasMore"`
		SongCount int  `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}
