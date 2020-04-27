package main

import (
	"github.com/gorilla/websocket"
	"context"
	"net/http"
	"encoding/json"
	mon "part5/internal/monitor"
	"io/ioutil"
	"part5/internal/incident"
	"log"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

// We declare a websocket handler for streaming real-time incidents to any connected
// clients.
func serveMonitors(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()
	mon.Run(context.Background(), mon.Monitor{
		WSConn: conn,
	})
}


func saveIncomingIncidents(w http.ResponseWriter, r *http.Request) {
	var evt incident.Incident

	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &evt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := incident.Persist(&evt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mon.Broadcast(evt)

	storedPayload, _ := json.Marshal(evt)
	w.WriteHeader(http.StatusCreated)
	w.Write(storedPayload)
}
