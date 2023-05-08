package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func bcryptSample() {
	fmt.Println("========== bcryptSample ==========")
	base64.StdEncoding.EncodeToString([]byte("user:password"))

	pass := "test password"

	hashPassword, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashPassword)
	if err != nil {
		log.Fatal(err)
	}

	err = comparePassword("pass", hashPassword)
	fmt.Printf("err: %v\n", err)
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generate password: %w", err)
	}

	return bytes, nil
}

func comparePassword(password string, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return fmt.Errorf("failed to compare password: %w", err)
	}

	return nil
}
