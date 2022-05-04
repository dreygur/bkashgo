package hooks

import (
	"net/url"

	"github.com/dreygur/bkashgo/common"
)

func GenerateURI(IsLiveStore bool, uri string) *url.URL {
	var storeUrl string
	if IsLiveStore {
		storeUrl = common.BKASH_LIVE_GATEWAY
	} else {
		storeUrl = common.BKASH_SANDBOX_GATEWAY
	}

	u, _ := url.ParseRequestURI(storeUrl)
	u.Path += uri

	return u
}
