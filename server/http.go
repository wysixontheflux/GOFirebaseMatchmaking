package server

import (
	"encoding/json"
	"firebase.google.com/go/v4/db"
	"fmt"
	"goMatchmaking/matchmaking"
	"goMatchmaking/models"
	"log"
	"net/http"
)

func JoinQueueHandler(queueRef *db.Ref) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err = matchmaking.AddUserToQueue(queueRef, user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when adding the player to the queue: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User successfully added to the queue",
		})
	}
}

func StartHTTPServer(queueRef *db.Ref) {
	http.HandleFunc("/join-queue", JoinQueueHandler(queueRef))

	log.Println("Http server starting on :8282")
	if err := http.ListenAndServe(":8282", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur HTTP : %v\n", err)
	}
}
