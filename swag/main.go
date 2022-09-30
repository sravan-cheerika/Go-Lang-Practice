package main

import (
    "net/http"
    docs "swaggerTest/docs"
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)
//@Author Subhash
// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
    g.JSON(http.StatusOK, "helloworld 1234")
}
func main() {
    r := gin.Default()
    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := r.Group("/api/v1")
    {
        eg := v1.Group("/example")
        {
            eg.GET("/hello123", Helloworld)
        }
    }
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    r.Run(":8080")
}