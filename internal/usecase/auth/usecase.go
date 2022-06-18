package auth

import "context"

type Usecase interface {
	Login(ctx context.Context, username string, password string) (accessToken string, err error)
}
