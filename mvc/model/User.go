package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Bio      string
	Image    string
}

type UserDao interface {
	FindById(id uint) (User, error)
}

type UserDaoImpl struct {
	DB *gorm.DB
}

func (dao *UserDaoImpl) FindById(id uint) (user User, error error) {
	error = dao.DB.Where("id = ?", id).First(&user).Error
	return
}
