package user

import (
	"github.com/j2ephyr/realworld-starter-gin/src/common"
	"github.com/jinzhu/gorm"
	"log"
)

type Following struct {
	gorm.Model
	User            User
	UserId          uint
	FollowingUser   User
	FollowingUserId uint
}

func following(userId uint, followingUserId uint) bool {
	return !common.DB.Model(&Following{}).
		Where("user_id = ? and following_user_id = ?", userId, followingUserId).
		RecordNotFound()
}

func saveFollow(userId uint, followingUserId uint) bool {
	if following(userId, followingUserId) {
		return true
	}

	err := common.DB.Save(&Following{
		UserId:          userId,
		FollowingUserId: followingUserId,
	}).Error
	if err != nil {
		log.Println("saveFollow error: ", err)
		return false
	}
	return true
}

func deleteFollow(userId uint, followingUserId uint) bool {
	if !following(userId, followingUserId) {
		return true
	}

	err := common.DB.Delete(&Following{
		UserId:          userId,
		FollowingUserId: followingUserId,
	}).Error
	if err != nil {
		log.Println("deleteFollow error: ", err)
		return false
	}
	return true
}
