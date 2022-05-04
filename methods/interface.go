package methods

type Bkash struct {
	Username  string
	Password  string
	AppKey    string
	AppSecret string

	IsLiveStore bool
}

type BkashTokenizedCheckoutService interface{}
