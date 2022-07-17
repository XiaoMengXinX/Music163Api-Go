package types

// AddCommentData 新增评论返回数据
type AddCommentData struct {
	RawJson          string `json:"-"`
	Code             int    `json:"code"`
	MusicianSaidTips struct {
		Toast          string `json:"toast"`
		InviteTitle    string `json:"inviteTitle"`
		InviteContent  string `json:"inviteContent"`
		ShowInvitation bool   `json:"showInvitation"`
	} `json:"musicianSaidTips"`
	Comment struct {
		User struct {
			LocationInfo interface{} `json:"locationInfo"`
			LiveInfo     interface{} `json:"liveInfo"`
			Anonym       int         `json:"anonym"`
			UserType     int         `json:"userType"`
			AuthStatus   int         `json:"authStatus"`
			ExpertTags   interface{} `json:"expertTags"`
			VipType      int         `json:"vipType"`
			RemarkName   interface{} `json:"remarkName"`
			AvatarUrl    string      `json:"avatarUrl"`
			Experts      interface{} `json:"experts"`
			VipRights    struct {
				Associator   interface{} `json:"associator"`
				MusicPackage struct {
					VipCode int  `json:"vipCode"`
					Rights  bool `json:"rights"`
				} `json:"musicPackage"`
				RedVipAnnualCount int `json:"redVipAnnualCount"`
				RedVipLevel       int `json:"redVipLevel"`
			} `json:"vipRights"`
			Nickname     string `json:"nickname"`
			UserId       int    `json:"userId"`
			AvatarDetail struct {
				UserType        int    `json:"userType"`
				IdentityLevel   int    `json:"identityLevel"`
				IdentityIconUrl string `json:"identityIconUrl"`
			} `json:"avatarDetail"`
		} `json:"user"`
		BeReplied           []interface{} `json:"beReplied"`
		PendantData         interface{}   `json:"pendantData"`
		ShowFloorComment    interface{}   `json:"showFloorComment"`
		Status              int           `json:"status"`
		CommentId           int           `json:"commentId"`
		Content             string        `json:"content"`
		Time                int           `json:"time"`
		LikedCount          int           `json:"likedCount"`
		ExpressionUrl       interface{}   `json:"expressionUrl"`
		CommentLocationType int           `json:"commentLocationType"`
		ParentCommentId     int           `json:"parentCommentId"`
		Decoration          struct {
		} `json:"decoration"`
		RepliedMark interface{} `json:"repliedMark"`
		Liked       bool        `json:"liked"`
	} `json:"comment"`
}

// ReplyCommentData 回复评论返回数据
type ReplyCommentData struct {
	RawJson string
	Code    int `json:"code"`
	Comment struct {
		User struct {
			LocationInfo interface{} `json:"locationInfo"`
			LiveInfo     interface{} `json:"liveInfo"`
			Anonym       int         `json:"anonym"`
			AuthStatus   int         `json:"authStatus"`
			AvatarUrl    string      `json:"avatarUrl"`
			Nickname     string      `json:"nickname"`
			VipRights    struct {
				Associator   interface{} `json:"associator"`
				MusicPackage struct {
					VipCode int  `json:"vipCode"`
					Rights  bool `json:"rights"`
				} `json:"musicPackage"`
				RedVipAnnualCount int `json:"redVipAnnualCount"`
				RedVipLevel       int `json:"redVipLevel"`
			} `json:"vipRights"`
			UserId       int `json:"userId"`
			AvatarDetail struct {
				UserType        int    `json:"userType"`
				IdentityLevel   int    `json:"identityLevel"`
				IdentityIconUrl string `json:"identityIconUrl"`
			} `json:"avatarDetail"`
			UserType   int         `json:"userType"`
			ExpertTags interface{} `json:"expertTags"`
			VipType    int         `json:"vipType"`
			RemarkName interface{} `json:"remarkName"`
			Experts    interface{} `json:"experts"`
		} `json:"user"`
		BeReplied []struct {
			User struct {
				LocationInfo interface{} `json:"locationInfo"`
				LiveInfo     interface{} `json:"liveInfo"`
				Anonym       int         `json:"anonym"`
				AuthStatus   int         `json:"authStatus"`
				AvatarUrl    string      `json:"avatarUrl"`
				Nickname     string      `json:"nickname"`
				VipRights    interface{} `json:"vipRights"`
				UserId       int         `json:"userId"`
				AvatarDetail interface{} `json:"avatarDetail"`
				UserType     int         `json:"userType"`
				ExpertTags   interface{} `json:"expertTags"`
				VipType      int         `json:"vipType"`
				RemarkName   interface{} `json:"remarkName"`
				Experts      interface{} `json:"experts"`
			} `json:"user"`
			BeRepliedCommentId int         `json:"beRepliedCommentId"`
			Content            string      `json:"content"`
			Status             int         `json:"status"`
			ExpressionUrl      interface{} `json:"expressionUrl"`
		} `json:"beReplied"`
		PendantData         interface{} `json:"pendantData"`
		ShowFloorComment    interface{} `json:"showFloorComment"`
		Status              int         `json:"status"`
		CommentId           int         `json:"commentId"`
		Content             string      `json:"content"`
		Time                int         `json:"time"`
		LikedCount          int         `json:"likedCount"`
		ExpressionUrl       interface{} `json:"expressionUrl"`
		CommentLocationType int         `json:"commentLocationType"`
		ParentCommentId     int         `json:"parentCommentId"`
		Decoration          struct {
		} `json:"decoration"`
		RepliedMark interface{} `json:"repliedMark"`
		Liked       bool        `json:"liked"`
	} `json:"comment"`
}

