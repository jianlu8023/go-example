package entity

import (
	"encoding/json"
)

type User struct {
	// Uid 如果需要自增，不指定type 可以实现 不清楚为什么增加type后，自增不生效
	Uid      int64  `json:"uid" gorm:"column:uid;primary_key;auto_increment;comment:'自增id'"`
	Username string `json:"username" gorm:"column:username;type:varchar(50);not null;comment:'用户名'"`
}

func (User) TableName() string {
	return "t_user"
}

func (u User) String() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}
