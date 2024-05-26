package main

import (
	"context"
	"log"
	"projectservice/app"
	"projectservice/database"

	"github.com/nats-io/nats.go"
)

func main() {
	ctx := context.Background()

	// Initialize the database if needed
	if err := database.Initialize(); err != nil {
		log.Fatal("Failed to initialize the database: ", err)
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Fatal("Failed to close the database: ", err)
		}
	}()

	// Run migrations
	if err := database.RunMigrations(ctx); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Register NATS subscribers for each project operation
	nc.Subscribe("CreateProject", app.CreateProjectHandler())
	nc.Subscribe("UpdateProject", app.UpdateProjectHandler())
	nc.Subscribe("ReadProject", app.ReadProjectHandler())
	nc.Subscribe("ReadAllProjects", app.ReadAllProjectsHandler())
	nc.Subscribe("DeleteProject", app.DeleteProjectHandler())

	// Keep the connection alive
	select {}
}
