package setup

import (
	"hubla/desafiofullstack/controllers"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appEngine struct {
	Router *gin.Engine
}

func NewAppEngine() *appEngine {
	Router := gin.Default()
	Router.MaxMultipartMemory = 8 << 20 // setting a max size at file, 8Mib

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	authRoute := cors.New(corsConfig)

	Router.Use(authRoute)
	return &appEngine{Router}
}

func (app *appEngine) InitializeRoutes() *appEngine {

	app.Router.POST("signup", controllers.CreateUser)
	app.Router.POST("signin", controllers.ValidateLogin)

	app.Router.POST("creator", utils.IsAuthorized, controllers.CreateNewCreator)
	app.Router.POST("creator/:id/product", utils.IsAuthorized, controllers.CreatorAddProduct)
	app.Router.POST("creator/:id/afiliate", utils.IsAuthorized, controllers.CreatorAddAfiliate)

	app.Router.POST("upload", controllers.NormalizeData)
	return app
}

func (app *appEngine) RunMigrations(db *gorm.DB) *appEngine {
	err := db.AutoMigrate(&entitys.Login{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&entitys.User{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&entitys.Creator{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&entitys.Afiliated{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&entitys.Product{})
	if err != nil {
		panic(err.Error())
	}

	return app
}

func (app *appEngine) Run(port string) *appEngine {
	app.Router.Run(port)
	return app
}
