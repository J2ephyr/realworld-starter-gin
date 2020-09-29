package user

import (
	"github.com/jinzhu/gorm"
)

type Following struct {
	gorm.Model
	User            User
	UserId          uint
	FollowingUser   User
	FollowingUserId uint
}
