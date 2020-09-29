package article

import (
	"github.com/j2ephyr/realworld-starter-gin/src/user"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	User   user.User
	UserId uint
	Body   string
}
