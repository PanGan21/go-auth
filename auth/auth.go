package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// TokenDetails defines the details of a token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUUID    string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails defines the details needed to aceess the api
type AccessDetails struct {
	TokenUUID string
	UserID    string
}

// AuthInterface is the authentication interface
type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type service struct {
	client *redis.Client
}

// Auth implements the AuthInterface
var _ AuthInterface = &service{}

// NewAuth returns a new Authentication service with redis client injected
func NewAuth(client *redis.Client) *service {
	return &service{client: client}
}

// Saves token metadata to Redis
func (tk *service) CreateAuth(userID string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(td.TokenUUID, userID, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(td.RefreshUUID, userID, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

func (tk *service) FetchAuth(tokenUUID string) (string, error) {
	userid, err := tk.client.Get(tokenUUID).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (tk *service) DeleteRefresh(refreshUUID string) error {
	//delete refresh token
	deleted, err := tk.client.Del(refreshUUID).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}

func (tk *service) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUUID := fmt.Sprintf("%s++%s", authD.TokenUUID, authD.UserID)
	//delete access token
	deletedAt, err := tk.client.Del(authD.TokenUUID).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(refreshUUID).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
