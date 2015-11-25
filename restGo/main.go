package main

import (
	//"fmt"
	"github.com/codegangsta/negroni"
	//"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"github.com/nicolas.garaza/router"
)

func main() {
	
	r:= router.RegisterEnrollmentRoutes()
	
	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	// Set Up router
	n.UseHandler(r)

	n.Run(":8080")
}



func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back...")
}
