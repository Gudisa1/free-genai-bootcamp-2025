package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	_ "github.com/mattn/go-sqlite3"
	"lang-portal/internal/models"
)

// Default target to run when none is specified
var Default = Run

// Install installs project dependencies
func Install() error {
	fmt.Println("Installing dependencies...")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return nil
}

// Run starts the development server
func Run() error {
	mg.Deps(Install)
	fmt.Println("Starting server...")
	cmd := exec.Command("go", "run", "cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Build builds the application
func Build() error {
	mg.Deps(Install)
	fmt.Println("Building...")
	return sh.Run("go", "build", "-o", "bin/server", "./cmd")
}

// InitDB initializes the database and runs migrations
func InitDB() error {
	fmt.Println("Initializing database...")

	// Initialize database
	db, err := models.NewDB("words.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}