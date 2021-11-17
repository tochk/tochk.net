package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./migrate [up|down] [dir]")
	}

	m, err := migrate.New("file://"+os.Args[2], os.Getenv("TOCHKNET_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "up":
		log.Println("apply all migrations")
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		log.Println("revert one migration")
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("completed")
}
