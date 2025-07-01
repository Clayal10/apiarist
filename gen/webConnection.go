package gen

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var webSocketData []byte

// This function will take a particle, run the network on the interval, and send
// it to the JS.
func visualDisplay(d []byte) {
	webSocketData = d

	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Server started on :8201")
	err := http.ListenAndServe(":8201", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	var ws *websocket.Conn
	var err error

	if ws, err = upgrader.Upgrade(w, r, nil); err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// Check for connection and send response.
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s", message)
		// Process message and send canvas commands
		err = ws.WriteMessage(websocket.BinaryMessage, webSocketData)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
