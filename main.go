package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nurrizkyimani/pahamifybackend/controller"
	"github.com/nurrizkyimani/pahamifybackend/database"
)

func main(){
	database.InitDatabase()

	defer database.DBConn.Close()

	r := gin.Default()
	r.Use(controller.CORSMiddleware())

	//READ route
	r.POST("/list", controller.ReadPoke)

	//CREATE route
	r.POST("/create", controller.CreatePoke )

	//UPDATE route 
	r.POST("/update", controller.UpdatePoke)

	//DELETE route
	r.DELETE("/delete", controller.DeletePoke)


	r.Run("127.0.0.1:8080")

}
