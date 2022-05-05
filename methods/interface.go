package methods

import (
	"net/http"

	"github.com/dreygur/bkashgo/models"
)

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
	CreateAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CreateAgreementResponse, error)
	CreateAgreementValidationListener(r *http.Request) (*models.CreateAgreementValidationResponse, error)
	ExecuteAgreement(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecuteAgreementResponse, error)
	QueryAgreement(request *models.AgreementRequest, token *models.TokenResponse) (*models.QueryAgreementResponse, error)
	CancelAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CancelAgreementResponse, error)
}
