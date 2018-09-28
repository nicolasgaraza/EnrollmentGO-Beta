package router

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nicolas.garaza/models/enrollment"
	"net/http"
	"github.com/nicolas.garaza/controllers/zipCode"
	zipCodeModel "github.com/nicolas.garaza/models/zipCode"
	//"io/ioutil"
)

func RegisterEnrollmentRoutes() *httprouter.Router {
	r := httprouter.New()

	myRouterConfig(r)

	return r
}

func myRouterConfig(router *httprouter.Router) {
	router.GET("/", homeHandler)

	router.GET("/enrollment", getEnrollments)
	
	router.GET("/zipCode/:code", func (w http.ResponseWriter, req *http.Request, ps httprouter.Params){
		fmt.Printf("%+v\n", ps.ByName("code"))
	})
	
	router.POST("/zipCode", func (w http.ResponseWriter, req *http.Request, ps httprouter.Params){
	
		controller := zipCode.New()	
	
		z := zipCodeModel.ZipCode{}
		
		json.NewDecoder(req.Body).Decode(&z)
		
		fmt.Printf("%+v\n", z)
		controller.Post(z);
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		
				
	})
	

	router.GET("/zipCode", func (w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	
		
		controller := zipCode.New()	
	
		z, er:= controller.Get()
		fmt.Printf("%+v\n", z)
		if er == nil{
			ej, _ := json.Marshal(z)
	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", ej)
			return
			
		}else {
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", er)
			return
		}
		
		ej, _ := json.Marshal(z)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", ej)
				
	})

}



func getEnrollments(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var e = enrollment.Enrollment{
		//Id: 123,
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
