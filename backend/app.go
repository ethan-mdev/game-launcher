package backend

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/Microsoft/go-winio"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartGame(username, apiKey string) error {
	pipePath := `\\.\pipe\game_launcher`

	pipe, err := winio.ListenPipe(pipePath, nil)
	if err != nil {
		return fmt.Errorf("failed to create pipe: %w", err)
	}
	defer pipe.Close()

	// Launch the game
	go func() {
		cmd := exec.Command(".\\Game.bin", "-i", "10.0.0.97", "-p", "9010")
		cmd.Start()
	}()

	// Wait for client to connect
	conn, err := pipe.Accept()
	if err != nil {
		return fmt.Errorf("failed to accept pipe connection: %w", err)
	}
	defer conn.Close()

	// Write credentials
	data := fmt.Sprintf(`{"username":"%s","api_key":"%s"}`, username, apiKey)
	_, err = conn.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}

	return nil
}
