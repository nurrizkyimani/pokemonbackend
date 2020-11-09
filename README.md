

# Pahamify Intern Pre-test

##  Technology Stack 
 * Golang,
 * GORM, 
 * Gin Gonic,
 * Heroku PSQL, 
 * Docker,
 * Godotenv
 * Git/Github

## Instruction

The database is based on Heroku PosgreSQL, therefore the env is open with the url needed 

Using docker, build the image 
```
docker build . -t pahamifybackend
```

Then after the image is builded, run 
```
docker run -i -t -p 8080:8080 pahamifybackend

```

# REST API

## Postman 

https://documenter.getpostman.com/view/11806282/TVejgVBi

### Reload the `Seen` in Algolia as False 
* Endpoint: `/list`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `NONE`

* Requst Body:
```json
  [
    {
    "limit": 2
    }
  ] 
  ```
  
* Response Body:
```json
  {
    "data": {
        "pokemons": [
            {
                "id": "57072197-e96e-4bf7-83e9-23c740183220",
                "number": "Bulbasaur",
                "name": "Bulbasaur",
                "types": [
                    "Grass",
                    "Poison",
                    "Grass",
                    "Poison"
                ]
            },
            {
                "id": "57072197-e96e-4bf7-83e9-23c740183221",
                "number": "Bulbasaur",
                "name": "Bulbasaur",
                "types": [
                    "Grass",
                    "Poison"
                ]
            }
        ]
    }
}
  ```

### Search with certain keyword from Algolia
* Endpoint: `/create`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `NONE`
  
* Request Body:
```json
  {
  "id": "57072197-e96e-4bf7-83e9-23c740183226",
  "number": "001",
  "name": "Bulbasaur",
  "types": [
    "Grass",
    "Poison"
  ]
}

  ```

* Response Body:
```json
  {
  "id": "57072197-e96e-4bf7-83e9-23c740183226",
  "number": "001",
  "name": "Bulbasaur",
  "types": [
    "Grass",
    "Poison"
  ]
}

  ```


### Update the Pokemon Data
* Endpoint: `/update`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `NONE`
  
* Request Body:
```json
  {
  "id": "57072197-e96e-4bf7-83e9-23c740183226",
  "number": "001",
  "name": "Bulbasaur",
  "types": [
    "Grass",
    "Poison"
  ]
}

  ```

* Response Body:
```json
  {
  "id": "57072197-e96e-4bf7-83e9-23c740183226",
  "number": "001",
  "name": "Bulbasaur",
  "types": [
    "Grass",
    "Poison"
  ]
}
```

### Update the Pokemon Data
* Endpoint: `/delete`
* HTTP Method: `DELETE`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `NONE`
  
* Request Body:
```json
  {
  "id": "57072197-e96e-4bf7-83e9-23c740183226",
  "number": "001",
  "name": "Bulbasaur",
  "types": [
    "Grass",
    "Poison"
  ]
}

  ```

* Response Body:
```json
  {
  "message": "OK",
  }

  ```