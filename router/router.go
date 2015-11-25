package router

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nicolas.garaza/models/enrollment"
	"net/http"
)

func RegisterEnrollmentRoutes() *httprouter.Router {
	r := httprouter.New()

	myRouterConfig(r)

	return r
}

func myRouterConfig(router *httprouter.Router) {
	router.GET("/", homeHandler)

	router.GET("/enrollment", getEnrollments)

}



func getEnrollments(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var e = enrollment.Enrollment{
		Id: 123,
	}
	
	array:= []enrollment.Enrollment{
		e,
	}
	
	ej, _ := json.Marshal(array)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", ej)
}

func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home Nico")
}
