package app

import (
	"fmt"

	"todocli/client"
	"todocli/server"
	"todocli/storage"
)

type App struct {
	server *server.Server
	client *client.Client
}

func ConfigureApp(storagePath string) (*App, error) {
	store, err := storage.New(storagePath)
	if err != nil {
		return nil, fmt.Errorf("error while configuring app: %w", err)
	}

	s, err := server.New(store)
	if err != nil {
		return nil, fmt.Errorf("error while configuring app: %w", err)
	}

	c, err := client.New(s)
	if err != nil {
		return nil, fmt.Errorf("error while configuring app: %w", err)
	}

	app := App{
		server: s,
		client: c,
	}

	return &app, nil
}

func (a *App) Run() error {
	for {
		a.client.HandleCommand()
	}
}
