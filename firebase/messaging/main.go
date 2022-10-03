package main

import (
	"fmt"
)

// https://github.com/firebase/firebase-admin-go/blob/bb055ed1cfbe6224367c63caedc4ba72f1437dcd/snippets/messaging.go
func main() {
	app, err := NewFirebaseApp()
	if err != nil {
		fmt.Printf("Failed to %s\n", "NewFirebaseApp")
		return
	}
	sendToToken(app)
}
