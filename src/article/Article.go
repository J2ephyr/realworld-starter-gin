package article

import (
	"github.com/j2ephyr/realworld-starter-gin/src/user"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	User        user.User
	UserId      uint
	TagList     []Tag
	Slug        string
	Title       string
	Description string
	Body        string
}
