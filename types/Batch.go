package types

// BatchUserSettingData 批处理 API 返回的 UserSetting API 数据
type BatchUserSettingData struct {
	ApiUserSetting struct {
		Setting struct {
			ProfileSetting                   int         `json:"profileSetting"`
			AgeSetting                       int         `json:"ageSetting"`
			AreaSetting                      int         `json:"areaSetting"`
			CollegeSetting                   int         `json:"collegeSetting"`
			VillageAgeSetting                int         `json:"villageAgeSetting"`
			FollowSingerSetting              int         `json:"followSingerSetting"`
			PersonalServiceSetting           int         `json:"personalServiceSetting"`
			NewSongDiskSetting               interface{} `json:"newSongDiskSetting"`
			TopLayerNotifySetting            int         `json:"topLayerNotifySetting"`
			ConcertSetting                   int         `json:"concertSetting"`
			PhoneFriendSetting               bool        `json:"phoneFriendSetting"`
			AllowKtvRoomNotify               bool        `json:"allowKtvRoomNotify"`
			UserId                           int         `json:"userId"`
			BroadcastSetting                 int         `json:"broadcastSetting"`
			ShareSetting                     int         `json:"shareSetting"`
			SocialSetting                    int         `json:"socialSetting"`
			PlayRecordSetting                int         `json:"playRecordSetting"`
			AllowSubscriptionNotify          bool        `json:"allowSubscriptionNotify"`
			AllowLikedNotify                 bool        `json:"allowLikedNotify"`
			AllowNewFollowerNotify           bool        `json:"allowNewFollowerNotify"`
			NeedRcmdEvent                    bool        `json:"needRcmdEvent"`
			AllowPlaylistShareNotify         bool        `json:"allowPlaylistShareNotify"`
			AllowDJProgramShareNotify        bool        `json:"allowDJProgramShareNotify"`
			AllowDJProgramSubscriptionNotify bool        `json:"allowDJProgramSubscriptionNotify"`
			AllowDJRadioSubscriptionNotify   bool        `json:"allowDJRadioSubscriptionNotify"`
			AllowPeopleCanSeeMyPlayRecord    bool        `json:"allowPeopleCanSeeMyPlayRecord"`
			AllowFollowedCanSeeMyPlayRecord  bool        `json:"allowFollowedCanSeeMyPlayRecord"`
			AllowOfflinePrivateMessageNotify bool        `json:"allowOfflinePrivateMessageNotify"`
			AllowOfflineCommentNotify        bool        `json:"allowOfflineCommentNotify"`
			AllowOfflineCommentReplyNotify   bool        `json:"allowOfflineCommentReplyNotify"`
			AllowOfflineNotify               bool        `json:"allowOfflineNotify"`
			AllowVideoSubscriptionNotify     bool        `json:"allowVideoSubscriptionNotify"`
			AllowOfflineForwardNotify        bool        `json:"allowOfflineForwardNotify"`
			SendMiuiMsg                      bool        `json:"sendMiuiMsg"`
			AllowImportDoubanPlaylist        bool        `json:"allowImportDoubanPlaylist"`
			AllowImportXiamiPlaylist         bool        `json:"allowImportXiamiPlaylist"`
			ImportedDoubanPlaylist           bool        `json:"importedDoubanPlaylist"`
			ImportedXiamiPlaylist            bool        `json:"importedXiamiPlaylist"`
			PeopleNearbyCanSeeMe             bool        `json:"peopleNearbyCanSeeMe"`
			FinishedFollowGuide              bool        `json:"finishedFollowGuide"`
			CommentSetting                   int         `json:"commentSetting"`
			MutualFollowSeeOnline            int         `json:"mutualFollowSeeOnline"`
			VoiceAutoPlay                    int         `json:"voiceAutoPlay"`
		} `json:"setting"`
		Code int `json:"code"`
	} `json:"/api/user/setting"`
}

