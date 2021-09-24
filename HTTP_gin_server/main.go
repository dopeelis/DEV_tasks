package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/about", getAbout)

			user := v1.Group("/user")
			{
				user.GET("/:id", getUserByID)
			}
		}
	}

	router.Run(":8080")
}

func getAbout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"page": "this is 'about' page",
	})
}

func getUserByID(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "can't convert user ID",
		})
	}
	if ID < 0 {
		ctx.JSON(406, gin.H{
			"error": "user ID has invalid value",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"id": ID,
	})
}
