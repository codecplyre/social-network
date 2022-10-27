package event

import (
	"fmt"
	"time"

	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func AllEventByGroup(groupId string, database *structs.DB) ([]structs.Event, error) {
	var event structs.Event
	var events []structs.Event
	var err error
	rows, err := database.DB.Query("SELECT * FROM Event WHERE groupId = '" + groupId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var eventId, userId, groupId2, name, description, location, startTime, endTime, createdAt string
	for rows.Next() {
		rows.Scan(&eventId, &userId, &groupId2, &name, &description, &location, &startTime, &endTime, &createdAt)
		event = structs.Event{
			EventId:     eventId,
			UserId:      userId,
			GroupId:     groupId2,
			Name:        name,
			Description: description,
			Location:    location,
			StartTime:   startTime,
			EndTime:     endTime,
			CreatedAt:   createdAt,
		}
		events = append([]structs.Event{event}, events...)
	}
	return events, nil
}

func AllEventParticipant(eventId string, database *structs.DB) ([]structs.EventParticipant, error) {
	var eventParticipant structs.EventParticipant
	var eventParticipants []structs.EventParticipant
	var err error
	rows, err := database.DB.Query("SELECT * FROM EventParticipant WHERE eventId = '" + eventId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var eventId2, userId, createdAt string
	for rows.Next() {
		rows.Scan(&eventId2, &userId, &createdAt)
		eventParticipant = structs.EventParticipant{
			EventId:   eventId2,
			UserId:    userId,
			CreatedAt: createdAt,
		}
		eventParticipants = append([]structs.EventParticipant{eventParticipant}, eventParticipants...)
	}
	return eventParticipants, nil
}

func GetEventByEventId(eventId string, database *structs.DB) (structs.Event, error) {
	var event structs.Event
	var err error

	rows, err := database.DB.Query("SELECT * FROM Event WHERE eventId = '" + eventId + "'")
	if err != nil {
		fmt.Print(err)
		return event, err
	}
	var eventId2, userId, groupId, name, description, location, startTime, endTime, createdAt string
	for rows.Next() {
		rows.Scan(&eventId2, &userId, &groupId, &name, &description, &location, &startTime, &endTime, &createdAt)
		event = structs.Event{
			EventId:     eventId2,
			UserId:      userId,
			GroupId:     groupId,
			Name:        name,
			Description: description,
			Location:    location,
			StartTime:   startTime,
			EndTime:     endTime,
			CreatedAt:   createdAt,
		}

	}
	return event, nil
}

func AllUserEvent(userId string, database *structs.DB) ([]structs.Event, error) {
	var eventStrA []string
	var events []structs.Event

	var err error
	rows, err := database.DB.Query("SELECT eventId FROM Event WHERE userId = '" + userId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var eventStr string
	for rows.Next() {
		rows.Scan(&eventStr)
		eventStrA = append([]string{eventStr}, eventStrA...)
	}

	for _, id := range eventStrA {
		userEvent, err := GetEventByEventId(id, database)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		events = append([]structs.Event{userEvent}, events...)
	}

	return events, nil
}

func AddEven(groupId string, event structs.Event, database *structs.DB) (string, error) {
	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	eventId := uuid.NewV4()

	stmt, _ := database.DB.Prepare(`
	INSERT INTO Event values (?, ?, ?, ?, ?, ?, ?, ?, ?)
`)

	_, err := stmt.Exec(eventId, event.UserId, groupId, event.Name, event.Description, event.Location, event.StartTime, event.EndTime, createdAt)
	if err != nil {
		fmt.Println("inside Create Addevent", err)
		return "", err
	}
	return eventId.String(), nil
}
