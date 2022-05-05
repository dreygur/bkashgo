package main

import (
	"fmt"

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

func main() {
	bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)
	token, err := bkash.GetToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	refreshToken, err := bkash.RefreshToken(&models.TokenRequest{RefreshToken: token.RefreshToken})
	if err != nil {
		panic(err)
	}
	fmt.Println(refreshToken)
}
