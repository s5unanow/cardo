package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Create a Gin router with default middleware: logger and recovery (crash-free) middleware
    router := gin.Default()

    // Define a route
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
        })
    })

    // Start serving the application
    router.Run(":8080")
}

