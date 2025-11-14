package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Location    string  `json:"location"`
	SensorID    string  `json:"sensor_id"`
	Temperature float64 `json:"temperature"`
}

func locationBySensorId(sensorID string) string {
	switch sensorID {
	case "1":
		return "Living Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}
}

func sensorIdByLocation(location string) string {
	switch location {
	case "Living Room":
		return "1"
	case "Bedroom":
		return "2"
	case "Kitchen":
		return "3"
	default:
		return "0"
	}
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	if location == "" && sensorID != "" {
		location = locationBySensorId(sensorID)
	}

	if sensorID == "" && location != "" {
		sensorID = sensorIdByLocation(location)
	}

	if location == "" && sensorID == "" {
		location = "Unknown"
		sensorID = "0"
	}

	rand.Seed(time.Now().UnixNano())
	temperature := 16 + rand.Float64()*(30-16)

	resp := Response{
		Location:    location,
		SensorID:    sensorID,
		Temperature: float64(int(temperature*10)) / 10, // округление до 1 знака
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/temperature", temperatureHandler)
	http.ListenAndServe(":8081", nil)
}
