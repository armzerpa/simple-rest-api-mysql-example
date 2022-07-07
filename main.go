package main

import (
	"github.com/armzerpa/simple-rest-api-mysql-example/cmd"
	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/config"
)

func main() {
	config := config.GetConfig()

	app := &cmd.App{}
	app.Initialize(config)
}
