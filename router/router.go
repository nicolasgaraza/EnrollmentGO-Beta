package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterEnrollmentRoutes() (*httprouter.Router) {
	r:= httprouter.New()
	
	myRouterConfig(r)
	
	return r;
}

func myRouterConfig(router *httprouter.Router) {
	router.GET("/pepe", func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "pepe")
	})

	router.GET("/", homeHandler)

}

func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home Nico")
}

