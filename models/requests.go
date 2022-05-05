// Package models provides the structs for the response
package models

// Grant/Refresh Token Request
type TokenRequest struct {
	AppKey       string `json:"app_key,omitempty"`
	AppSecret    string `json:"app_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Create Agreement/Payment Request
type CreateRequest struct {
	Mode                    string `json:"mode,omitempty"`
	PayerReference          string `json:"payerReference,omitempty"`
	CallbackURL             string `json:"callbackURL,omitempty"`
	Amount                  string `json:"amount,omitempty"`
	Currency                string `json:"currency,omitempty"`
	Intent                  string `json:"intent,omitempty"`
	MerchantInvoiceNumber   string `json:"merchantInvoiceNumber,omitempty"`
	AgreementID             string `json:"agreementID,omitempty"`
	MerchantAssociationInfo string `json:"merchantAssociationInfo,omitempty"`
}

// Execute Agreement/Payment/Query Request
type ExecuteRequest struct {
	PaymentID string `json:"paymentID,omitempty"`
}

// Query/Cancel Agreement Request
type AgreementRequest struct {
	AgreementID string `json:"agreementID,omitempty"`
}

// Search Transaction Request
type SearchTransactionRequest struct {
	TrxID string `json:"trxID,omitempty"`
}

// Refund Transaction/Status Request
type RefundRequest struct {
	PaymentID string `json:"paymentID,omitempty"`
	Amount    string `json:"amount,omitempty"`
	TrxID     string `json:"trxID,omitempty"`
	SKU       string `json:"sku,omitempty"`
	Reason    string `json:"reason,omitempty"`
}
