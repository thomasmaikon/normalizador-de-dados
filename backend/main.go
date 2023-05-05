package main

import (
	"hubla/desafiofullstack/setup"
)

func main() {
	setup.NewAppEngine().
		InitializeRoutes().
		RunMigrations().
		Run(":8080")

}
