package Routers

import (
	Controller "go_practice/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/helloworld", Controller.Hello)
	r.POST("/register", func(c *gin.Context) {
		Controller.Register(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		Controller.Login(c, db)
	})
	return r
}
