package main

import (
	"fmt"
	"genity/Config"
	"genity/Models"
	"genity/Routes"
)
var err error

func main() {

	err := Config.BuildDB()

	if err != nil {
		fmt.Println("Status:", err)
	}

	err = Config.DB.AutoMigrate(&Models.Data{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	r := Routes.SetupRouter()
	//running
	r.Run()


}