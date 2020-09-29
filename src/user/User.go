package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/j2ephyr/realworld-starter-gin/src/common"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Bio      string
	Image    string
}

func Route(authorization *gin.RouterGroup, api *gin.RouterGroup) {
	//Get Current User
	authorization.GET("/user", getCurrentUser)
	authorization.PUT("/user", updateUser)
	api.POST("/users", registration)
	api.POST("/users/login", login)
}

type registrationForm struct {
	User struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

func registration(c *gin.Context) {
	form := registrationForm{}
	_ = c.ShouldBind(&form)
	user, notFound := getUserByEmail(&form.User.Email)
	if !notFound {
		common.Error(c, "Email has been registered!")
		return
	}
	user = User{
		Username: form.User.Username,
		Email:    form.User.Email,
		Password: form.User.Password,
	}
	common.GetDB().Save(&user)
	c.JSON(200, doLogin(&user))
}

type loginForm struct {
	User struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

func login(c *gin.Context) {
	form := loginForm{}
	_ = c.ShouldBind(&form)
	fmt.Printf("user: %s", form.User)
	user, notFound := getUserByEmail(&form.User.Email)
	if notFound {
		common.Error(c, "wrong password!")
		return
	}
	if form.User.Password != user.Password {
		common.Error(c, "wrong password!")
		return
	}
	c.JSON(200, doLogin(&user))
}

func doLogin(user *User) map[string]interface{} {
	token, err := GenerateToken(user)
	if err != nil {
		fmt.Println("GenerateToken Error:", err)
	}
	return map[string]interface{}{
		"user": map[string]interface{}{
			"email":    user.Email,
			"token":    token,
			"username": user.Username,
			"bio":      user.Bio,
			"image":    user.Bio,
		},
	}
}

func getCurrentUser(c *gin.Context) {
	userId, _ := c.Get("UserId")
	var user User
	notFound := common.DB.Where("id = ?", userId).First(&user).RecordNotFound()
	if notFound {
		c.AbortWithStatus(401)
		return
	}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Println("GenerateToken Error:", err)
	}
	c.JSON(200, map[string]interface{}{
		"user": map[string]interface{}{
			"email":    user.Email,
			"token":    token,
			"username": user.Username,
			"bio":      user.Bio,
			"image":    user.Bio,
		},
	})
}

type UpdateUserForm struct {
	User struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Bio      string `json:"bio" binding:"required"`
		Image    string `json:"image" binding:"required"`
	} `json:"user" binding:"required"`
}

func updateUser(c *gin.Context) {
	userId, _ := c.Get("UserId")
	var user User
	notFound := common.DB.Where("id = ?", userId).First(&user).RecordNotFound()
	if notFound {
		c.AbortWithStatus(401)
		return
	}
	var form UpdateUserForm
	_ = c.ShouldBind(&form)
	user.Username = form.User.Username
	user.Email = form.User.Email
	user.Password = form.User.Password
	user.Bio = form.User.Bio
	user.Image = form.User.Image
	user.ID = userId.(uint)
	err := common.DB.Model(&user).Updates(&user).Error
	if err != nil {
		common.Error(c, "DB Error:"+err.Error())
		return
	}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Println("GenerateToken Error:", err)
	}
	common.DB.Where("id = ?", userId).First(&user)
	c.JSON(200, map[string]interface{}{
		"user": map[string]interface{}{
			"email":    user.Email,
			"token":    token,
			"username": user.Username,
			"bio":      user.Bio,
			"image":    user.Image,
		},
	})
}

func getUserByEmail(email *string) (User, bool) {
	var user User
	fmt.Println("user:", &user)
	notFound := common.GetDB().Where(&User{Email: *email}).Find(&user).RecordNotFound()
	return user, notFound
}

func GenerateToken(user *User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := UserClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "realworld",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(common.JwtKey))
	return token, err
}

type UserClaims struct {
	UserID uint
	jwt.StandardClaims
}
