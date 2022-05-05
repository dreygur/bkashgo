package methods

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/dreygur/bkashgo/models"
)

func getMessageBytesToSign(msg *models.BkashIPNPayload) []byte {
	var builtSignature bytes.Buffer
	signableKeys := []string{"Message", "MessageId", "Subject", "SubscribeURL", "Timestamp", "Token", "TopicArn", "Type"}
	for _, key := range signableKeys {
		reflectedStruct := reflect.ValueOf(msg)
		field := reflect.Indirect(reflectedStruct).FieldByName(key)
		value := field.String()
		if field.IsValid() && value != "" {
			builtSignature.WriteString(key + "\n")
			builtSignature.WriteString(value + "\n")
		}
	}
	return builtSignature.Bytes()
}

func IsMessageSignatureValid(msg *models.BkashIPNPayload) error {
	resp, err := http.Get(msg.SigningCertURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("unable to get certificate err: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	p, _ := pem.Decode(body)
	cert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return err
	}

	base64DecodedSignature, err := base64.StdEncoding.DecodeString(msg.Signature)
	if err != nil {
		return err
	}

	if err := cert.CheckSignature(x509.SHA1WithRSA, getMessageBytesToSign(msg), base64DecodedSignature); err != nil {
		return err
	}

	return nil
}
