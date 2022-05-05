package methods

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

func (b *Bkash) CreatePayment(request *models.CreateRequest, token *models.TokenResponse) (*models.CreatePaymentResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.Mode == "" || request.CallbackURL == "" {
		return nil, common.ErrEmptyRequiredField
	}

	// Mode validation
	if request.Mode != "0001" && request.Mode != "0011" {
		return nil, errors.New("invalid mode value")
	}

	createPaymentURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_CREATE_PAYMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, createPaymentURL, true)
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

func (b *Bkash) ExecutePayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.ExecutePaymentResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	executePayment := hooks.GenerateURI(b.IsLiveStore, common.BKASH_EXECUTE_PAYMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, executePayment, true)
	if err != nil {
		// if error is timeout then call query payment
		// if complete return success payload (*models.ExecutePaymentResponse, nil)
		// if initiated - return something that should be handled by client (maybe return some kind of timeout error)
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

func (b *Bkash) QueryPayment(request *models.ExecuteRequest, token *models.TokenResponse) (*models.QueryPaymentResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || request.PaymentID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	queryPaymentURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_QUERY_PAYMENT_URI)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, queryPaymentURL, true)
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
