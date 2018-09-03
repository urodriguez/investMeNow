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
	  { Name: "Real State / Buildings", Description: "Real State / Buildings Real State / Buildings Real State / Buildings Real State / Buildings", Image:"https://www.apafacadesystems.com/wp-content/uploads/2017/12/theroy_newresized.jpg" },
      { Name: "Copyrights", Description: "Copyrights Copyrights Copyrights Copyrights", Image:"http://majesticproject.com/final-year-projects/images/content/feature/3.jpg" },
      { Name: "Independent Project", Description: "Independent Project Independent Project Independent Project Independent Project", Image:"http://www.gruposirem.mx/images/includes/projectimage5.jpg" },
      { Name: "Next Sport Stars", Description: "Next Sport Stars Next Sport Stars Next Sport Stars Next Sport Stars", Image:"https://i.pinimg.com/originals/ff/05/fa/ff05faceda3bbab1c53a0f7eed5a9a92.jpg" }}

    json.NewEncoder(w).Encode(assets)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/assets", AllAssetsEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}