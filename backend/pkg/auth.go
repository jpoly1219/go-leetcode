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

func HandleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func GenerateToken(userid int, username string) (*token, error) {
	accessKey := os.Getenv("ACCESSSECRETKEY")
	refreshKey := os.Getenv("REFRESHSECRETKEY")
	accessExp := time.Now().Add(time.Minute * 15).Unix()
	refreshExp := time.Now().Add(time.Hour * 24).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"exp":      accessExp,
	})
	accessTokenString, err := accessToken.SignedString([]byte(accessKey))
	if err != nil {
		fmt.Println("failed to sign access token: ", err)
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"exp":      refreshExp,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(refreshKey))
	if err != nil {
		fmt.Println("failed to sign refresh token: ", err)
		return nil, err
	}

	tokenPair := token{AccessToken: accessTokenString, RefreshToken: refreshTokenString}
	return &tokenPair, nil
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// handle preflight OPTIONS request
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// read form data and check if form is valid
	var formData user
	json.NewDecoder(r.Body).Decode(&formData)
	fmt.Println(formData)

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	if err != nil {
		fmt.Println("failed to generate password hash", err)
		return
	}

	// insert user data to database
	userid := 0
	username := ""

	err = Db.QueryRow(
		"INSERT INTO users (username, fullname, email, password) VALUES ($1, $2, $3, $4) RETURNING id, username;",
		formData.Username, formData.Fullname, formData.Email, string(passwordHash),
	).Scan(&userid, &username)
	// check if user already exists
	if err != nil {
		fmt.Println("signup() query failed:", err)
		return
	}

	// generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
	tokenPair, err := GenerateToken(userid, username)
	if err != nil {
		fmt.Println("failed to generate token pair:", err)
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
	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(tokenPair.AccessToken)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// handle preflight OPTIONS request
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// read form data and check if form is valid
	var formData user
	json.NewDecoder(r.Body).Decode(&formData)
	// fmt.Println(formData)

	// compare form data to database
	dbPasswordHash := ""
	err := Db.QueryRow(
		"SELECT password FROM users WHERE username = $1;",
		formData.Username,
	).Scan(&dbPasswordHash)
	// fmt.Println(dbPasswordHash)
	// check if user does not exist
	if err != nil {
		fmt.Println("login() query failed:", err)
		return
	}

	pwMatchErr := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(formData.Password))
	if pwMatchErr != nil {
		fmt.Println("password hashes don't match:", pwMatchErr)
		return
	} else {
		// generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
		tokenPair, err := GenerateToken(formData.Userid, formData.Username)
		if err != nil {
			fmt.Println("failed to generate token pair:", err)
			return
		}

		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "refreshToken",
			Value:    tokenPair.RefreshToken,
			Domain:   "jpoly1219devbox.xyz",
			Path:     "/auth/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)
		w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
		json.NewEncoder(w).Encode(tokenPair.AccessToken)
	}
}

func SilentRefresh(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

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
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)
		w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
		json.NewEncoder(w).Encode(tokenPair.AccessToken)
	} else {
		fmt.Println(err)
	}
}
