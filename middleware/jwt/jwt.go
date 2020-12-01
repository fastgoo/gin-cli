package jwt

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"github.com/caarlos0/env/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	jwt.StandardClaims
	UserId   uint32 `json:"user_id"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	Status   int32  `json:"status"`
}

type jwtConfig struct {
	Secret     string `env:"JWT_SECRET,required"`
	ExpireTime int    `env:"JWT_EXPIRE_TIME,required"`
}

var jwtConf = jwtConfig{}

func init() {
	if err := env.Parse(&jwtConf); err != nil {
		log.Fatal(err)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.Fail(c, http.StatusUnauthorized, e.ERR_AUTH_TOKEN_NOTFOUND)
			c.Abort()
			return
		}
		claims, err := parseToken(strings.Fields(token)[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERR_AUTH_CHECK_TOKEN_FAIL
			}
		}else{
			// 匹配登录的ip是否为当前ip,如果不是的话授权失败，需要重新去登录授权
			if claims.IP != c.ClientIP() {
				code = e.ERR_AUTH_CHECK_TOKEN_FAIL
			}
		}
		if code != e.SUCCESS {
			response.Fail(c, http.StatusUnauthorized, code, "")
			c.Abort()
			return
		}
		c.Set("authInfo", claims)
		c.Next()
	}
}

func parseToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtConf.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

func CreateToken(c Claims) (string, error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(jwtConf.ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(jwtConf.Secret))
}
