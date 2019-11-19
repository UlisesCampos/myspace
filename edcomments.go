package main

import (
	"flag"
	"log"

	"github.com/UlisesCampos/myspace/migration"
)

func main() {

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migracion..")
		migration.Migrate()
		log.Println("Finalizo la migracion")
	}
}
