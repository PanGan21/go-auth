package auth

import "github.com/go-redis/redis"

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
// var _ AuthInterface = &service{}

// NewAuth returns a new Authentication service with redis client injected
func NewAuth(client *redis.Client) *service {
	return &service{client: client}
}
