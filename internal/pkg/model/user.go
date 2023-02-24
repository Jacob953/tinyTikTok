package model

import (
	"fmt"

	"github.com/marmotedu/component-base/pkg/auth"
)

type User struct {
	Id              int64  `gorm:"column:user_id"          protobuf:"varint,1,req,name=id"                                   json:"user_id,omitempty"`          // 用户id
	Name            string `gorm:"column:name"             protobuf:"bytes,2,req,name=name"                                  json:"name,omitempty"`             // 用户名称
	FollowCount     int64  `gorm:"column:follow_count"     protobuf:"varint,3,opt,name=follow_count,json=followCount"        json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count"   protobuf:"varint,4,opt,name=follower_count,json=followerCount"    json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `gorm:"column:is_follow"        protobuf:"varint,5,req,name=is_follow,json=isFollow"              json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar          string `gorm:"column:avatar"           protobuf:"bytes,6,opt,name=avatar"                                json:"avatar,omitempty"`           //用户头像
	BackgroundImage string `gorm:"column:background_image" protobuf:"bytes,7,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `gorm:"column:signature"        protobuf:"bytes,8,opt,name=signature"                             json:"signature,omitempty"`        //个人简介
	TotalFavorited  int64  `gorm:"column:total_favorited"  protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited"  json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `gorm:"column:work_count"       protobuf:"varint,10,opt,name=work_count,json=workCount"           json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `gorm:"column:favorite_count"   protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount"   json:"favorite_count,omitempty"`   //点赞数量
	Password        string `gorm:"column:password"                                                                           json:"password,omitempty"`         // 用户密码
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Compare(pwd string) error {
	if err := auth.Compare(u.Password, pwd); err != nil {
		return fmt.Errorf("failed to compile password: %w", err)
	}

	return nil
}
