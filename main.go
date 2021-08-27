package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Roberto-GH/Operacion-Fuego-de-Quazar/controllers"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	File, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(File)

	router := mux.NewRouter()

	router.HandleFunc("/topsecret", controllers.AllSatellites).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", controllers.SatelliteSplit).Methods("POST")
	router.HandleFunc("/topsecret_split/", controllers.DecodeMessageSplit).Methods("GET")
	router.HandleFunc("/topsecret_split/", controllers.GetSatelliteSplit).Methods("POST")
	router.HandleFunc("/topsecret_split/", controllers.DeleteSatelliteSplit).Methods("DELETE")
	fmt.Println("Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
