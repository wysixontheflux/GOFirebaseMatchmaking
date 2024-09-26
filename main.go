package main

import (
	"goMatchmaking/firebase"
	"goMatchmaking/matchmaking"
	"goMatchmaking/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	log.Printf("FIREBASE_CREDENTIALS_PATH: %s\n", os.Getenv("FIREBASE_CREDENTIALS_PATH"))
	log.Printf("FIREBASE_DATABASE_URL: %s\n", os.Getenv("FIREBASE_DATABASE_URL"))

	client := firebase.InitializeFirebase()

	queueRef := client.NewRef("matchmakingQueue")
	roomsRef := client.NewRef("rooms")

	log.Println("Matchmaking system and servers started..")

	go server.StartHTTPServer(queueRef)

	//go server.StartWSServer()

	matchmaking.Matchmaking(queueRef, roomsRef)
}
