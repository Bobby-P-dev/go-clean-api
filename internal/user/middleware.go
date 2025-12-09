package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// 1. Tangkap error (ganti _ dengan err)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// 2. PASTIKAN INI SAMA DENGAN DI LOGIN SERVICE
			// Misal kita sepakat pakai "JWT_SECRET"
			secret := os.Getenv("JWT_KEY")
			if secret == "" {
				// Print error ke terminal server biar sadar
				fmt.Println("CRITICAL: JWT_SECRET is empty in middleware")
				return nil, errors.New("secret key missing")
			}

			return []byte(secret), nil
		})

		// 3. Cek Error Parse secara detail
		if err != nil {
			// Ini akan memberitahumu KENAPA gagal di response JSON
			fmt.Println("Token validation error:", err) // Log ke terminal
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 4. SAMAKAN KEY CONTEXT DENGAN HANDLER
			c.Set("user_id", claims["user_id"])
			c.Set("username", claims["username"]) // Ubah "user_name" jadi "username" biar konsisten
			c.Set("email", claims["email"])
			c.Set("exp", claims["exp"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
	}
}
