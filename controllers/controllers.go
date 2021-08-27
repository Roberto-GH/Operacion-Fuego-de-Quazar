package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Roberto-GH/Operacion-Fuego-de-Quazar/models"
	"github.com/Roberto-GH/Operacion-Fuego-de-Quazar/services"
	"github.com/gorilla/mux"
)

func AllSatellites(w http.ResponseWriter, r *http.Request) {
	var topSecret models.TopSecret
	_ = json.NewDecoder(r.Body).Decode(&topSecret)

	if len(topSecret.Satellites) == 3 {

		message := services.GetMessage(topSecret.Satellites[0].Message, topSecret.Satellites[1].Message, topSecret.Satellites[2].Message)
		x, y := services.GetLocation(topSecret.Satellites[0].Distance, topSecret.Satellites[1].Distance, topSecret.Satellites[2].Distance)

		var responseTopSecret models.ResponseTopSecret
		responseTopSecret.Message = message
		responseTopSecret.Position.X = x
		responseTopSecret.Position.Y = y

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseTopSecret)
		log.Printf("Respuesta all satellites: %v ", responseTopSecret)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func SatelliteSplit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var newSatellite models.Satellite
	newSatellite.Name = params["satellite_name"]

	_ = json.NewDecoder(r.Body).Decode(&newSatellite)
	newSatellite.Name = params["satellite_name"]

	if len(models.SatellitesBD) >= 3 {
		responseErrorSplit := models.ResponseSplit{
			"error": "Cantidad de satelites ingresados superada",
			"info":  "Usar el metodo delete para eliminar satelites ingresados /topsecret_split/",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseErrorSplit)
		log.Printf("Respuesta satelites split: %v ", responseErrorSplit)
	} else {
		models.SatellitesBD = append(models.SatellitesBD, newSatellite)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newSatellite)
		log.Printf("Respuesta satelites split: %v ", newSatellite)
	}

}

func GetSatelliteSplit(w http.ResponseWriter, r *http.Request) {

	if len(models.SatellitesBD) == 0 {
		responseErrorSplit := models.ResponseSplit{
			"status": "200",
			"body":   "Aun no has ingresado satelites",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseErrorSplit)
		log.Printf("Respuesta satelites get split: %v ", responseErrorSplit)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.SatellitesBD)
		log.Printf("Respuesta satelites get split: %v ", models.SatellitesBD)
	}

}

func DecodeMessageSplit(w http.ResponseWriter, r *http.Request) {

	if len(models.SatellitesBD) == 3 {

		message := services.GetMessage(models.SatellitesBD[0].Message, models.SatellitesBD[1].Message, models.SatellitesBD[2].Message)
		x, y := services.GetLocation(models.SatellitesBD[0].Distance, models.SatellitesBD[1].Distance, models.SatellitesBD[2].Distance)

		var responseTopSecret models.ResponseTopSecret
		responseTopSecret.Message = message
		responseTopSecret.Position.X = x
		responseTopSecret.Position.Y = y

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseTopSecret)
		log.Printf("Respuesta mensaje codificado: %v ", responseTopSecret)
	} else {
		var responseError models.ResponseError
		responseError.Er = "no hay suficiente información"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseError)
		log.Printf("Respuesta mensaje codificado: %v ", responseError)
	}
}

func DeleteSatelliteSplit(w http.ResponseWriter, r *http.Request) {
	models.SatellitesBD = nil
	responseErrorSplit := models.ResponseSplit{
		"status": "200",
		"body":   "Satelites eliminados satisfactoriamente",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseErrorSplit)
	log.Printf("Respuesta eliminar satelites: %v ", responseErrorSplit)
}
