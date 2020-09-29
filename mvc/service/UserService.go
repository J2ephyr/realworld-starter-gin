package service

import (
	"fmt"
	"github.com/j2ephyr/realworld-starter-gin/mvc/model"
	"github.com/j2ephyr/realworld-starter-gin/src/common"
)

var userDao model.UserDao = &model.UserDaoImpl{DB: common.DB}

func login(id uint) {
	user, err := userDao.FindById(id)
	fmt.Println("user:", user, "err:", err)
}
