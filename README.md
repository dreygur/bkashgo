# bKashGo

This go SDK aim to implement bKash tokenized api

[![Go Reference](https://pkg.go.dev/badge/github.com/dreygur/bkashgo.svg)](https://pkg.go.dev/github.com/dreygur/bkashgo)  [![CodeQL](https://github.com/dreygur/bkashgo/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dreygur/bkashgo/actions/workflows/codeql-analysis.yml)

Ref:
- [Integration Guide](https://developer.bka.sh/docs/tokenized-checkout-process)
- [Specification/Reference](https://developer.bka.sh/reference)

## Feautures
### Tokenized Checkout
- Grant Token
- Refresh Token
- Create Agreement
- Execute Agreement
- Query Agreement
- Cancel Agreement
- Create Payment
- Execute Payment
- Query Payment
- Search Transaction
- Refund Transaction
- Refund Status

## Examples:

__To Generate a Token__
```go
// bKash Instance
bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

// Token and error message
token, err := bkash.GetToken()
if err != nil {
  panic(err)
}
fmt.Println(token)
```

__To Generate Refresh Token__
```go
// bKash Instance
bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

// Token and error message
token, err := bkash.GetToken()
if err != nil {
  panic(err)
}

refreshToken, err := bkash.RefreshToken(&models.TokenRequest{
  RefreshToken: token.RefreshToken
})
if err != nil {
  panic(err)
}
fmt.Println(refreshToken)
```

Made with ❤️ by [Rakibul Yeasin](https://facebook.com/dreygur)
