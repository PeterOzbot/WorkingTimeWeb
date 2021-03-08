package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"workingtimeweb/server/generator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port = flag.String("port", "8080", "Port on which the server runs.")

func main() {

	// read flags
	flag.Parse()

	// initialize routing
	router := gin.Default()

	// enable CORS
	initCORS(router)

	// initialize http
	initHTTP(router)

	// run
	router.Run(fmt.Sprintf(":%s", *port))
}

func initCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func initHTTP(router *gin.Engine) {

	// set the routs
	router.POST("/generate", generate)
}

func generate(c *gin.Context) {
	// deserialize generator request
	var request *generator.Request = &generator.Request{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate hours
	generatedWorkingDays := generator.Generate(request)

	// return result
	c.JSON(200, generatedWorkingDays)
}
