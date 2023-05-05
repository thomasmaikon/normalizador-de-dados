package setup

import (
	"hubla/desafiofullstack/controllers"
	"hubla/desafiofullstack/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appEngine struct {
	Router *gin.Engine
}

func NewAppEngine() *appEngine {
	Router := gin.Default()
	return &appEngine{Router}
}

func (app *appEngine) InitializeRoutes() *appEngine {

	app.Router.POST("/user", controllers.CreateUser)

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
	app.Router.Run(port)
	return app
}
