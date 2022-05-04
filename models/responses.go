package models

// Error response
type BkashError struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Grant/Refresh Token Response
type TokenResponse struct {
	ExpiresIn     int    `json:"expires_in,omitempty"`
	IdToken       string `json:"id_token,omitempty"`
	RefreshToken  string `json:"refresh_token,omitempty"`
	TokenType     string `json:"token_type,omitempty"`
	StatusCode    string `json:"statusCode,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

type CreateAgreementValidationResponse struct {
	PaymentID string `json:"paymentID,omitempty"`
	Status    string `json:"status,omitempty"`
}

// CreateAgreement Response
type CreateAgreementResponse struct {
	PaymentID            string `json:"paymentID,omitempty"`
	BkashURL             string `json:"bkashURL,omitempty"`
	CallbackURL          string `json:"callbackURL,omitempty"`
	SuccessCallbackURL   string `json:"successCallbackURL,omitempty"`
	FailureCallbackURL   string `json:"failureCallbackURL,omitempty"`
	CancelledCallbackURL string `json:"cancelledCallbackURL,omitempty"`
	StatusCode           string `json:"statusCode,omitempty"`
	StatusMessage        string `json:"statusMessage,omitempty"`
}

// ExecuteAgreement Response
type ExecuteAgreementResponse struct {
	PaymentID            string `json:"paymentID,omitempty"`
	AgreementID          string `json:"agreementID,omitempty"`
	CustomerMsisdn       string `json:"customerMsisdn,omitempty"`
	PayerReference       string `json:"payerReference,omitempty"`
	AgreementExecuteTime string `json:"agreementExecuteTime,omitempty"`
	AgreementStatus      string `json:"agreementStatus,omitempty"`
	StatusCode           string `json:"statusCode,omitempty"`
	StatusMessage        string `json:"statusMessage,omitempty"`
}

// QueryAgreement Response
type QueryAgreementResponse struct {
	PaymentID            string `json:"paymentID,omitempty"`
	AgreementID          string `json:"agreementID,omitempty"`
	PayerReference       string `json:"payerReference,omitempty"`
	CustomerMsisdn       string `json:"customerMsisdn,omitempty"`
	AgreementCreateTime  string `json:"agreementCreateTime,omitempty"`
	AgreementExecuteTime string `json:"agreementExecuteTime,omitempty"`
	AgreementVoidTime    string `json:"agreementVoidTime,omitempty"`
	AgreementStatus      string `json:"agreementStatus,omitempty"`
	StatusCode           string `json:"statusCode,omitempty"`
	StatusMessage        string `json:"statusMessage,omitempty"`
}

// CancelAgreement Response
type CancelAgreementResponse struct {
	PaymentID         string `json:"paymentID,omitempty"`
	AgreementID       string `json:"agreementID,omitempty"`
	PayerReference    string `json:"payerReference,omitempty"`
	AgreementVoidTime string `json:"agreementVoidTime,omitempty"`
	AgreementStatus   string `json:"agreementStatus,omitempty"`
	StatusCode        string `json:"statusCode,omitempty"`
	StatusMessage     string `json:"statusMessage,omitempty"`
}

// CreatePayment Response
type CreatePaymentResponse struct {
	PaymentID             string `json:"paymentID,omitempty"`
	AgreementID           string `json:"agreementID,omitempty"`
	PaymentCreateTime     string `json:"paymentCreateTime,omitempty"`
	TransactionStatus     string `json:"transactionStatus,omitempty"`
	Amount                string `json:"amount,omitempty"`
	Currency              string `json:"currency,omitempty"`
	Intent                string `json:"intent,omitempty"`
	MerchantInvoiceNumber string `json:"merchantInvoiceNumber,omitempty"`
	BkashURL              string `json:"bkashURL,omitempty"`
	CallbackURL           string `json:"callbackURL,omitempty"`
	SuccessCallbackURL    string `json:"successCallbackURL,omitempty"`
	FailureCallbackURL    string `json:"failureCallbackURL,omitempty"`
	CancelledCallbackURL  string `json:"cancelledCallbackURL,omitempty"`
	StatusCode            string `json:"statusCode,omitempty"`
	StatusMessage         string `json:"statusMessage,omitempty"`
}

// ExecutePayment Response
type ExecutePaymentResponse struct {
	PaymentID             string `json:"paymentID,omitempty"`
	AgreementID           string `json:"agreementID,omitempty"`
	CustomerMsisdn        string `json:"customerMsisdn,omitempty"`
	PayerReference        string `json:"payerReference,omitempty"`
	AgreementExecuteTime  string `json:"agreementExecuteTime,omitempty"`
	AgreementStatus       string `json:"agreementStatus,omitempty"`
	PaymentExecuteTime    string `json:"paymentExecuteTime,omitempty"`
	TrxID                 string `json:"trxID,omitempty"`
	TransactionStatus     string `json:"transaction_status,omitempty"`
	Amount                string `json:"amount,omitempty"`
	Currency              string `json:"currency,omitempty"`
	Intent                string `json:"intent,omitempty"`
	MerchantInvoiceNumber string `json:"merchantInvoiceNumber,omitempty"`
	StatusCode            string `json:"statusCode,omitempty"`
	StatusMessage         string `json:"statusMessage,omitempty"`
}

// QueryPayment Response
type QueryPaymentResponse struct {
	PaymentID              string `json:"paymentID,omitempty"`
	Mode                   string `json:"mode,omitempty"`
	PayerReference         string `json:"payerReference,omitempty"`
	PaymentCreateTime      string `json:"paymentCreateTime,omitempty"`
	PaymentExecuteTime     string `json:"paymentExecuteTime,omitempty"`
	TrxID                  string `json:"trxID,omitempty"`
	TransactionStatus      string `json:"transaction_status,omitempty"`
	Amount                 string `json:"amount,omitempty"`
	Currency               string `json:"currency,omitempty"`
	Intent                 string `json:"intent,omitempty"`
	MerchantInvoiceNumber  string `json:"merchantInvoiceNumber,omitempty"`
	UserVerificationStatus string `json:"userVerificationStatus,omitempty"`
	StatusCode             string `json:"statusCode,omitempty"`
	StatusMessage          string `json:"statusMessage,omitempty"`
}

// SearchTransaction Response
type SearchTransactionResponse struct {
	Amount                string `json:"amount,omitempty"`
	CompletedTime         string `json:"completed_time,omitempty"`
	Currency              string `json:"currency,omitempty"`
	CustomerMsisdn        string `json:"customerMsisdn,omitempty"`
	InitiationTime        string `json:"initiationTime,omitempty"`
	OrganizationShortCode string `json:"organizationShortCode,omitempty"`
	TransactionReference  string `json:"transactionReference,omitempty"`
	TransactionStatus     string `json:"transactionStatus,omitempty"`
	TransactionType       string `json:"transactionType,omitempty"`
	StatusCode            string `json:"statusCode,omitempty"`
	StatusMessage         string `json:"statusMessage,omitempty"`
	TrxID                 string `json:"trxID,omitempty"`
}

// Refund Transaction/Status Response
type RefundResponse struct {
	CompletedTime     string `json:"completedTime,omitempty"`
	TransactionStatus string `json:"transactionStatus,omitempty"`
	OriginalTrxID     string `json:"originalTrxID,omitempty"`
	RefundTrxID       string `json:"refundTrxID,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Currency          string `json:"currency,omitempty"`
	Charge            string `json:"charge,omitempty"`
}
