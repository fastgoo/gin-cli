package jwt

import (
	"gin-cli/pkg/response"
	"net/http"
	"os"
	"strings"

	"gin-cli/pkg/e"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.ERR_AUTH_TOKEN_NOTFOUND
		} else {
			_, err := parseToken(strings.Fields(token)[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}
		if code != e.SUCCESS {
			response.Fail(c, http.StatusUnauthorized, code)
			c.Abort()
			return
		}
		c.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
