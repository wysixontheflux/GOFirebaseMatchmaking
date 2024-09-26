package common

import "log"

func NotifyRoomCreation(roomID string) {
	log.Printf("Notification : Room %s created.\n", roomID)
}
