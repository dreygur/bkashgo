package common

import "errors"

const BKASH_SANDBOX_GATEWAY = "https://tokenized.sandbox.bka.sh/v1.2.0-beta/tokenized/checkout"
const BKASH_LIVE_GATEWAY = "https://tokenized.pay.bka.sh/v1.2.0-beta/tokenized/checkout"
const BKASH_GRANT_TOKEN_URI = "/token/grant"
const BKASH_REFRESH_TOKEN_URI = "/token/refresh"
const BKASH_CREATE_AGREEMENT_URI = "/create"
const BKASH_EXECUTE_AGREEMENT_URI = "/execute"
const BKASH_QUERY_AGREEMENT_URI = "/agreement/status"
const BKASH_CANCEL_AGREEMENT_URI = "/agreement/cancel"
const BKASH_CREATE_PAYMENT_URI = "/create"
const BKASH_EXECUTE_PAYMENT_URI = "/execute"
const BKASH_QUERY_PAYMENT_URI = "/payment/status"
const BKASH_SEARCH_TRXN_URI = "/general/searchTransaction"

var ErrEmptyRequiredField = errors.New("empty required field")
var ErrTimeout = errors.New("api request timeout")
