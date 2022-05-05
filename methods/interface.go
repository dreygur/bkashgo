package methods

import "github.com/dreygur/bkashgo/models"

type Bkash struct {
	Username  string
	Password  string
	AppKey    string
	AppSecret string

	IsLiveStore bool
}

type BkashTokenizedCheckoutService interface {
	GrantToken() (*models.TokenResponse, error)
	RefreshToken(token *models.TokenRequest) (*models.TokenResponse, error)
}
