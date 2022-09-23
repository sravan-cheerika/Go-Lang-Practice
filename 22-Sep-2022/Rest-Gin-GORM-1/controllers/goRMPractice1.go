package main

import(
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)


func main(){

	router := gin.Default()
	router.GET("/",getAllUsers)
}

func getAllUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users)
}