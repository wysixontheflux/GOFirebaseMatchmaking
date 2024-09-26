package matchmaking

import (
	"context"
	"firebase.google.com/go/v4/db"
	"goMatchmaking/models"
	"log"
)

func AddPlayerToQueue(queueRef *db.Ref, user models.User) {
	err := queueRef.Child(user.ID).Set(context.Background(), user)
	if err != nil {
		log.Fatalf("Error adding player to queue: %v\n", err)
	}
	log.Printf("Player added : %s\n", user.Name)
}
