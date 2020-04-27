package monitor

import (
	"log"
	"sync"
	"encoding/json"
	"part5/internal/incident"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Monitor struct {
	uuid string
	update chan incident.Incident
	WSConn *websocket.Conn
}

var (
	monitors  map[string]Monitor
	mtx *sync.Mutex
	initialiser sync.Once
)

func InitialiseRegistry() {
	initialiser.Do(func(){
		mtx = &sync.Mutex{}
		monitors = make(map[string]Monitor)
	})
}

func Run(ctx context.Context, mon Monitor) {
	mon.update = make(chan incident.Incident)
	mon.uuid = uuid.New().String()

	mtx.Lock()
	monitors[mon.uuid] = mon
	mtx.Unlock()

	defer func(){
		close(mon.update)
		delete(monitors, mon.uuid)
	}()

	for {
		select {
		case update := <-mon.Update:
			if err := send(mon, update); err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func Broadcast(incident incident.Incident) {
	for _, m := range monitors {
		m.update <- incident
	}
}

func send(mon Monitor, incident incident.Incident) error {
	payload, err := json.Marshal(incident)
	if err != nil {
		// Faulty payload, log and ignore this error: it doesn't reflect
		// the connection state of the monitor.
		log.Printf("unable to dispatch payload: %v", err)
		return nil
	}

	if err := mon.WSConn.WriteMessage(websocket.TextMessage, payload); err != nil {
		return err
	}

	return nil
}
