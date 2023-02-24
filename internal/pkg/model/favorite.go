package model

type Favorite struct {
	VideoId  int64  `gorm:"video_id"  json:"video_id"`  // 视频ID
	UserId   *User  `gorm:"user_id"   json:"user_id"`   // 点赞用户ID
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间戳
}

func (u *Favorite) TableName() string {
	return "favorite"
}
