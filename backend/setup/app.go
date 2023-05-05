package setup

import (
	"hubla/desafiofullstack/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func (app *appEngine) RunMigrations(db *gorm.DB) *appEngine {
	err := db.AutoMigrate(&models.Login{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}

	return app
}

func (app *appEngine) Run(port string) *appEngine {
	app.router.Run(port)
	return app
}
