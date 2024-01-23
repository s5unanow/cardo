package main

import (
    "github.com/gin-gonic/gin"
    "math/rand"
    "net/http"
    "path/filepath"
    "time"
)

// UserCredentials for simple auth testing
type UserCredentials struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}

func main() {
    router := gin.Default()

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Login endpoint
    router.POST("/login", func(c *gin.Context) {
        var creds UserCredentials
        if err := c.ShouldBindJSON(&creds); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // For testing purposes, you can have predefined credentials or implement your checking logic
        if creds.Login == "test" && creds.Password == "password" {
            token := generateToken()
            c.JSON(http.StatusOK, gin.H{"token": token})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        }
    })

    // Middleware to check for the presence of an authorization token
    authMiddleware := func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        
        // Here, you'd implement your logic to validate the token
        // For simplicity, we assume any non-empty token is valid
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
            return
        }
        c.Next()
    }

    // Apply the authMiddleware to all routes except /login
    router.Use(authMiddleware)

    // Random error simulation for /plants route
    router.GET("/plants", func(c *gin.Context) {
        if rand.Intn(100) < 5 {
            c.AbortWithStatus(http.StatusInternalServerError)
        } else {
            resourcesDir := "resources"
            filePath := filepath.Join(resourcesDir, "plants.json")
            c.File(filePath)
        }
    })

    router.Run(":8080")
}

// generateToken creates a dummy token. In a real application, use a more robust method.
func generateToken() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, 20)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