// BatchUserInfoData 批处理 UserInfo 数据
type BatchUserInfoData struct {
	ApiV1UserInfo struct {
		Level     int `json:"level"`
		UserPoint struct {
			Balance      int   `json:"balance"`
			BlockBalance int   `json:"blockBalance"`
			Status       int   `json:"status"`
			UpdateTime   int64 `json:"updateTime"`
			UserId       int   `json:"userId"`
			Version      int   `json:"version"`
		} `json:"userPoint"`
		MobileSign       bool   `json:"mobileSign"`
		PcSign           bool   `json:"pcSign"`
		Viptype          int    `json:"viptype"`
		Expiretime       int64  `json:"expiretime"`
		BackupExpireTime int64  `json:"backupExpireTime"`
		Storeurl         string `json:"storeurl"`
		MallDesc         string `json:"mallDesc"`
		StoreTitle       string `json:"storeTitle"`
		Pubwords         string `json:"pubwords"`
		GameConfig       struct {
			ExpirationTime       int64  `json:"expirationTime"`
			GameCenterEntryTitle string `json:"gameCenterEntryTitle"`
			GameCenterPic        string `json:"gameCenterPic"`
			GameCenterReddot     int    `json:"gameCenterReddot"`
			GameCenterText       string `json:"gameCenterText"`
			GameCenterUrl        string `json:"gameCenterUrl"`
		} `json:"gameConfig"`
		RingConfig interface{} `json:"ringConfig"`
		FmConfig   struct {
			Image string `json:"image"`
			Text  string `json:"text"`
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"fmConfig"`
		TicketConfig interface{} `json:"ticketConfig"`
		Code         int         `json:"code"`
	} `json:"/api/v1/user/info"`
	ApiUserSetting struct {
		Setting struct {
			ProfileSetting                   int         `json:"profileSetting"`
			AgeSetting                       int         `json:"ageSetting"`
			AreaSetting                      int         `json:"areaSetting"`
			CollegeSetting                   int         `json:"collegeSetting"`
			VillageAgeSetting                int         `json:"villageAgeSetting"`
			FollowSingerSetting              int         `json:"followSingerSetting"`
			PersonalServiceSetting           int         `json:"personalServiceSetting"`
			NewSongDiskSetting               interface{} `json:"newSongDiskSetting"`
			TopLayerNotifySetting            int         `json:"topLayerNotifySetting"`
			ConcertSetting                   int         `json:"concertSetting"`
			PhoneFriendSetting               bool        `json:"phoneFriendSetting"`
			BroadcastSetting                 int         `json:"broadcastSetting"`
			ShareSetting                     int         `json:"shareSetting"`
			SocialSetting                    int         `json:"socialSetting"`
			PlayRecordSetting                int         `json:"playRecordSetting"`
			AllowSubscriptionNotify          bool        `json:"allowSubscriptionNotify"`
			AllowLikedNotify                 bool        `json:"allowLikedNotify"`
			AllowNewFollowerNotify           bool        `json:"allowNewFollowerNotify"`
			NeedRcmdEvent                    bool        `json:"needRcmdEvent"`
			AllowPlaylistShareNotify         bool        `json:"allowPlaylistShareNotify"`
			AllowDJProgramShareNotify        bool        `json:"allowDJProgramShareNotify"`
			AllowDJRadioSubscriptionNotify   bool        `json:"allowDJRadioSubscriptionNotify"`
			AllowPeopleCanSeeMyPlayRecord    bool        `json:"allowPeopleCanSeeMyPlayRecord"`
			AllowFollowedCanSeeMyPlayRecord  bool        `json:"allowFollowedCanSeeMyPlayRecord"`
			AllowOfflinePrivateMessageNotify bool        `json:"allowOfflinePrivateMessageNotify"`
			AllowOfflineCommentNotify        bool        `json:"allowOfflineCommentNotify"`
			AllowOfflineCommentReplyNotify   bool        `json:"allowOfflineCommentReplyNotify"`
			AllowOfflineNotify               bool        `json:"allowOfflineNotify"`
			AllowVideoSubscriptionNotify     bool        `json:"allowVideoSubscriptionNotify"`
			AllowOfflineForwardNotify        bool        `json:"allowOfflineForwardNotify"`
			SendMiuiMsg                      bool        `json:"sendMiuiMsg"`
			AllowImportDoubanPlaylist        bool        `json:"allowImportDoubanPlaylist"`
			AllowImportXiamiPlaylist         bool        `json:"allowImportXiamiPlaylist"`
			ImportedDoubanPlaylist           bool        `json:"importedDoubanPlaylist"`
			ImportedXiamiPlaylist            bool        `json:"importedXiamiPlaylist"`
			PeopleNearbyCanSeeMe             bool        `json:"peopleNearbyCanSeeMe"`
			FinishedFollowGuide              bool        `json:"finishedFollowGuide"`
			AllowDJProgramSubscriptionNotify bool        `json:"allowDJProgramSubscriptionNotify"`
			UserId                           int         `json:"userId"`
			AllowKtvRoomNotify               bool        `json:"allowKtvRoomNotify"`
			CommentSetting                   int         `json:"commentSetting"`
			MutualFollowSeeOnline            int         `json:"mutualFollowSeeOnline"`
			VoiceAutoPlay                    int         `json:"voiceAutoPlay"`
		} `json:"setting"`
		Code int `json:"code"`
	} `json:"/api/user/setting"`
	Code int `json:"code"`
}

// BatchSongURLData 批处理 SongURL 数据
type BatchSongURLData struct {
	ApiSongEnhancePlayerUrlV1 SongURLData `json:"/api/song/enhance/player/url/v1"`
}

// BatchSongDetailData 批处理 SongDetail 数据
type BatchSongDetailData struct {
	ApiV3SongDetail SongDetailData `json:"/api/v3/song/detail"`
}
