package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Car struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	IsElectric bool `json:"isElectric"`
	HorsePower int64 `json:"horsePower"`
	BasePrice float64 `json:"basePrice"`
	StorePrice float64 `json:"storePrice"`
}

var cars []Car

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if len(cars) < 1 {
		http.Error(w, "No car is available", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func getCarById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if len(cars) < 1 {
		http.Error(w, "No car is available", http.StatusNotFound)
		return
	}

	params := mux.Vars(r)

	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}


	http.Error(w, "No car with the given ID was found", http.StatusNotFound)
}

func calculateStorePrice(input *Car) {
	if input.IsElectric {
		input.StorePrice = input.BasePrice + (float64(input.HorsePower) / 4)
		return
	}

	input.StorePrice = input.BasePrice + (float64(input.HorsePower) / 2) + 100
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCar Car
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, "Error createCar input", http.StatusInternalServerError)
		return
	}

	newCar.ID = uuid.NewString()

	calculateStorePrice(&newCar)

	cars = append(cars, newCar)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func updateCarById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCar Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Error updatedCar input", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			// delete old data
			cars = append(cars[:index], cars[index+1:]...)

			// insert new data but with same carID
			updatedCar.ID = item.ID
			calculateStorePrice(&updatedCar)
			cars = append(cars, updatedCar)

			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}

	http.Error(w, "No car with the given ID was found", http.StatusNotFound)
}

func deleteCarById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			// delete old data
			cars = append(cars[:index], cars[index+1:]...)

			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "No car with the given ID was found", http.StatusNotFound)
}

func main() {
	// init router
	router := mux.NewRouter()

	// route endpoints
	router.HandleFunc("/api/v1/cars", getCars).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cars/{id}", getCarById).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cars", createCar).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/cars/{id}", updateCarById).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/cars/{id}", deleteCarById).Methods(http.MethodDelete)


	var port string

	if os.Getenv("PORT") != "" {
		port =  os.Getenv("PORT")
	} else {
		port = "8001"
	}

	fmt.Println("TechnoCar REST API is running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}