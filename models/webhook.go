package models

// BkashIPN Request/Response payload
type BkashIPNPayload struct {
	Type             string      `json:"Type,omitempty"`
	MessageId        string      `json:"MessageId,omitempty"`
	Token            string      `json:"Token,omitempty"`
	TopicArn         string      `json:"TopicArn,omitempty"`
	SubscribeURL     string      `json:"SubscribeURL,omitempty"`
	Timestamp        string      `json:"Timestamp,omitempty"`
	SignatureVersion string      `json:"SignatureVersion,omitempty"`
	Signature        string      `json:"Signature,omitempty"`
	SigningCertURL   string      `json:"SigningCertURL,omitempty"`
	UnSubscribeURL   string      `json:"UnSubscribeURL,omitempty"`
	Message          interface{} `json:"Message,omitempty"`
}

// Message for response
type Message struct {
	DateTime               string `json:"dateTime,omitempty"`
	DebitMSISDN            string `json:"debitMSISDN,omitempty"`
	CreditOrganizationName string `json:"creditOrganizationName,omitempty"`
	CreditShortCode        string `json:"creditShortCode,omitempty"`
	TrxID                  string `json:"trxID,omitempty"`
	TransactionStatus      string `json:"transactionStatus,omitempty"`
	TransactionType        string `json:"transactionType,omitempty"`
	Amount                 string `json:"amount,omitempty"`
	Currency               string `json:"currency,omitempty"`
	TransactionReference   string `json:"transactionReference,omitempty"`
}
