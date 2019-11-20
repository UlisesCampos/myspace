package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/UlisesCampos/myspace/migration"
	"github.com/UlisesCampos/myspace/routes"
	"github.com/urfave/negroni"
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

	//Inicia las rutas
	router := routes.InitRoutes()

	//Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciando el servidor ")
	log.Println(server.ListenAndServe())
}
