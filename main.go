package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("works")
}

// User type and json mapping
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

// CreateToken creates and returns a token with expiration time 15 minutes
func CreateToken(userID int64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
