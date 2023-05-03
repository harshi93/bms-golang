# BMS With Golang
book management system using golang

# Compilation
cd cmd/main
Run `go build`

# Starting server
Run `./main`

# Accessing server
Open in Browser `localhost:9080`

# Routes
#### Add book
Use POST method with Postman
`curl localhost:9080/addbook` 

#### Get all books 
`curl localhost:9080/getbooks`

#### Get single book by id 
`curl localhost:9080/listbook/{bookId}`

#### Update book
Use POST method with Postman
`curl localhost:9080/modbook/{bookId}`

#### Del book by id
Use POST method with Postman
`curl localhost:9080/delbook/{bookId}`

# Sample payload
```json
{
    Name: "",
    Author: "",
    Publication: ""
}
```