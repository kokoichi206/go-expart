package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

var key = []byte{}

func initKey() {
	for i := 0; i < 64; i++ {
		key = append(key, byte(i))
	}
}

func signMessage(msg []byte) ([]byte, error) {
	// https://pkg.go.dev/crypto/hmac#New
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to write: %w", err)
	}

	signature := h.Sum(nil)

	return signature, nil
}

func checkSignature(msg []byte, signature []byte) error {
	newSignature, err := signMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to sign: %w", err)
	}

	if !hmac.Equal(newSignature, signature) {
		return fmt.Errorf("signature is not equal")
	}

	return nil
}
