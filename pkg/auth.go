package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(userid int, username string) (*token, error) {
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
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
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

func Signup(w http.ResponseWriter, r *http.Request) {
	// read form data and check if form is valid
	var formData user
	json.NewDecoder(r.Body).Decode(&formData)
	fmt.Println(formData)

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert user data to database
	userid := 0
	username := ""

	err = Db.QueryRow(
		"INSERT INTO users (username, fullname, email, password) VALUES (?, ?, ?, ?) RETURNING (userid, username);",
		formData.Username, formData.Fullname, formData.Email, string(passwordHash),
	).Scan(&userid, &username)
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
	tokenPair, err := GenerateToken(userid, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	cookie := http.Cookie{
		HttpOnly: true,
		Name:     "refreshToken",
		Value:    tokenPair.RefreshToken,
		Domain:   "jpoly1219devbox.xyz",
		Path:     "/auth/",
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(tokenPair.AccessToken)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// read form data and check if form is valid
	var formData user
	json.NewDecoder(r.Body).Decode(&formData)
	fmt.Println(formData)

	// compare form data to database
	dbPasswordHash := ""
	err := Db.QueryRow(
		"SELECT password FROM users WHERE username = ?;",
		formData.Username,
	).Scan(&dbPasswordHash)
	if err != nil {
		fmt.Println(err)
		return
	}

	pwMatchErr := bcrypt.CompareHashAndPassword([]byte(formData.Password), []byte(dbPasswordHash))
	if pwMatchErr != nil {
		fmt.Println(pwMatchErr)
		return
	} else {
		// generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
		tokenPair, err := GenerateToken(formData.Userid, formData.Username)
		if err != nil {
			fmt.Println(err)
			return
		}

		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "refreshToken",
			Value:    tokenPair.RefreshToken,
			Domain:   "jpoly1219devbox.xyz",
			Path:     "/auth/",
		}
		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(tokenPair.AccessToken)
	}
}

func SilentRefresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refreshToken")
	if err != nil {
		fmt.Println(err)
		return
	}
	refreshToken := c.Value

	// check if refresh token is valid
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESSSECRETKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if so then generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
		userid := claims["userid"]
		username := claims["username"]
		useridStr := userid.(int)
		usernameStr := username.(string)
		tokenPair, err := GenerateToken(useridStr, usernameStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "refreshToken",
			Value:    tokenPair.RefreshToken,
			Domain:   "jpoly1219devbox.xyz",
			Path:     "/auth/",
		}
		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(tokenPair.AccessToken)
	} else {
		fmt.Println(err)
	}
}
