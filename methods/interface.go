package methods

import (
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
	// Token
	GetToken() (*models.TokenResponse, error)
	RefreshToken(token *models.TokenRequest) (*models.TokenResponse, error)

	// Agreement
	CreateAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CreateAgreementResponse, error)
	ExecuteAgreement(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecuteAgreementResponse, error)
	QueryAgreement(request *models.AgreementRequest, token *models.TokenResponse) (*models.QueryAgreementResponse, error)
	CancelAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CancelAgreementResponse, error)

	// Payments
	CreatePayment(request *models.CreateRequest, token *models.TokenResponse) (*models.CreatePaymentResponse, error)
	ExecutePayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecutePaymentResponse, error)
	QueryPayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.QueryPaymentResponse, error)

	// Search Transactions
	Search(trxID string, token *models.TokenResponse) (*models.SearchTransactionResponse, error)

	// Refund
	// This function can be used to refund a payment or check the status of a refund
	Refund(r *models.RefundRequest, t *models.TokenResponse) (*models.RefundResponse, error)
}
