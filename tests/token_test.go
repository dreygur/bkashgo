package tests

import (
	"testing"

	"github.com/dreygur/bkashgo"
	"github.com/dreygur/bkashgo/models"
)

var (
	username    = "sandboxTokenizedUser02"
	password    = "sandboxTokenizedUser02@12345"
	appKey      = "4f6o0cjiki2rfm34kfdadl1eqq"
	appSecret   = "2is7hdktrekvrbljjh44ll3d9l1dtjo4pasmjvs5vl5qr3fug4b"
	isLiveStore = false
)

func TestGrantToken(t *testing.T) {
	bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

	token, err := bkash.GrantToken()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Error("invalid token")
	}

	t.Run("test RefreshToken", func(t *testing.T) {
		refreshToken, err := bkash.RefreshToken(&models.TokenRequest{
			RefreshToken: token.RefreshToken,
		})
		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}
		if refreshToken == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
			t.Error("invalid token")
		}
	})
}
