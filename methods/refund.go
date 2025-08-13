package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

// This function can be used to refund a payment or check the status of a refund
func (b *Bkash) Refund(r *models.RefundRequest, t *models.TokenResponse) (*models.RefundResponse, error) {
	if b.AppKey == "" || t.IdToken == "" || r.PaymentID == "" || r.TrxID == "" {
		return nil, common.ErrEmptyRequiredField
	}

	refundURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_REFUND_URI)

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    r,
		Username:   t.IdToken,
		Password:   b.AppKey,
		Url:        refundURL,
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
