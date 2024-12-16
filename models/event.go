package models

import (
	"errors"

	"azno-space.com/azno/db"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Price       string `binding:"required"`
	UserId      int64
}

func (event *Event) SaveEvent() error {
	query := `INSERT INTO events (name,description,location,price,user_id) VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(&event.Name, &event.Description, &event.Location, &event.Price, &event.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	events := []Event{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tempEvent Event
		err := rows.Scan(&tempEvent.Id, &tempEvent.Name, &tempEvent.Description, &tempEvent.Location, &tempEvent.Price, &tempEvent.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, tempEvent)
	}

	return events, nil

}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = ? 
	`

	row := db.DB.QueryRow(query, id)
	tempEvent := Event{}
	err := row.Scan(&tempEvent.Id, &tempEvent.Name, &tempEvent.Description, &tempEvent.Location, &tempEvent.Price, &tempEvent.UserId)

	if err != nil {
		return nil, err
	}

	return &tempEvent, nil

}

func DeleteEventById(id int64) error {

	query := `DELETE FROM events WHERE id = ? `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return err

}

func UpdateEventById(id int64, event Event) error {
	query := `UPDATE events SET 
	name = ?, 
	description = ?,
	location = ?,
	price = ?, 
	user_id = ? 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Price, event.UserId, id)

	return err
}

func (e *Event) Register(userId int64) (int64, error) {
	query := `
		INSERT INTO registration (event_id , user_id) VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant add registeration for this event")
	}

	result, err := stmt.Exec(e.Id, userId)

	if err != nil {
		return -1, errors.New("cant save registration for this event ")
	}

	id, err := result.LastInsertId()
	defer stmt.Close()
	return id, err

}

func (e *Event) CancelRegisteration(userId int64) (int64, error) {
	query := `DELETE FROM registration WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, errors.New("Cant delete this registration")
	}

	result, err := stmt.Exec(e.Id, userId)
	rows, err := result.RowsAffected()

	return rows, err

}
