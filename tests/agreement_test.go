package tests

import (
	"testing"

	"github.com/dreygur/bkashgo"
	"github.com/dreygur/bkashgo/models"
)

var (
	createAgreementResponse *models.CreateAgreementResponse
	executeAgreementReponse *models.ExecuteAgreementResponse
	// createPaymentResponse   *models.CreatePaymentResponse
)

func TestAgreement(t *testing.T) {
	bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

	token, err := bkash.GrantToken()
	if err != nil {
		t.Fatal(err)
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		t.Fatal(err)
	}

	t.Run("test CreateAgreement", func(t *testing.T) {
		req := &models.CreateRequest{
			Mode:           "0000",
			PayerReference: "dsfsodjf-w3y2sdjf83493-sdhfis",
			CallbackURL:    "https://api.shikho.net/payment",
			Currency:       "BDT",
			Intent:         "Shikho Subscription",
		}
		resp, err := bkash.CreateAgreement(req, token)

		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}

		if resp == nil || resp.StatusCode != "0000" {
			t.Fatal("Invalid create agreement response")
		}

		createAgreementResponse = resp
	})

	t.Run("test ExecuteAgreement", func(t *testing.T) {
		req := &models.ExecuteRequest{
			PaymentID: createAgreementResponse.PaymentID,
		}
		resp, err := bkash.ExecuteAgreement(req, token)

		if err != nil {
			t.Error(err.Error())
			t.Fail()
		}

		if resp == nil || resp.StatusCode != "0000" {
			t.Fatal("Invalid execute agreement response")
		}

		executeAgreementReponse = resp
	})

	// t.Run("test CreatePayment", func(t *testing.T) {
	// 	req := &models.CreateRequest{
	// 		Mode:                    "0001",
	// 		PayerReference:          "01723888888",
	// 		CallbackURL:             "https://shikho.tech/payment",
	// 		AgreementID:             executeAgreementReponse.AgreementID,
	// 		Amount:                  "12",
	// 		Currency:                "BDT",
	// 		Intent:                  "sale",
	// 		MerchantInvoiceNumber:   "Inv0124",
	// 		MerchantAssociationInfo: "MI05MID54RF09123456One",
	// 	}
	// 	resp, err := bkash.CreatePayment(req, token, false)

	// 	if err != nil {
	// 		t.Error(err.Error())
	// 		t.Fail()
	// 	}

	// 	if resp == nil || resp.StatusCode != "0000" {
	// 		t.Fatal("payment creattion failed")
	// 	}

	// 	createPaymentResponse = resp
	// })

	// t.Run("test executePayment", func(t *testing.T) {
	// 	req := &models.ExecutePaymentRequest{
	// 		PaymentID: createPaymentResponse.PaymentID,
	// 	}
	// 	resp, err := paymentService.ExecutePayment(req, token, false)

	// 	if err != nil {
	// 		t.Error(err.Error())
	// 		t.Fail()
	// 	}

	// 	if resp == nil || resp.StatusCode != "0000" {
	// 		t.Fatal("payment creattion failed")
	// 	}
	// })
}
