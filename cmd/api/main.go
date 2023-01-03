package main

import (
	"fmt"

	transportHttp "github.com/bajalnyt/chi-router-playground/internal/transport/http"
)

// Run instantiates and starts the app
func Run() error {
	fmt.Println("Starting up application")

	httpHandler := transportHttp.NewHandler()
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
