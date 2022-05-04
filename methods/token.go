package methods

import (
	"encoding/json"

	"github.com/dreygur/bkashgo/common"
	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

func (b *Bkash) GrantToken() (*models.TokenResponse, error) {
	if b.AppKey == "" || b.AppSecret == "" || b.Username == "" || b.Password == "" {
		return nil, common.ErrEmptyRequiredField
	}

	data := models.TokenRequest{
		AppKey:    b.AppKey,
		AppSecret: b.AppSecret,
	}

	u := hooks.GenerateURI(b.IsLiveStore, common.BKASH_GRANT_TOKEN_URI)
	grantTokenURL := u.String()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, b.Username, b.Password, grantTokenURL)
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

func (b *Bkash) RefreshToken(token *models.TokenRequest) (*models.TokenResponse, error) {
	if b.AppKey == "" || b.AppSecret == "" || token.RefreshToken == "" || b.Username == "" || b.Password == "" {
		return nil, common.ErrEmptyRequiredField
	}

	data := models.TokenRequest{
		AppKey:       b.AppKey,
		AppSecret:    b.AppSecret,
		RefreshToken: token.RefreshToken,
	}

	u := hooks.GenerateURI(b.IsLiveStore, common.BKASH_REFRESH_TOKEN_URI)
	refreshTokenURL := u.String()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := hooks.DoRequest(jsonData, b.Username, b.Password, refreshTokenURL)
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
