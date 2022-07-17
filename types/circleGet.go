package types

// GetCircleData 获取云圈动态API返回数据
type GetCircleData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
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
				Post   string `json:"post"`
				Read   string `json:"read"`
				Member string `json:"member"`
				Online string `json:"online"`
			} `json:"count"`
			ActivityLabel    bool        `json:"activityLabel"`
			Source           int         `json:"source"`
			Followed         interface{} `json:"followed"`
			UserType         int         `json:"userType"`
			CircleOuterId    string      `json:"circleOuterId"`
			CircleOuterName  string      `json:"circleOuterName"`
			CircleOrpheusUrl string      `json:"circleOrpheusUrl"`
			BtnName          interface{} `json:"btnName"`
			BtnOrpheus       interface{} `json:"btnOrpheus"`
			IsCircleSquare   bool        `json:"isCircleSquare"`
			SquareButton     interface{} `json:"squareButton"`
			PrivateCircle    bool        `json:"privateCircle"`
		} `json:"circle"`
		Airborne []string `json:"airborne"`
		ShareUrl string   `json:"shareUrl"`
		Tabs     []struct {
			Type              int           `json:"type"`
			Name              string        `json:"name"`
			Count             string        `json:"count"`
			HasNew            bool          `json:"hasNew"`
			Url               *string       `json:"url"`
			TopInfos          []interface{} `json:"topInfos"`
			ArtistOnly        *bool         `json:"artistOnly"`
			ArtistOnlyText    *string       `json:"artistOnlyText"`
			DefaultFilterType *int          `json:"defaultFilterType"`
			Extension         interface{}   `json:"extension"`
		} `json:"tabs"`
		V2Tabs     interface{} `json:"v2Tabs"`
		DefaultTab int         `json:"defaultTab"`
		Orpheus    struct {
			Text  string      `json:"text"`
			Url   string      `json:"url"`
			H5Url string      `json:"h5Url"`
			Game  interface{} `json:"game"`
			Type  int         `json:"type"`
		} `json:"orpheus"`
		Members struct {
			AvatarUrls []string `json:"avatarUrls"`
			Url        string   `json:"url"`
		} `json:"members"`
		CountInfo struct {
			OnlineCount string `json:"onlineCount"`
			NumDesc     string `json:"numDesc"`
		} `json:"countInfo"`
		ManageEntrance struct {
			Show       bool   `json:"show"`
			OrpheusUrl string `json:"orpheusUrl"`
		} `json:"manageEntrance"`
		SettingOrpheus    string `json:"settingOrpheus"`
		Official          bool   `json:"official"`
		PublishInfoLayout struct {
			PublishInfos []struct {
				PublishType    int    `json:"publishType"`
				PublishText    string `json:"publishText"`
				PublishOrpheus string `json:"publishOrpheus"`
				PublishIcon    string `json:"publishIcon"`
			} `json:"publishInfos"`
		} `json:"publishInfoLayout"`
		PublishInfo struct {
			PublishType    int         `json:"publishType"`
			PublishText    interface{} `json:"publishText"`
			PublishOrpheus interface{} `json:"publishOrpheus"`
			PublishIcon    string      `json:"publishIcon"`
		} `json:"publishInfo"`
		RecommendModel struct {
			IsOpen   bool   `json:"isOpen"`
			Title    string `json:"title"`
			SubTitle string `json:"subTitle"`
			Action   struct {
				Title string      `json:"title"`
				Link  interface{} `json:"link"`
			} `json:"action"`
		} `json:"recommendModel"`
		ModalVo interface{} `json:"modalVo"`
	} `json:"data"`
}
