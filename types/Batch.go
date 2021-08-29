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
