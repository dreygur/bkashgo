package models

type RefundStatusRequest struct {
	PaymentID string `json:"paymentID,omitempty"`
	TrxID     string `json:"trxID,omitempty"`
}

type RefundTransactionRequest struct {
	PaymentID string `json:"paymentID,omitempty"`
	Amount    string `json:"amount,omitempty"`
	TrxID     string `json:"trxID,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Reason    string `json:"reason,omitempty"`
}
