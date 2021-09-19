package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userid, username string) (*token, error) {
	accessKey := os.Getenv("ACCESSSECRETKEY")
	refreshKey := os.Getenv("REFRESHSECRETKEY")

	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(accessKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString(refreshKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tokenPair := token{AccessToken: accessTokenString, RefreshToken: refreshTokenString}
	fmt.Println("generated token pair")
	return &tokenPair, nil
}
