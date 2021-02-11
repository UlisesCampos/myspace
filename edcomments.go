package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Neil-uli/myspace/commons"
	"github.com/Neil-uli/myspace/migration"
	"github.com/Neil-uli/myspace/routes"
	"github.com/urfave/negroni"
)

func main() {

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}
	log.Printf("Iniciando el servidor en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
}
