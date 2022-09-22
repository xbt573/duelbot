package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

func prompt(name string) (string, error) {
	fmt.Print("Enter " + name + ": ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

type authenticator struct{}

func (authenticator) Phone(_ context.Context) (string, error) {
	return prompt("phone")
}

func (authenticator) Password(_ context.Context) (string, error) {
	return prompt("password")
}

func (authenticator) Code(
	ctx context.Context,
	sentCode *tg.AuthSentCode,
) (string, error) {
	return prompt("code")
}

func (authenticator) AcceptTermsOfService(
	ctx context.Context,
	tos tg.HelpTermsOfService,
) error {
	return nil
}

func (authenticator) SignUp(_ context.Context) (auth.UserInfo, error) {
	panic("signup is not supported")
}

func main() {
	id, _ := strconv.Atoi(os.Getenv("APP_ID"))
	if id <= 0 {
		fmt.Println("invalid APP_ID")
		os.Exit(1)
	}

	sessionFile := os.Getenv("SESSION_FILE")
	if sessionFile == "" {
		sessionFile = "session.json"
	}

	client := telegram.NewClient(id, os.Getenv("APP_HASH"), telegram.Options{
		SessionStorage: &session.FileStorage{Path: sessionFile},
	})

	if err := client.Run(context.Background(), func(ctx context.Context) error {
		status, err := client.Auth().Status(ctx)
		if err != nil {
			return err
		}

		if status.Authorized {
			return errors.New("already authorized")
		}

		return auth.NewFlow(
			authenticator{},
			auth.SendCodeOptions{},
		).Run(ctx, client.Auth())
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
