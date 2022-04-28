package main

import (
	"github.com/KoLLlaka/_learning/session/internal/model"
	"github.com/KoLLlaka/_learning/session/internal/routes"
)

func main() {
	model.Setup()

	routes.Setup()
}
