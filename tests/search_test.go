package tests

import (
	"testing"

	"github.com/dreygur/bkashgo"
)

func TestSearch(t *testing.T) {
	bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

	token, err := bkash.GetToken()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Error("invalid token")
	}

	res, err := bkash.Search("9DT9NBLAOT", token)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if res == nil || res.StatusMessage == "" || res.StatusCode != "0000" {
		t.Error("invalid response")
	}
}
