package matchmaking

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/db"
	"fmt"
	"goMatchmaking/common"
	"goMatchmaking/models"
	"log"
	"sort"
	"strconv"
	"time"
)

const maxPlayersPerRoom = 10

func AddUserToQueue(queueRef *db.Ref, user models.User) error {
	err := queueRef.Child(user.ID).Set(context.Background(), user)
	if err != nil {
		log.Printf("Error adding player to queue: %v\n", err)
		return err
	}
	log.Printf("Player added : %s\n", user.Name)
	return nil
}

func Matchmaking(queueRef, roomsRef *db.Ref) {
	for {
		log.Println("Start of matchmaking loop iteration")

		var rawData interface{}
		err := queueRef.Get(context.Background(), &rawData)
		if err != nil {
			log.Printf("Error retrieving tail: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if rawData == nil {
			log.Println("The queue is empty")
			time.Sleep(1 * time.Second)
			continue
		}

		data, ok := rawData.(map[string]interface{})
		if !ok {
			log.Printf("Unexpected data type for matchmakingQueue: %T\n", rawData)
			time.Sleep(1 * time.Second)
			continue
		}

		var queueList []models.User
		for _, value := range data {
			// COnvert user data to user struct
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
			queueList = append(queueList, user)
		}

		sort.Slice(queueList, func(i, j int) bool {
			idI, _ := strconv.Atoi(queueList[i].ID[1:])
			idJ, _ := strconv.Atoi(queueList[j].ID[1:])
			return idI < idJ
		})

		log.Printf("Queue before room creation : %v\n", queueList)

		if len(queueList) >= maxPlayersPerRoom {
			log.Println("Enough players to create a room")

			createRoom(queueList[:maxPlayersPerRoom], roomsRef)

			log.Println("Room created, players removed from queue")

			for _, user := range queueList[:maxPlayersPerRoom] {
				log.Printf("Attempt to delete player ID: %s", user.ID)
				err := queueRef.Child(user.ID).Delete(context.Background())
				if err != nil {
					log.Printf("Error when delete user %s (%s) from the queue: %v", user.Name, user.ID, err)
				} else {
					log.Printf("Player %s (%s) deleted from the queu", user.Name, user.ID)
				}
			}

			verifyQueueUpdate(queueRef, len(queueList)-maxPlayersPerRoom)

			var newRawData interface{}
			err = queueRef.Get(context.Background(), &newRawData)
			if err != nil {
				log.Printf("Error retrieving queue after deletion: %v\n", err)
				continue
			}

			if newRawData == nil {
				log.Printf("Tail after deletion : []\n")
			} else {
				newData, ok := newRawData.(map[string]interface{})
				if !ok {
					log.Printf("Unexpected data type for matchmakingQueue after deletion: %T\n", newRawData)
				} else {
					var remainingUsers []string
					for key := range newData {
						remainingUsers = append(remainingUsers, key)
					}
					log.Printf("Queue after deletion : %v\n", remainingUsers)
				}
			}
		}

		log.Println("End of matchmaking loop iteration")
		time.Sleep(1 * time.Second)
	}
}

func createRoom(players []models.User, roomsRef *db.Ref) {
	roomID := fmt.Sprintf("room-%d", time.Now().Unix())
	var playerIDs []string
	for _, player := range players {
		playerIDs = append(playerIDs, player.ID)
	}

	room := models.Room{
		ID:      roomID,
		Players: playerIDs,
	}

	err := roomsRef.Child(roomID).Set(context.Background(), room)
	if err != nil {
		log.Printf("Error creating room: %v\n", err)
		return
	}

	fmt.Printf("Room %s created with players: %v\n", roomID, playerIDs)

	common.NotifyRoomCreation(roomID)
}

func verifyQueueUpdate(queueRef *db.Ref, expectedCount int) {
	var rawData interface{}
	err := queueRef.Get(context.Background(), &rawData)
	if err != nil {
		log.Printf("Error checking queue after update: %v\n", err)
		return
	}

	actualCount := 0
	if rawData != nil {
		data, ok := rawData.(map[string]interface{})
		if ok {
			actualCount = len(data)
		} else {
			log.Printf("Unexpected data type for matchmakingQueue: %T\n", rawData)
		}
	}

	if actualCount != expectedCount {
		log.Printf("Queue not updated correctly! Expected number : %d, Actual number : %d\n", expectedCount, actualCount)
	} else {
		log.Printf("Queue successfully updated. Current count : %d\n", actualCount)
	}
}
