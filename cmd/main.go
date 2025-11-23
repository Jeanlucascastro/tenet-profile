package main

import "tenet-profile/config"

func main() {

	db, err := config.InitDataBase()
	if err != nil {
		panic("Failed to connect to database")
	}

	config.RunMigrations(db)
}
