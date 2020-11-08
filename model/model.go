package model

type Pokemon struct {
 ID uint `gorm:"primary_key"` 
	Name string `json:"name"`
 Number string `json:"number"`
 Types []Poketype `json:"types" gorm:"embedded"`

}


type Limit struct {	
	Num int `json:"limit"`    
}

type PokeID struct {	
	Num int `json:"id"`    
}

//Poketype is Struct of type of pokemon
type Poketype struct {
	Element string `json:"element"`
}


// type ProductSeenUpdate struct  {
// 	ObjectID string `json:"objectID"`
// 	Seen bool`json:"Seen"`
// }

// //Response is a xxxx
// type Respoonse struct {
// 	Params string `json:"Params"`
// 	Hits []Hit `json:"Hits"`
// }


// type Hit struct {
// 	ProductName string `json:"ProductName"`
// 	ObjectID string `json:"ObjectID"`
// }