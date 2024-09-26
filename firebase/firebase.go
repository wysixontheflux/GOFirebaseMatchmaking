package firebase

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"fmt"
	"goMatchmaking/models"
	"google.golang.org/api/option"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func InitializeFirebase() *db.Client {
	client := initializeClient()
	initializeDatabase(client)
	return client
}

func initializeClient() *db.Client {
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	databaseURL := os.Getenv("FIREBASE_DATABASE_URL")

	log.Printf("Using Firebase Database URL: %s\n", databaseURL)

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v\n", err)
	}

	client, err := app.DatabaseWithURL(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Firebase Database Connection Error: %v\n", err)
	}

	return client
}

func initializeDatabase(client *db.Client) {
	queueRef := client.NewRef("matchmakingQueue")
	roomsRef := client.NewRef("rooms")

	log.Println("Resetting the matchmakingQueue and rooms...")

	err := queueRef.Delete(context.Background())
	if err != nil {
		log.Fatalf("Error resetting matchmakingQueue: %v\n", err)
	}

	err = roomsRef.Delete(context.Background())
	if err != nil {
		log.Fatalf("Error resetting rooms: %v\n", err)
	}

	go addFakePlayersPeriodically(queueRef)

	go func() {
		for {
			var rawData interface{}
			err = queueRef.Get(context.Background(), &rawData)
			if err != nil {
				log.Printf("Error retrieving matchmakingQueue content: %v\n", err)
			} else {
				if rawData == nil {
					log.Printf("Current MatchmakingQueue Content : []\n")
				} else {
					data, ok := rawData.(map[string]interface{})
					if !ok {
						log.Printf("Unexpected data type for matchmakingQueue: %T\n", rawData)
					} else {
						var queueData []models.User
						for _, value := range data {
							userData, err := json.Marshal(value)
							if err != nil {
								log.Printf("Error marshalling user: %v\n", err)
								continue
							}
							var user models.User
							err = json.Unmarshal(userData, &user)
							if err != nil {
								log.Printf("Error unmarshalling user: %v\n", err)
								continue
							}
							queueData = append(queueData, user)
						}

						// Sort the list by ID
						sort.Slice(queueData, func(i, j int) bool {
							idI, _ := strconv.Atoi(queueData[i].ID[1:])
							idJ, _ := strconv.Atoi(queueData[j].ID[1:])
							return idI < idJ
						})

						log.Printf("Current MatchmakingQueue Content : %v\n", queueData)
					}
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func addFakePlayersPeriodically(queueRef *db.Ref) {
	playerID := 1
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		newPlayer := models.User{
			ID:   fmt.Sprintf("u%d", playerID),
			Name: fmt.Sprintf("Player%d", playerID),
		}

		err := queueRef.Child(newPlayer.ID).Set(context.Background(), newPlayer)
		if err != nil {
			log.Printf("Error adding player %v : %v\n", newPlayer.Name, err)
		} else {
			log.Printf("Player added : %v\n", newPlayer.Name)
		}

		playerID++
	}
}
