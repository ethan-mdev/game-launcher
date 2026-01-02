package backend

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/Microsoft/go-winio"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// StartGame launches the game with the provided credentials
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

	runtime.Quit(a.ctx)

	return nil
}
