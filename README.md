# webapi-go
testing api using go, gin and gorm


# To init the configuration
go mod init webapi

go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Database configuration - PostreSQL
on file databases\database.go
## Default
host=localhost port=5432 user=postgres dbname=books sslmode=disable password=docker

# Starting the Server
go run main.go

# Methods available
## POST
URL: http://localhost:5000/api/v1/books

```
{
	"name" : "Livro 2",
	"description" : "Outro livro legal",
	"medium_price" : 99.99,
	"author" : "Bem Aquele",
	"img_url" : "www.site.com"
}
```

## GET all books
http://localhost:5000/api/v1/books

## GET book by ID
http://localhost:5000/api/v1/books/1

## PUT
http://localhost:5000/api/v1/books
```
{
	"id" : 2,
	"name" : "Livro 2 - NOVO",
	"description" : "Outro livro legal",
	"medium_price" : 99.99,
	"author" : "Bem Aquele",
	"img_url" : "www.site.com"
}

```
## DELETE
http://localhost:5000/api/v1/books/1