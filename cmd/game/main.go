package main

import (
	"go-m2-server/internal/network/auth"
	"go-m2-server/internal/network/game"
)

func main() {
	// Start Auth server
	go auth.StartServer(23070, true)

	// Start Game server
	game.StartServer(23000, true)

}
