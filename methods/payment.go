package methods

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
	"github.com/dreygur/bkashgo/utils"
)

// CreatePayment Initiates a payment request for an user
func (b *Bkash) CreatePayment(request *models.CreateRequest, token *models.TokenResponse) (*models.CreatePaymentResponse, error) {
	if !utils.RequireNonEmpty(b.AppKey, token.IdToken, request.Mode, request.CallbackURL) {
		return nil, common.ErrEmptyRequiredField
	}

	// Mode validation
	if request.Mode != "0001" && request.Mode != "0011" {
		return nil, errors.New("invalid mode value")
	}

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        hooks.GenerateURI(b.IsLiveStore, common.BKASH_CREATE_PAYMENT_URI),
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.CreatePaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ExecutePayment executes the agreement using the paymentID received from CreateAgreementResponse
func (b *Bkash) ExecutePayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecutePaymentResponse, error) {
	if !utils.RequireNonEmpty(b.AppKey, token.IdToken, request.PaymentID) {
		return nil, common.ErrEmptyRequiredField
	}

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        hooks.GenerateURI(b.IsLiveStore, common.BKASH_EXECUTE_PAYMENT_URI),
		Authorized: true,
	}

	body, err := hooks.DoRequest(payload)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			queryResp, err := b.QueryPayment(&models.ExecuteRequest{
				PaymentID: request.PaymentID,
			}, token)
			if err != nil {
				return nil, err
			}

			if queryResp.StatusCode == "0000" && queryResp.TransactionStatus == "Completed" {
				return &models.ExecutePaymentResponse{
					PaymentID:             queryResp.PaymentID,
					PayerReference:        queryResp.PayerReference,
					PaymentExecuteTime:    queryResp.PaymentExecuteTime,
					TrxID:                 queryResp.TrxID,
					TransactionStatus:     queryResp.TransactionStatus,
					Amount:                queryResp.Amount,
					Currency:              queryResp.Currency,
					Intent:                queryResp.Intent,
					MerchantInvoiceNumber: queryResp.MerchantInvoiceNumber,
					StatusCode:            queryResp.StatusCode,
					StatusMessage:         queryResp.StatusMessage,
				}, nil
			} else {
				return nil, common.ErrTimeout
			}
		} else {
			return nil, err
		}
	}

	var resp models.ExecutePaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// QueryPayment queries a payment by paymentID
func (b *Bkash) QueryPayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.QueryPaymentResponse, error) {
	if !utils.RequireNonEmpty(b.AppKey, token.IdToken, request.PaymentID) {
		return nil, common.ErrEmptyRequiredField
	}

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        hooks.GenerateURI(b.IsLiveStore, common.BKASH_QUERY_PAYMENT_URI),
		Authorized: true,
	}

	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.QueryPaymentResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
