#Pahamify Intern Pre-test

# REST API

### Reload the `Seen` in Algolia as False 
* Endpoint: `/reload`
* HTTP Method: `GET`
* Request Header:
    * Accept: `NONE`
    * Content-type: `NONE`
    * Authorization: `NONE`
  
* Response Body:
```json
  [
    {
        "ProductName": "Sofa Java",
        "ObjectID": "8750990000"
    }
  ] 
  ```

### Search with certain keyword from Algolia
* Endpoint: `/search/sofa%20dudukan`
* HTTP Method: `GET`
* Request Header:
    * Accept: `NONE`
    * Content-type: `NONE`
    * Authorization: `NONE`
  
* Response Body:
```json
  {
    "ProductName": "Sofa Java",
    "ObjectID": "8750990000"
  }
  ```
