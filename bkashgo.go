package bkashgo

import "github.com/dreygur/bkashgo/methods"

func GetBkash(username, password, appKey, appSecret string, isLiveStore bool) methods.BkashTokenizedCheckoutService {
	return &methods.Bkash{
		Username:  username,
		Password:  password,
		AppKey:    appKey,
		AppSecret: appSecret,

		IsLiveStore: isLiveStore,
	}
}
