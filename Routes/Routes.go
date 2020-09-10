package Routes

import (
	"genity/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()


	router.POST("/post-data/:title", Controllers.CreateData)
	router.GET("/get-data/:title", Controllers.GetDataByTitle)

	return router
}