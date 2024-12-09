package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func VerifyWebhookSignature(signature, payload, secret string) error {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedMAC := mac.Sum(nil)
	expectedSignature := hex.EncodeToString(expectedMAC)

	if signature != expectedSignature {
		return errors.New("invalid webhook signature")
	}
	return nil
}
