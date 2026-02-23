package main

import (
	"fmt"
	"log"

	"github.com/nanch/claude-task-app/backend/infrastructure/config"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := config.Migrate(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully")
}
