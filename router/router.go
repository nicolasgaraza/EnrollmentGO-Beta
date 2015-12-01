package router

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nicolas.garaza/models/enrollment"
	"net/http"
	"github.com/nicolas.garaza/controllers/zipCode"
)

func RegisterEnrollmentRoutes() *httprouter.Router {
	r := httprouter.New()

	myRouterConfig(r)

	return r
}

func myRouterConfig(router *httprouter.Router) {
	router.GET("/", homeHandler)

	router.GET("/enrollment", getEnrollments)

	router.GET("/zipCode", func (w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	
		
		controller := zipCode.New()	
	
		z, er:= controller.Get()
		
		if er != nil{
			ej, _ := json.Marshal(z)
	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", ej)
			
		}else {
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", er)
		}
		
		ej, _ := json.Marshal(z)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", ej)
				
	})

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
