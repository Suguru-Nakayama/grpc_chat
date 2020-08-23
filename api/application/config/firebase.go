package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

/*
 * Firebase Authentication CLientオブジェクトを取得
 */
func NewFirebaseAuthClient() (*auth.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SECRET"))
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v\n", err)
	}

	return client, nil
}
