package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

func (b *Bkash) Search(trxID string, token *models.TokenResponse) (*models.SearchTransactionResponse, error) {
	if b.AppKey == "" || token.IdToken == "" || trxID == "" {
		return nil, common.ErrEmptyRequiredField
	}
	executePayment := hooks.GenerateURI(b.IsLiveStore, common.BKASH_SEARCH_TRXN_URI)

	request := models.SearchTransactionRequest{
		TrxID: trxID,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, token.IdToken, b.AppKey, executePayment, true)
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
