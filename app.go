package main

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

const appName string = "p2p-chat"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) DoesAppHaveDir() (bool, error) {
	appDir := filepath.Join(xdg.DataHome, appName)
	_, err := os.Stat(appDir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (a *App) ReadAppFiles() string {
	log.Println("Home data directory:", xdg.DataHome)
	return xdg.DataHome
}

func (a *App) ReadAppConfig() string {
	log.Println("Config data directory:", xdg.ConfigHome)
	return xdg.ConfigHome
}

type User struct {
	Username string `json:"username"`
}

func (a *App) ReadUsers() ([]User, error) {
	err := errors.New("no users")
	return nil, err
}
