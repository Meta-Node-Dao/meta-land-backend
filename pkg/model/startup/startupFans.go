package startup

type StartupFansInfo struct {
	ComerId      uint64 `gorm:"column:comer_id" json:"comerId"`
	ComerAvatar  string `gorm:"column:avatar" json:"comerAvatar"`
	ComerName    string `gorm:"column:name" json:"comerName"`
	FollowedByMe *bool  `json:"followedByMe"`
}
type StartupFans []*StartupFansInfo
