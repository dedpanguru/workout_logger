package main

import (
	"fmt"
	"os"

	"workout-journal/api"
	"workout-journal/database"
)

//os.Getenv("DATABASE_URL")
const dbURL = "postgres://username:password@host.docker.internal:5432/workouts"

func main() {
	dbpool, ctx, err := database.NewConnPool(dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	router := api.NewRouter(dbpool, ctx)
	router.Logger.Fatal(router.Start(":8080"))
}
