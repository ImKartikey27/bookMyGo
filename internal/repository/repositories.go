package repository

import "gorm.io/gorm"

type Repositories struct {
    Theater TheaterRepository
    Hall    HallRepository
    Movie   MovieRepository
    Show    ShowRepository
    Seat    SeatRepository
    Booking BookingRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
    return &Repositories{
        Theater: NewTheaterRepository(db),
        Hall:    NewHallRepository(db),
        Movie:   NewMovieRepository(db),
        Show:    NewShowRepository(db),
        Seat:    NewSeatRepository(db),
        Booking: NewBookingRepository(db),
    }
}