package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"learn-middleware-example/app/services"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get header
		tknheader := c.GetHeader("Authorization")

		//cek tokenheader, jika kosong maka return err / nil data
		if tknheader == "" {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Unathorized"})
			return
		}

		//replace
		tknheader = strings.Replace(tknheader, "Bearer ", "", 1)

		//parse jwt
		token, err := services.ParseJWT(tknheader)
		//if smth got err and token isn't valid
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Invalid Token"})
			return
		}

		//set the username in the context
		//Jika jwt valid dan klaim berhasil diakses, middleware akan melanjutkan pemrosesan permintaan selanjutnya dengan menggunakan c.Next()
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			return
		}
		c.Next()
	}
}
