package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurrizkyimani/pahamifybackend/database"
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

func ReadPoke(c *gin.Context){
    conn := database.DBConn
    var lim model.Limit
    var poke model.Pokemon

    if err := c.ShouldBindJSON(&lim); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

        res :=  conn.Preload("Poketypes").Find(&poke)

    c.JSON(200, res)
    }

func CreatePoke( c *gin.Context){

		conn := 	database.DBConn
		var pokePost model.PokeTest

		if err := c.ShouldBindJSON(&pokePost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pokePost.Number == "" || pokePost.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		
		var pkts  []model.Poketype
		for _, p  := range pokePost.Poketypes {
			pkts =  append(pkts, model.Poketype{
					Element: p,
					UserID: pokePost.ID,
			})
		}

		finalpokemon := model.Pokemon{
			ID: pokePost.ID,
			Number: pokePost.Name,
			Name: pokePost.Name,
			Poketypes: pkts ,
			
		}
		res := conn.Create(&finalpokemon)
		c.JSON(200, res)

    }
    
func UpdatePoke( c *gin.Context){

		conn := 	database.DBConn
		var pokePost model.PokeTest
		var pokeQuery model.Pokemon

		if err := c.ShouldBindJSON(&pokePost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pokePost.Number == "" || pokePost.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		
		var pkts  []model.Poketype
		for _, p  := range pokePost.Poketypes {
			pkts =  append(pkts, model.Poketype{
					Element: p,
					UserID: pokePost.ID,
			})
		}

		finalpokemon := model.Pokemon{
			ID: pokePost.ID,
			Number: pokePost.Name,
			Name: pokePost.Name,
			Poketypes: pkts ,
			
		}
		
		conn.First(&pokeQuery, pokePost.ID)

		res := conn.Model(&pokeQuery).Updates(finalpokemon)

		res.Value

		c.JSON(200, res)

    }
    

func DeletePoke( c *gin.Context){
		
		// var PokeID model.PokeID

		conn := 	database.DBConn
		var pokePost model.PokeTest

		if err := c.ShouldBindJSON(&pokePost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pokePost.Number == "" || pokePost.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		
		var pkts  []model.Poketype
		for _, p  := range pokePost.Poketypes {
			pkts =  append(pkts, model.Poketype{
					Element: p,
					UserID: pokePost.ID,
			})
		}

		finalpokemon := model.Pokemon{
			ID: pokePost.ID,
			Number: pokePost.Name,
			Name: pokePost.Name,
			Poketypes: pkts ,
		}

		res := conn.Delete(&finalpokemon)

		// fmt.Println(PokeID.Num)

		if res.Error != nil {
				c.JSON(200, gin.H{
					"message": res.Error,
				})
		}

		c.JSON(200, gin.H{
			"message": "OK",
		})

	}