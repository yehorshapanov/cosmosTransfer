package app

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"strings"
)

func CalculateTransactionID(txs string) string {
	sDec, _ := b64.StdEncoding.DecodeString(txs)

	h := sha256.New()
	h.Write(sDec)
	res := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))

	return res
}
