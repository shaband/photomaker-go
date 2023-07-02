package middleware

import (
	"log"
	"net/http"
	"github.com/ansel1/merry"
	"github.com/gin-gonic/gin"
)
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // Create a Merry error from the recovered panic
                merryErr := merry.Wrap(err.(error))
                // Log the error, you can replace this with your preferred logging mechanism
                log.Println(merryErr)


                
                // Set the response status code
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Internal Server Error",
                })

                // Abort the current request
                c.Abort()
            }
        }()

        // Continue processing the request
        c.Next()
    }
}