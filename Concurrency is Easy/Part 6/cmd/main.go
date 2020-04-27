package main

import (
	"part5/internal/monitor"
	"net/http"
	"log"
)

func main(){
	monitor.InitialiseRegistry()

	http.HandleFunc("/api/incident", saveIncomingIncidents)
	http.HandleFunc("/ws/monitor", serveMonitors)
	log.Fatal(http.ListenAndServe(":8080", nil))
}