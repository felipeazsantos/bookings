package dbrepo

import (
	"errors"
	"time"

	"github.com/felipeazsantos/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesAndRoomID returns true if availability exists for roomID, and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesAndRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gest a room by id
func (m *testDBRepo) GetRoomByID(roomID int) (models.Room, error) {
	var room models.Room
	if roomID > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}

// GetUserByID returns a user by ID
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var user models.User
	return user, nil
}

// UpdateUser update the user in the database
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate authenticate the user
func (m *testDBRepo) Authenticate(email, password string) (int, string, error) {
	return 0, "", nil
}

// AllReservations returns a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	return nil, nil
}

// AllNewReservations returns a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	return nil, nil
}

// GetReservationByID returns one reservation by id
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	return models.Reservation{}, nil
}

// UpdateReservation update a reservation in the database
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// DeleteReservation deletes one reservation by id in the database
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by id
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	return nil, nil
}

// GetRestrictionsForRoomByDate returns restrictions for a room by date range
func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	return nil, nil
}
