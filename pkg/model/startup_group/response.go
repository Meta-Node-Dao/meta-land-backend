package startup_group

import "time"

type GroupMember struct {
	StartupId   uint64     `gorm:"column:startup_id" json:"startupId"`
	ComerId     uint64     `gorm:"column:comer_id" json:"comerId"`
	ComerName   string     `gorm:"column:comer_name" json:"comerName"`
	ComerAvatar string     `gorm:"column:avatar" json:"comerAvatar"`
	Position    string     `gorm:"position" json:"position"`
	GroupId     *uint64    `gorm:"column:group_id" json:"groupId"`
	GroupName   *string    `gorm:"column:group_name" json:"groupName"`
	JoinedTime  *time.Time `gorm:"column:created_at" json:"joinedTime"`
}
