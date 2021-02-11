package routes

import (
	"github.com/Neil-uli/myspace/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetCommentRouter creando la ruta parala cracion de comentarios
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")
	subRouter.HandleFunc("/", controllers.CommentGetAll).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
