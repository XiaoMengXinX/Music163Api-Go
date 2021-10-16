package types

// AddCommentData 新增评论返回数据
type AddCommentData struct {
	RawJson          string
	Code             int `json:"code"`
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
