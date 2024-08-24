package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/shanth1/golang-microservices/authorization/sso/internal/domain/models"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appId int) (models.App, error)
}

// New returns a new instance of the Auth service
func New(
	log *slog.Logger,
	usrSaver UserSaver,
	usrProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    usrSaver,
		usrProvider: usrProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

// Login checks if user with given credentials exists in the system
//
// If user exists, but password is incorrect, returns error
// If user doesn't exists, returns error
func (a *Auth) Login(
	ctx context.Context,
	email string,
	psw string,
	appID int,
) (string, error) {
	panic("not implemented")
}

// RegisterNewUser registers new user in the system and returns id
func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	psw string,
) (int64, error) {
	panic("not implemented")
}

// IsAdmin checks if user is admin
func (a *Auth) IsAdmin(
	ctx context.Context,
	userId int,
) (bool, error) {
	panic("not implemented")
}
