package main

import (
	"strconv"
	"log"
	"net/http"
	"github.com/agentmilindu/soap2rest/gen"
	"github.com/hooklift/gowsdl/soap"
	"github.com/gin-gonic/gin"
)

var client = soap.NewClient("{{ config['repository']['url'] }}")
var service = gen.NewCalculatorSoap(client)


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

{% for endpoint in config['endpoints'] %}
	// Get user value
	r.GET("{{ endpoint['endpoint'] }}", func(c *gin.Context) {

		// Getting params and doing transisions on them should be happening from config.yaml
		a, err := strconv.ParseInt(c.Params.ByName("a"), 10, 32)
		b, err := strconv.ParseInt(c.Params.ByName("b"), 10, 32)

                if err != nil {
                        c.JSON(http.StatusOK, gin.H{"error": err})
                        log.Fatalf("error: %v", err)
                }

		reply, err := service.{{ endpoint['mapping'] }}(&gen.{{ endpoint['mapping'] }}{IntA: int32(a), IntB: int32(b)})

	        if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err})
        	        log.Fatalf("error: %v", err)
        	}

        	response := &reply

                if err != nil {
                        c.JSON(http.StatusOK, gin.H{"error": err})
                        log.Fatalf("error: %v", err)
                }

		c.JSON(http.StatusOK, response)

	})
{% endfor %}
	return r
}


func main() {

        r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
