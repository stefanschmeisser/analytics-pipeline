package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// GetAnalyticsStream is the handler function to receive the analytics result
// data from Apache Spark
func GetAnalyticsStream(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			return
		}
		log.Println(string(message))

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			return
		}

	}

}
