package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

// GetToken creates a access token using bkash credentials
func (b *Bkash) GetToken() (*models.TokenResponse, error) {
	if b.AppKey == "" || b.AppSecret == "" || b.Username == "" || b.Password == "" {
		return nil, common.ErrEmptyRequiredField
	}

	data := models.TokenRequest{
		AppKey:    b.AppKey,
		AppSecret: b.AppSecret,
	}

	grantTokenURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_GRANT_TOKEN_URI)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, b.Username, b.Password, grantTokenURL, false)
	if err != nil {
		return nil, err
	}

	var resp models.TokenResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// RefreshToken refreshes the access token
func (b *Bkash) RefreshToken(token *models.TokenRequest) (*models.TokenResponse, error) {
	if b.AppKey == "" || b.AppSecret == "" || token.RefreshToken == "" || b.Username == "" || b.Password == "" {
		return nil, common.ErrEmptyRequiredField
	}

	data := models.TokenRequest{
		AppKey:       b.AppKey,
		AppSecret:    b.AppSecret,
		RefreshToken: token.RefreshToken,
	}

	refreshTokenURL := hooks.GenerateURI(b.IsLiveStore, common.BKASH_REFRESH_TOKEN_URI)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, b.Username, b.Password, refreshTokenURL, false)
	if err != nil {
		return nil, err
	}

	var resp models.TokenResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
