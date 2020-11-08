package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"github.com/nurrizkyimani/pahamifybackend/model"
)

//CORSMiddleware is middle ware for cors
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main(){
	db.InitDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())


	r.POST("/list", func(c *gin.Context){

		conn := db.DBConn 
		var lim model.Limit
		var poke model.Pokemon

		if err := c.ShouldBindJSON(&lim); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := conn.Limit(lim.Num).Find(&poke)

		c.JSON(200, res)
	})


	//post route
	r.GET("/post", func( c *gin.Context){

		conn := db.DBConn 
		var pokePost model.Pokemon

		if err := c.ShouldBindJSON(&pokePost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pokePost.Number == "" || pokePost.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 

		res := conn.Create(&pokePost)
		c.JSON(200, res)

	})

	//update route 
	r.GET("/update", func( c *gin.Context){

		conn := db.DBConn 
		var pokePost model.Pokemon

		if err := c.ShouldBindJSON(&pokePost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pokePost.Number == "" || pokePost.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 

		res := conn.Model(&pokePost).Updates(
			model.Pokemon{
				Name : "Hello",
				Number : "002",
				
			},
		)

		c.JSON(200, res )

	})

	r.DELETE("/delete", func( c *gin.Context){

		conn := db.DBConn 
		var pokePost model.Pokemon

		var PokeID model.PokeID

		if err := c.ShouldBindJSON(&PokeID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		conn.Delete(&pokePost, PokeID.Num)

		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})


	r.Run("127.0.0.1:8080")

}
