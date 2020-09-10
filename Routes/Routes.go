package Routes

import (
	"genity/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("")
	{
		v1.POST("/post-data/:data", Controllers.CreateData)
		v1.GET("/get-data/:title", Controllers.GetDataByTitle)

	}

	return router
}