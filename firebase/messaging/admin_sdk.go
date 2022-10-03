package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	// option "google.golang.org/api/option"
)

func NewFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	// app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func sendToToken(app *firebase.App) {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "YOUR_REGISTRATION_TOKEN"

	// See documentation on defining a message payload.
	androidNotification := &messaging.AndroidNotification{
		Title:    "Title from golang",
		Body:     "body body body: heyheyheyhey",
		ImageURL: "https://pekepeke",
		Priority: messaging.PriorityHigh,
	}
	config := &messaging.AndroidConfig{
		Notification: androidNotification,
	}
	notification := &messaging.Notification{
		Title:    "Title from golang",
		Body:     "Body Of Knowledge",
		ImageURL: "https://pekepeke",
	}
	apnsConfig := &messaging.APNSConfig{
		Headers: map[string]string{
			"apns-priority": "10",
		},
		Payload: &messaging.APNSPayload{
			Aps: &messaging.Aps{
				Alert: &messaging.ApsAlert{
					Title: "APNs Title",
					Body:  "Body from golang",
				},
			},
		},
	}
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Android:      config,
		APNS:         apnsConfig,
		Notification: notification,
		Token:        registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
