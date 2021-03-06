package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyusgupta/pgtuner/backend/api"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func main() {
	flag.Parse()
	api.HealthCheckInit()
	router := mux.NewRouter()
	router.HandleFunc("/", api.HealthCheckHandler)
	router.HandleFunc("/db/settings/", api.PGSettingHandler)
	router.HandleFunc("/db/tables/", api.PGTableHandler)
	// URLs for stats, can be used for charting
	router.HandleFunc("/stats/{clientID:[0-9]+}/{chartName}/", api.GetResourceMetricsHandler)
	log.Println("starting server  at port", *addr, "...")
	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
