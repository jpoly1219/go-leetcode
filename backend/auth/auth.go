package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/jpoly1219/go-leetcode/backend/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func GenerateToken(userid uuid.UUID, username string) (*models.Token, error) {
	accessKey := os.Getenv("ACCESSSECRETKEY")
	refreshKey := os.Getenv("REFRESHSECRETKEY")
	accessExp := time.Now().Add(time.Hour * 24).Unix()
	refreshExp := time.Now().Add(time.Hour * 24 * 7).Unix()

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

	tokenPair := models.Token{AccessToken: accessTokenString, RefreshToken: refreshTokenString}
	return &tokenPair, nil
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// handle preflight OPTIONS request
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// read form data and check if form is valid
	var formData models.User
	json.NewDecoder(r.Body).Decode(&formData)
	fmt.Println(formData)

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	if err != nil {
		fmt.Println("failed to generate password hash", err)
		return
	}

	// insert user data to database
	var userid uuid.UUID
	username := ""
	defaultProfilePic := "https://isobarscience.com/wp-content/uploads/2020/09/default-profile-picture1.jpg"

	err = models.Db.QueryRow(
		"INSERT INTO users (username, fullname, email, password, profile_pic) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, username;",
		formData.Username, formData.Fullname, formData.Email, string(passwordHash), defaultProfilePic,
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
		Domain:   "54.145.220.238:3000",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
	json.NewEncoder(w).Encode(tokenPair.AccessToken)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// handle preflight OPTIONS request
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// read form data and check if form is valid
	var formData models.User
	json.NewDecoder(r.Body).Decode(&formData)
	// fmt.Println(formData)

	// compare form data to database
	dbPasswordHash := ""
	err := models.Db.QueryRow(
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
			Domain:   "54.145.220.238:3000",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)
		w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
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
	fmt.Println("refreshToken value: ", refreshToken)

	// check if refresh token is valid
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESHSECRETKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if so then generate token pair and send it to user. Access token and exp as JSON, refresh token as HttpOnly cookie.
		userid := claims["userid"]
		fmt.Println(userid, reflect.TypeOf(userid))
		username := claims["username"]
		useridStr, ok := userid.(uuid.UUID)
		if !ok {
			fmt.Println("Error: %v")
		}
		usernameStr := username.(string)
		tokenPair, err := GenerateToken(useridStr, usernameStr)
		if err != nil {
			fmt.Println("Error generating token: ", err)
			return
		}

		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "refreshToken",
			Value:    tokenPair.RefreshToken,
			Domain:   "frontend",
			Path:     "/auth/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)
		w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
		json.NewEncoder(w).Encode(tokenPair.AccessToken)
	} else {
		fmt.Println(err)
	}
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// read form data and check if form is valid
	var formData models.NewUserData
	json.NewDecoder(r.Body).Decode(&formData)
	fmt.Println(formData)

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(formData.NewPassword), 14)
	if err != nil {
		fmt.Println("failed to generate password hash", err)
		return
	}

	var userid uuid.UUID
	var username string
	// update user data to database
	err = models.Db.QueryRow(
		"UPDATE users SET username = $1, fullname = $2, email = $3, password = $4 WHERE username = $5 RETURNING user_id, username;",
		formData.NewUsername, formData.NewFullname, formData.NewEmail, string(passwordHash), formData.OldUsername,
	).Scan(&userid, &username)
	if err != nil {
		fmt.Println("INSERT to EditProfile failed:", err)
		w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
		w.Write([]byte("Request failed. Username or email already exists."))
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
		Domain:   "54.145.220.238:3000",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Access-Control-Allow-Origin", "http://54.145.220.238:3000")
	json.NewEncoder(w).Encode(tokenPair.AccessToken)
}
