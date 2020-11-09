package controller

import (
	"encoding/json"
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
    var pokes[] model.Pokemon

    if err := c.ShouldBindJSON(&lim); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
		res :=  conn.Preload("Poketypes").Limit(lim.Num).Find(&pokes)
		resval := res.Value
		
		// marshaling the hits json; and unmarshalling;
		b, err := json.Marshal(resval)
		if err != nil {
			panic("Error in marshal updatepoke")
		}

		//configure the response structure 
		type Data struct {
			Pokemons []model.PokeTest `json:"pokemons"`
		}

		type JSONResp struct {
			Data Data `json:"data"`
		}

		var a[] model.PokeResp 
		err = json.Unmarshal(b, &a)

		var t []model.PokeTest

		for _, p  := range a {
				var tipe []string
				for _, each :=  range p.Poketypes {
					tipe = append(tipe, each.Element)
				}
				
				finalpokemon := model.PokeTest{
						ID: p.ID,
						Number: p.Number,
						Name: p.Name,
						Poketypes: tipe ,
				}

				t = append(t, finalpokemon)
		}

		jsonRespo := JSONResp {
			Data: Data{
			Pokemons: t,
			},
		}

    c.JSON(200, jsonRespo)
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

		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error"})
		}

		c.JSON(200, pokePost)

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

		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error"})
		}

		c.JSON(200, pokePost)

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