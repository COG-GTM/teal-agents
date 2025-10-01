package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/COG-GTM/teal-agents/pkg/config"
)

func main() {
	cfg := config.GetInstance()

	fmt.Println("Teal Agents Go Server")
	fmt.Println("Configuration loaded from:", cfg.Get("TA_SERVICE_CONFIG"))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Server initialization will be implemented in future sessions")
	fmt.Println("Press Ctrl+C to exit")
	<-sigChan
	fmt.Println("\nShutting down gracefully...")
}
