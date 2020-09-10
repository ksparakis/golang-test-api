package Controllers

import (
	"genity/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetUsers ... Get all users
func GetDataByTitle(c *gin.Context) {
	var data Models.Data
	title := c.Params.ByName("title")
	err := Models.GetDataByTitle(&data, title)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//CreateUser ... Create User
func CreateData(c *gin.Context) {
	var data Models.Data
	data.BeforeCreate(c)
	err := c.BindJSON(&data)
	if err != nil{
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	err = Models.CreateData(&data)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
