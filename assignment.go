package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Config struct {
	Name string `json:"name,omitempty"`
	Data `json:"data,omitempty"`
}

type Data struct {
	Key string `json:"key,omitempty"`
	Id  int    `json:"id,omitempty"`
}

var configs []Config

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Application is up and running")
}

func GetConfigs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(configs)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, config := range configs {
		if config.Name == params["name"] {
			json.NewEncoder(w).Encode(config)
			return
		}
	}
	json.NewEncoder(w).Encode(&Config{})
}

func CreateConfig(w http.ResponseWriter, r *http.Request) {
	var config Config
	_ = json.NewDecoder(r.Body).Decode(&config)
	configs = append(configs, config)
	json.NewEncoder(w).Encode(config)
}

func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var config Config
	_ = json.NewDecoder(r.Body).Decode(&config)
	configs = append(configs, config)
	json.NewEncoder(w).Encode(config)
}

func DeleteConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, config := range configs {
		if config.Name == params["name"] {
			configs = append(configs[:index], configs[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(configs)
	}
}

func SearchConfig(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	n := r.Form.Get("name")

	d := r.Form.Get("data.id")
	data, _ := strconv.Atoi(d)

	for _, config := range configs {
		if config.Name == n && config.Id == data {
			json.NewEncoder(w).Encode(config)
			return
		}
	}
	json.NewEncoder(w).Encode(&Config{})
}

func main() {
	var PORT string
	if PORT = os.Getenv("SERVE_PORT"); PORT == "" {
		PORT = "3000"
	}
	router := mux.NewRouter()
	configs = append(configs, Config{"config1", Data{"val1", 1}})
	configs = append(configs, Config{"config2", Data{"val2", 2}})

	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	router.HandleFunc("/configs", GetConfigs).Methods("GET")
	router.HandleFunc("/configs/{name}", GetConfig).Methods("GET")
	router.HandleFunc("/configs", CreateConfig).Methods("POST")
	router.HandleFunc("/configs/{name}", DeleteConfig).Methods("DELETE")
	router.HandleFunc("/search", SearchConfig).Methods("GET")
	router.HandleFunc("/configs/{name}", UpdateConfig).Methods("PUT")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
