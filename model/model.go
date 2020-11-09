package model

type Pokemon struct {
	ID string `json:"id" gorm:"primary_key"` 
	Number string `json:"number"`
	Name string `json:"name"`
	Poketypes []Poketype `json:"types" gorm:"foreignKey:UserID;references:ID"`

}

//Poketype is Struct of type of pokemon
type Poketype struct {
	IDtype uint `gorm:"primary_key;auto_increment;not_null"` 
	Element string `json:"element"`
	UserID string `json:"userID"`
}

type PokeTest struct {
	ID string `json:"id" gorm:"primary_key"` 
	Number string `json:"number"`
	Name string `json:"name"`
	Poketypes []string `json:"types"`
}

type PokeResp struct {
	ID string `json:"id" gorm:"primary_key"` 
	Number string `json:"number"`
	Name string `json:"name"`
	Poketypes []struct{
		Element string `json:"element"`
	} `json:"types"`

}

type Limit struct {	
	Num int `json:"limit"`    
}

type PokeID struct {	
	Num string `json:"ID"`    
}


