package methods

import (
	"encoding/json"
	"errors"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

// Debug
func (b *Bkash) Debug(enable bool) *Bkash {
	b.debug = enable
	return b
}

// CreateAgreement Initiates an agreement request for a user
func (b *Bkash) CreateAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CreateAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.Mode == "" || request.CallbackURL == "" {
		return nil, common.ErrEmptyRequiredField
	}

	// Mode validation
	if request.Mode != "0000" {
		return nil, errors.New("invalid mode value")
	}

	createAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_CREATE_AGREEMENT_URI)

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        createAgreementURL,
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.CreateAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ExecuteAgreement executes the agreement using the paymentID received from CreateAgreementResponse
func (b *Bkash) ExecuteAgreement(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecuteAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	executeAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_EXECUTE_AGREEMENT_URI)

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        executeAgreementURL,
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.ExecuteAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// QueryAgreement queries an agreement by agreementID
func (b *Bkash) QueryAgreement(request *models.AgreementRequest, token *models.TokenResponse) (*models.QueryAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	queryAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_QUERY_AGREEMENT_URI)

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        queryAgreementURL,
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.QueryAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// CancelAgreement cancels an agreement by agreementID
func (b *Bkash) CancelAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CancelAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	cancelAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_CANCEL_AGREEMENT_URI)

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        cancelAgreementURL,
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.CancelAgreementResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
