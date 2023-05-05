package main

import (
	"hubla/desafiofullstack/setup"
	"hubla/desafiofullstack/utils"
)

func main() {
	db := utils.GetDB()
	defer db.DB()

	setup.NewAppEngine().
		InitializeRoutes().
		RunMigrations(db).
		Run(":8080")

}
