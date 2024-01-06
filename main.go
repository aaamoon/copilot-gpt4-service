package main

import (
	"copilot-gpt4-service/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080")
}
