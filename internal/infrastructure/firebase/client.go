package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirebaseApp(serviceAccountPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(serviceAccountPath)
	return firebase.NewApp(context.Background(), nil, opt)
}