// DelCommentData 删除评论返回数据
type DelCommentData struct {
	RawJson string
	Code    int `json:"code"`
}

// GetCommentData 获取评论返回数据
type GetCommentData struct {
	RawJson string
	Code    int `json:"code"`
	Data    struct {
		Comments []struct {
			User struct {
				AvatarDetail *struct {
					UserType        int    `json:"userType"`
					IdentityLevel   int    `json:"identityLevel"`
					IdentityIconUrl string `json:"identityIconUrl"`
				} `json:"avatarDetail"`
				CommonIdentity interface{} `json:"commonIdentity"`
				LocationInfo   interface{} `json:"locationInfo"`
				LiveInfo       interface{} `json:"liveInfo"`
				Followed       bool        `json:"followed"`
				VipRights      *struct {
					Associator struct {
						VipCode int  `json:"vipCode"`
						Rights  bool `json:"rights"`
					} `json:"associator"`
					MusicPackage      interface{} `json:"musicPackage"`
					RedVipAnnualCount int         `json:"redVipAnnualCount"`
					RedVipLevel       int         `json:"redVipLevel"`
				} `json:"vipRights"`
				RelationTag interface{} `json:"relationTag"`
				Anonym      int         `json:"anonym"`
				UserId      int64       `json:"userId"`
				UserType    int         `json:"userType"`
				Nickname    string      `json:"nickname"`
				AvatarUrl   string      `json:"avatarUrl"`
				AuthStatus  int         `json:"authStatus"`
				ExpertTags  interface{} `json:"expertTags"`
				Experts     interface{} `json:"experts"`
				VipType     int         `json:"vipType"`
				RemarkName  interface{} `json:"remarkName"`
				IsHug       bool        `json:"isHug"`
			} `json:"user"`
			BeReplied       interface{} `json:"beReplied"`
			CommentId       int64       `json:"commentId"`
			Content         string      `json:"content"`
			Status          int         `json:"status"`
			Time            int64       `json:"time"`
			LikedCount      int         `json:"likedCount"`
			Liked           bool        `json:"liked"`
			ExpressionUrl   interface{} `json:"expressionUrl"`
			ParentCommentId int         `json:"parentCommentId"`
			RepliedMark     bool        `json:"repliedMark"`
			PendantData     *struct {
				Id       int    `json:"id"`
				ImageUrl string `json:"imageUrl"`
			} `json:"pendantData"`
			ShowFloorComment struct {
				ReplyCount     int         `json:"replyCount"`
				Comments       interface{} `json:"comments"`
				ShowReplyCount bool        `json:"showReplyCount"`
				TopCommentIds  interface{} `json:"topCommentIds"`
				Target         interface{} `json:"target"`
			} `json:"showFloorComment"`
			Decoration struct {
			} `json:"decoration"`
			CommentLocationType int    `json:"commentLocationType"`
			Args                string `json:"args"`
			Tag                 struct {
				Datas []struct {
					Text          string `json:"text"`
					Type          int    `json:"type"`
					RecommendType string `json:"recommendType"`
					CommentType   int    `json:"commentType"`
				} `json:"datas"`
				RelatedCommentIds interface{} `json:"relatedCommentIds"`
			} `json:"tag"`
			Source *struct {
				Type    int           `json:"type"`
				Keys    []interface{} `json:"keys"`
				IconUrl string        `json:"iconUrl"`
				Text    string        `json:"text"`
				Orpheus string        `json:"orpheus"`
			} `json:"source"`
			ExtInfo struct {
			} `json:"extInfo"`
			CommentVideoVO struct {
				ShowCreationEntrance bool    `json:"showCreationEntrance"`
				AllowCreation        bool    `json:"allowCreation"`
				CreationOrpheusUrl   *string `json:"creationOrpheusUrl"`
				PlayOrpheusUrl       *string `json:"playOrpheusUrl"`
				VideoCount           int     `json:"videoCount"`
				ForbidCreationText   string  `json:"forbidCreationText"`
			} `json:"commentVideoVO"`
		} `json:"comments"`
		CurrentComment interface{} `json:"currentComment"`
		TotalCount     int         `json:"totalCount"`
		HasMore        bool        `json:"hasMore"`
		Cursor         string      `json:"cursor"`
		SortType       int         `json:"sortType"`
		SortTypeList   []struct {
			SortType     int    `json:"sortType"`
			SortTypeName string `json:"sortTypeName"`
			Target       string `json:"target"`
		} `json:"sortTypeList"`
	} `json:"data"`
}
