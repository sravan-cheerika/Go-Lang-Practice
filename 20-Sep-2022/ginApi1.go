package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
var db = make(map[string]string)
func setupRouter() *gin.Engine {
    // Disable Console Color
    // gin.DisableConsoleColor()
    r := gin.Default()
    // Ping test
    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    //http://localhost:8080/welcome?firstname=Jane&lastname=Doe
    r.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })
	r.GET("/json", func(c * gin.Context) {
		c.JSON(200, gin.H{"html" : "<b> Hello world ! </b>"})
	})

    return r
}
func main() {
    r := setupRouter()
    // Listen and Server in 0.0.0.0:8080
    r.Run(":8080")
}