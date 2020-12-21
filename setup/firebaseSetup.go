package setup

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// Creates firebase auth.Client for the given context and credentials file
func FirebaseAuthClient(ctx context.Context, credentialsFile string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}