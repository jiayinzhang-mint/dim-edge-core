package auth

import (
	"dim-edge/core/utils"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

// UserPool is a user info container
type UserPool struct{}

var userInstance *UserPool
var userOnce sync.Once

var claims jwt.MapClaims
var ok bool

// GetUserInstance of current user
func GetUserInstance() *UserPool {
	userOnce.Do(func() {
		userInstance = &UserPool{}
	})
	return userInstance
}

// GetUser claims
func (m *UserPool) GetUser() jwt.MapClaims {
	return claims
}

func decodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Error("Unexpected signing method")
			return nil, fmt.Errorf("")
		}

		hmacSampleSecret := []byte(viper.GetString("tokenKey"))

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	return token, err
}

// CheckAuth check user access token
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get authorization header
		authString := r.Header.Get("Authorization")
		if len(authString) < 10 {
			utils.RespondWithError(w, r, 402, "invalid token")
			return
		}

		// Check token head
		if strings.Fields(authString)[0] != "DIMEDGE" {
			utils.RespondWithError(w, r, 402, "invalid token")
			return
		}

		// Cut header to get jwt string
		tokenString := strings.Fields(authString)[1]

		token, err := decodeToken(tokenString)

		if err != nil {
			utils.RespondWithError(w, r, 402, "invalid token")
			return
		}

		if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check expire time
			now := time.Now()
			exp := time.Unix(int64(claims["exp"].(float64)), 0)
			if now.After(exp) {
				utils.RespondWithError(w, r, 402, "token expired")
				return
			}

		} else {
			utils.RespondWithError(w, r, 402, "invalid token")
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
