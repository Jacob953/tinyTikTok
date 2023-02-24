package model

type Video struct {
	Id            int64  `gorm:"video_id"                          protobuf:"varint,1,req,name=id"                                json:"id,omitempty"`     // 视频唯一标识
	Author        User   `gorm:"foreignKey:AuthorID;references:Id" protobuf:"bytes,2,req,name=author"                             json:"author,omitempty"` // 视频作者信息
	AuthorID      int    `gorm:"column:author_id"                                                                                 json:"author_id"`
	PlayUrl       string `gorm:"play_url"                          protobuf:"bytes,3,req,name=play_url,json=playUrl"              json:"play_url,omitempty"`       // 视频播放地址
	CoverUrl      string `gorm:"cover_url"                         protobuf:"bytes,4,req,name=cover_url,json=coverUrl"            json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64  `gorm:"favorite_count"                    protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `gorm:"comment_count"                     protobuf:"varint,6,req,name=comment_count,json=commentCount"   json:"comment_count,omitempty"`  // 视频的评论总数
	IsFavorite    bool   `gorm:"is_favorite"                       protobuf:"varint,7,req,name=is_favorite,json=isFavorite"       json:"is_favorite,omitempty"`    // true-已点赞，false-未点赞
	Title         string `gorm:"title"                             protobuf:"bytes,8,req,name=title"                              json:"title,omitempty"`          // 视频标题
	CreateAt      int64  `gorm:"create_at"                                                                                        json:"timestamp"`
}

func (u *Video) TableName() string {
	return "video"
}
