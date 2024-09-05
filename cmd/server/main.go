package main

import (
	"github.com/zhandos717/go-http-api/pkg/systems"
)

func main() {

	token := systems.BotToken()

	println(token)
}
