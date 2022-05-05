package main

import (
	"log"

	"github.com/dreygur/bkashgo"
	"github.com/dreygur/bkashgo/models"
)

var (
	username    = "sandboxTokenizedUser02"
	password    = "sandboxTokenizedUser02@12345"
	appKey      = "4f6o0cjiki2rfm34kfdadl1eqq"
	appSecret   = "2is7hdktrekvrbljjh44ll3d9l1dtjo4pasmjvs5vl5qr3fug4b"
	isLiveStore = false
)

var (
	createAgreementResponse *models.CreateAgreementResponse
	executeAgreementReponse *models.ExecuteAgreementResponse
	createPaymentResponse   *models.CreatePaymentResponse
)

func main() {
	// bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)
	// token, err := bkash.GrantToken()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(token)
	TestAgreement()
}

func TestAgreement() {
	bkash := bkashgo.GetBkash(username, password, appKey, appSecret, isLiveStore)

	token, err := bkash.GrantToken()
	if err != nil {
		log.Fatal(err)
	}

	if token == nil || len(token.IdToken) == 0 || len(token.RefreshToken) == 0 || token.StatusCode != "0000" {
		log.Fatal(err)
	}

	func() {
		req := &models.CreateRequest{
			Mode:           "0000",
			PayerReference: "dsfsodjf-w3y2sdjf83493-sdhfis",
			CallbackURL:    "https://api.shikho.net/payment",
			Currency:       "BDT",
			Intent:         "Shikho Subscription",
		}
		resp, err := bkash.CreateAgreement(req, token)

		if err != nil {
			log.Fatal(err.Error())
		}

		if resp == nil || resp.StatusCode != "0000" {
			log.Fatal("Invalid create agreement response")
		}

		createAgreementResponse = resp
	}()

	func() {
		req := &models.ExecuteRequest{
			PaymentID: createAgreementResponse.PaymentID,
		}
		resp, err := bkash.ExecuteAgreement(req, token)

		if err != nil {
			log.Fatal(err.Error())
		}

		if resp == nil || resp.StatusCode != "0000" {
			log.Print(req)

			log.Print(resp)
			log.Fatal("Invalid execute agreement response")
		}

		executeAgreementReponse = resp
	}()

	func() {
		req := &models.CreateRequest{
			Mode:                    "0001",
			PayerReference:          "01723888888",
			CallbackURL:             "https://shikho.tech/payment",
			AgreementID:             executeAgreementReponse.AgreementID,
			Amount:                  "12",
			Currency:                "BDT",
			Intent:                  "sale",
			MerchantInvoiceNumber:   "Inv0124",
			MerchantAssociationInfo: "MI05MID54RF09123456One",
		}
		resp, err := bkash.CreatePayment(req, token)

		if err != nil {
			log.Fatal(err.Error())
		}

		if resp == nil || resp.StatusCode != "0000" {
			log.Fatal("payment creattion failed")
		}

		createPaymentResponse = resp
	}()

	func() {
		req := &models.ExecuteRequest{
			PaymentID: createPaymentResponse.PaymentID,
		}
		resp, err := bkash.ExecutePayment(req, token)

		if err != nil {
			log.Fatal(err.Error())
		}

		if resp == nil || resp.StatusCode != "0000" {
			log.Fatal("payment creattion failed")
		}
	}()
}
