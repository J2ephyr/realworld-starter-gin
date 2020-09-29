package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/j2ephyr/realworld-starter-gin/src/article"
	"github.com/j2ephyr/realworld-starter-gin/src/common"
	"github.com/j2ephyr/realworld-starter-gin/src/user"
	"github.com/jinzhu/gorm"
	"strings"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	initTable(db)
	route := gin.Default()
	authorization := route.Group("/api")
	authorization.Use(authHandler())
	api := route.Group("/api")
	user.Route(authorization, api)
	route.Run(":1018")
}

func authHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) == 0 {
			c.AbortWithStatus(401)
			return
		}
		tokenString := strings.Split(auth, " ")[1]
		if len(tokenString) == 0 {
			c.AbortWithStatus(401)
			return
		}
		claims := user.UserClaims{}
		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return []byte(common.JwtKey), nil
		})
		fmt.Println("token: ", token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Set("UserId", claims.UserID)
		c.Next()
	}
}

func initTable(db *gorm.DB) {
	db.AutoMigrate(user.User{}, user.Following{}, article.Article{}, article.Comment{}, article.Favorite{}, article.Tag{})
}
