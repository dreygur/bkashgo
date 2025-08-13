package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
	"github.com/dreygur/bkashgo/utils"
)

// This function can be used to refund a payment or check the status of a refund
func (b *Bkash) Refund(request *models.RefundRequest, token *models.TokenResponse) (*models.RefundResponse, error) {
	if !utils.RequireNonEmpty(b.AppKey, token.IdToken, request.PaymentID, request.TrxID) {
		return nil, common.ErrEmptyRequiredField
	}

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        hooks.GenerateURI(b.IsLiveStore, common.BKASH_REFUND_URI),
		Authorized: true,
	}

	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.RefundResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
