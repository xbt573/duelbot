package tdapi

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/gotd/contrib/bg"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/peers"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
)

type Client struct {
	Telegram *telegram.Client
	Peers    *peers.Manager
}

func New() (*Client, error) {
	id, _ := strconv.Atoi(os.Getenv("APP_ID"))
	if id <= 0 {
		return nil, errors.New("invalid APP_ID")
	}

	sessionFile := os.Getenv("SESSION_FILE")
	if sessionFile == "" {
		sessionFile = "session.json"
	}

	var h telegram.UpdateHandler

	client := telegram.NewClient(id, os.Getenv("APP_HASH"), telegram.Options{
		SessionStorage: &session.FileStorage{Path: sessionFile},
		UpdateHandler: telegram.UpdateHandlerFunc(
			func(ctx context.Context, u tg.UpdatesClass) error {
				if h != nil {
					return h.Handle(ctx, u)
				}
				return nil
			},
		),
	})

	if _, err := bg.Connect(client); err != nil {
		return nil, err
	}

	status, err := client.Auth().Status(context.Background())
	if err != nil {
		return nil, err
	}

	if !status.Authorized {
		return nil, errors.New("not authorized")
	}

	manager := peers.Options{}.Build(client.API())
	if err := manager.Init(context.Background()); err != nil {
		return nil, err
	}

	gaps := updates.New(updates.Config{
		Handler:      tg.NewUpdateDispatcher(),
		AccessHasher: manager,
	})
	h = manager.UpdateHook(gaps)

	return &Client{client, manager}, nil
}
