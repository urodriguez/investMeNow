package main


import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Asset struct {
    Name      string `json:"name"`
    Description string `json:"description"`
    Image       string `json:"image"`
}

type Assets []Asset

func AllAssetsEndPoint(w http.ResponseWriter, r *http.Request) {	

	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	assets := Assets{
	  { Name: "Real State / Buildings", Description: "Real State / Buildings Real State / Buildings Real State / Buildings Real State / Buildings", Image:"https://images.pexels.com/photos/290275/pexels-photo-290275.jpeg?cs=srgb&dl=city-skyline-buildings-290275.jpg&fm=jpg" },
      { Name: "Copyrights", Description: "Copyrights Copyrights Copyrights Copyrights", Image:"https://sites.google.com/site/derechosdedeautorr/_/rsrc/1468755296557/para-que-sirve/11504987-vector-sello-danado-los-derechos-de-autor.jpg" },
      { Name: "Independent Project", Description: "Independent Project Independent Project Independent Project Independent Project", Image:"http://factorypyme.thestandardit.com/wp-content/uploads/sites/5/2014/09/Adquisicion.jpg" },
      { Name: "Next Sport Stars", Description: "Next Sport Stars Next Sport Stars Next Sport Stars Next Sport Stars", Image:"https://pbs.twimg.com/profile_images/875393235354439680/qWkpDQnh_400x400.jpg" }}

    json.NewEncoder(w).Encode(assets)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/assets", AllAssetsEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}