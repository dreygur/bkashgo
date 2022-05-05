package methods

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

func (b *Bkash) CreateAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CreateAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.Mode == "" || request.CallbackURL == "" {
		return nil, common.ErrEmptyRequiredField
	}

	// Mode validation
	if request.Mode != "0000" {
		return nil, errors.New("invalid mode value")
	}

	createAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_CREATE_AGREEMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, createAgreementURL, true)
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

func (b *Bkash) CreateAgreementValidationListener(r *http.Request) (*models.CreateAgreementValidationResponse, error) {
	if r.Method != "POST" {
		return nil, errors.New("method not allowed")
	}

	var agreementTValidationResponse models.CreateAgreementValidationResponse

	err := json.NewDecoder(r.Body).Decode(&agreementTValidationResponse)
	if err != nil {
		return nil, err
	}

	return &agreementTValidationResponse, nil
}

func (b *Bkash) ExecuteAgreement(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecuteAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	executeAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_EXECUTE_AGREEMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, executeAgreementURL, true)
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

func (b *Bkash) QueryAgreement(request *models.AgreementRequest, token *models.TokenResponse) (*models.QueryAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	queryAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_QUERY_AGREEMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, queryAgreementURL, true)
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

func (b *Bkash) CancelAgreement(request *models.CreateRequest, token *models.TokenResponse) (*models.CancelAgreementResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.AgreementID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	cancelAgreementURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_CANCEL_AGREEMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, cancelAgreementURL, true)
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
