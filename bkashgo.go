// Package bkashgo provides the methods for tokenized checkout
package bkashgo

import "github.com/dreygur/bkashgo/methods"

// Initiates a new BkashGo object
// It also publishes sets of methods for tokenized checkout
func GetBkash(username, password, appKey, appSecret string, isLiveStore bool) methods.BkashTokenizedCheckoutService {
	return &methods.Bkash{
		Username:  username,
		Password:  password,
		AppKey:    appKey,
		AppSecret: appSecret,

		IsLiveStore: isLiveStore,
	}
}
