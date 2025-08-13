// Package methods provides the methods for tokenized checkout
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
	debug       bool
}

type BkashTokenizedCheckoutService interface {
	// GetToken creates a access token using bkash credentials
	GetToken() (*models.TokenResponse, error)
	// RefreshToken refreshes the access token
	RefreshToken(token *models.TokenRequest) (*models.TokenResponse, error)
	// CreateAgreement Initiates an agreement request for an user
	CreateAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CreateAgreementResponse, error)
	// ExecuteAgreement executes the agreement using the paymentID received from CreateAgreementResponse
	ExecuteAgreement(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecuteAgreementResponse, error)
	// QueryAgreement queries an agreement by agreementID
	QueryAgreement(request *models.AgreementRequest, token *models.TokenResponse) (*models.QueryAgreementResponse, error)
	// CancelAgreement cancels an agreement by agreementID
	CancelAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CancelAgreementResponse, error)

	// CreatePayment Initiates a payment request for an user
	CreatePayment(request *models.CreateRequest, token *models.TokenResponse) (*models.CreatePaymentResponse, error)
	// ExecutePayment executes the agreement using the paymentID received from CreateAgreementResponse
	ExecutePayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecutePaymentResponse, error)
	// QueryPayment queries a payment by paymentID
	QueryPayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.QueryPaymentResponse, error)

	// Search for Transactions
	Search(trxID string, token *models.TokenResponse) (*models.SearchTransactionResponse, error)

	// This function can be used to refund a payment or check the status of a refund
	Refund(r *models.RefundRequest, t *models.TokenResponse) (*models.RefundResponse, error)

	// This method allows to get RAW JSON for Request/Response
	Debug(enable bool) *Bkash
}
