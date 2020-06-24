package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required,min=1,max=32"`
	Password string `json:"password" binding:"required,min=1,max=16"`
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Home",
	})
}

func PostHome(c *gin.Context) {
	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"UserInfo": reqBody,
		})
	}
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"Name": name,
		"Age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"Name": name,
		"Age":  age,
	})

}

func NoFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"error": "requested API not found",
	})

}
