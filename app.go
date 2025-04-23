package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
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
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) KillProcess(pid string) error {
	cmd := exec.Command("kill", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ListProcesses() ([]string, error) {
	cmd := exec.Command("ps", "-eo", "pid,user,%cpu,time,command,state")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	return lines, nil
}

func (a *App) PauseProcess(pid string) error {
	fmt.Println("Pausing process with PID:", pid)
	cmd := exec.Command("kill", "-STOP", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ResumeProcess(pid string) error {
	cmd := exec.Command("kill", "-CONT", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
