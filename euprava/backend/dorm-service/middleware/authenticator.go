package middleware

import (
	"log"
	"net/http"
	"strings"

	helper "dorm-service/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		clientToken = strings.Replace(clientToken, "Bearer ", "", 1)
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)

		c.Next()

	}
}
func AuthorizeRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_type, _ := c.Get("user_type")
		l := log.New(gin.DefaultWriter, "auth: ", log.LstdFlags)
		userRole := user_type.(string)
		l.Printf("User role: %s", userRole)
		authorized := false
		for _, role := range roles {
			l.Printf("Comparing: %s and %s", userRole, role)
			if userRole == role {
				l.Printf("%s and %s are the same", userRole, role)
				authorized = true
				break
			}
		}

		if !authorized {
			l.Printf("Access forbidden to role %s", userRole)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
