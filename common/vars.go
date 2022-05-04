package common

import "errors"

const BKASH_SANDBOX_GATEWAY = "https://tokenized.sandbox.bka.sh/v1.2.0-beta"
const BKASH_LIVE_GATEWAY = "https://tokenized.pay.bka.sh/v1.2.0-beta"
const BKASH_GRANT_TOKEN_URI = "/tokenized/checkout/token/grant"
const BKASH_REFRESH_TOKEN_URI = "/tokenized/checkout/token/refresh"
const BKASH_CREATE_AGREEMENT_URI = "/tokenized/checkout/create"
const BKASH_EXECUTE_AGREEMENT_URI = "/tokenized/checkout/execute"
const BKASH_QUERY_AGREEMENT_URI = "/tokenized/checkout/agreement/status"
const BKASH_CANCEL_AGREEMENT_URI = "/tokenized/checkout/agreement/cancel"
const BKASH_CREATE_PAYMENT_URI = "/tokenized/checkout/create"
const BKASH_EXECUTE_PAYMENT_URI = "/tokenized/checkout/execute"
const BKASH_QUERY_PAYMENT_URI = "/tokenized/checkout/payment/status"

var ErrEmptyRequiredField = errors.New("empty required field")
var ErrTimeout = errors.New("api request timeout")
