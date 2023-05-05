package setup

import "github.com/gin-gonic/gin"

type appEngine struct {
	router *gin.Engine
}

func NewAppEngine() *appEngine {
	router := gin.Default()
	return &appEngine{router}
}

func (app *appEngine) InitializeRoutes() *appEngine {
	return app
}

func (app *appEngine) RunMigrations() *appEngine {
	return app
}

func (app *appEngine) Run(port string) *appEngine {
	app.router.Run(port)
	return app
}
