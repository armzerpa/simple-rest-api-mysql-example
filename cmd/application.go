package cmd

import (
	"fmt"

	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize(config *config.Config) {
	fmt.Println("in application file")
}
