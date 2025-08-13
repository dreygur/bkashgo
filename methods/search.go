package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

// Search for Transactions
func (b *Bkash) Search(trxID string, token *models.TokenResponse) (*models.SearchTransactionResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || trxID == "" {
		return nil, common.ErrEmptyRequiredField
	}
	executePayment := hooks.GenerateURI(b.IsLiveStore, common.BKASH_SEARCH_TRXN_URI)

	request := models.SearchTransactionRequest{
		TrxID: trxID,
	}

	payload := &hooks.Request{
		Debug:      b.debug,
		Payload:    request,
		Username:   token.IdToken,
		Password:   b.AppKey,
		Url:        executePayment,
		Authorized: true,
	}
	body, err := hooks.DoRequest(payload)
	if err != nil {
		return nil, err
	}

	var resp models.SearchTransactionResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
