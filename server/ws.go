package server

/*


import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

/*func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Erreur WebSocket: %v\n", err)
	}
	defer ws.Close()

	clients[ws] = true

	for msg := range common.Broadcast {
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("Erreur d'envoi de message: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func StartWSServer() {
	http.HandleFunc("/ws", WsEndpoint)
	go func() {
		log.Println("Démarrage du serveur WebSocket sur le port 8081...")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Erreur lors du démarrage du serveur WebSocket: %v\n", err)
		}
	}()
}
*/
