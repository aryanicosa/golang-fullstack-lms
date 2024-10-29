# Learning Management System with Golang

### Initialize Go module
go mod init golang-fullstack-lms

### Install required packages
go get -u github.com/gin-gonic/gin && /
go get -u gorm.io/gorm && /
go get -u gorm.io/driver/postgres && /
go get -u github.com/golang-jwt/jwt/v4 && /
go get -u golang.org/x/crypto/bcrypt && /
go get -u github.com/joho/godotenv && /
go get -u github.com/stripe/stripe-go/v72 && /
go get -u github.com/swaggo/gin-swagger && /
go get -u github.com/swaggo/files

### Install development tools
#### For hot reloading
go install github.com/cosmtrek/air@latest
#### For testing
go install github.com/golang/mock/mockgen@latest
#### for swagger
go install github.com/swaggo/swag/cmd/swag@latest