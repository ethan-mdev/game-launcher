package backend

import (
	"context"
	"os"
	"os/exec"
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
func (a *App) StartGame(username, password string) {
	exec.Command(".\\LauncherUpdater.exe", username, password).Output()
	os.Exit(0)
}
