package methods

import (
	"encoding/json"
	"net/url"

	"github.com/dreygur/bkashgo/hooks"
	"github.com/dreygur/bkashgo/models"
)

func (b *Bkash) GrantToken() (*models.TokenResponse, error) {
	// Mandatory field validation
	if b.AppKey == "" || b.AppSecret == "" || b.Username == "" || b.Password == "" {
		return nil, ErrEmptyRequiredField
	}

	var data = make(map[string]string)

	data["app_key"] = b.AppKey
	data["app_secret"] = b.AppSecret

	var storeUrl string
	if b.IsLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_GRANT_TOKEN_URI

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
	// Mandatory field validation
	if b.AppKey == "" || b.AppSecret == "" || token.RefreshToken == "" || b.Username == "" || b.Password == "" {
		return nil, ErrEmptyRequiredField
	}

	data := models.TokenRequest{
		AppKey:       b.AppKey,
		AppSecret:    b.AppSecret,
		RefreshToken: token.RefreshToken,
	}

	var storeUrl string
	if b.IsLiveStore {
		storeUrl = BKASH_LIVE_GATEWAY
	} else {
		storeUrl = BKASH_SANDBOX_GATEWAY
	}
	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += BKASH_REFRESH_TOKEN_URI

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
