package model

import (
	"fmt"
	"time"
)

type Comment struct {
	Id         int64  `gorm:"column:comment_id"  protobuf:"varint,1,req,name=id"                         json:"id,omitempty"`          // 视频评论id
	User       *User  `gorm:"column:user"        protobuf:"bytes,2,req,name=user"                        json:"user,omitempty"`        // 评论用户信息
	Content    string `gorm:"column:content"     protobuf:"bytes,3,req,name=content"                     json:"content,omitempty"`     // 评论内容
	CreateDate string `gorm:"column:create_date" protobuf:"bytes,4,req,name=create_date,json=createDate" json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
	VideoId    int64  `gorm:"column:video_id"                                                            json:"video_id,omitempty"`
}

func (u *Comment) TableName() string {
	return "comment"
}

func (u *Comment) FormatDate() {
	date := time.Now()
	u.CreateDate = fmt.Sprintf("%02d-%02d", date.Month(), date.Day())
}
