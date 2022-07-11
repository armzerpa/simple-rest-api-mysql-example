package cmd

import (
	"database/sql"
	"log"

	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/api/handler"
	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/config"
	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/repository/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize(config *config.Config) {
	db := a.initDB(config.DB)
	bookHandler := handler.NewBookHandler(db)
	a.initRouter(config.Route, bookHandler)
}

func (a *App) initDB(config *config.DBConfig) *sql.DB {
	db, err := db.InitDatabase(*config)
	if err != nil {
		log.Println("Could not connect to database")
	}
	return db
}

func (a *App) initRouter(config *config.RouteConfig, handlerBook *handler.HandlerBook) {
	a.Router = gin.Default()
	version := a.Router.Group(config.Version)
	{
		version.GET("/ping", handler.Ping)
		version.GET("/books", handlerBook.GetBooks)
		version.GET("/books/:id", handlerBook.GetBookById)
	}
	a.Router.Run(config.Port)
}
