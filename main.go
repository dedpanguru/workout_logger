package main

import (
	"fmt"
	"os"

	"github.com/dedpanguru/workout_logger/api"
	"github.com/dedpanguru/workout_logger/database"
)

//os.Getenv("DATABASE_URL")
const dbURL = "postgres://username:password@host.docker.internal:5432/workouts"

func main() {
	db, err := database.NewConn(dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	router := api.NewRouter(db)
	router.Logger.Fatal(router.Start(":8080"))
}
