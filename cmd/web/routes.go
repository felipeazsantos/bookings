package main

import (
	"net/http"

	"github.com/felipeazsantos/bookings/internal/config"
	"github.com/felipeazsantos/bookings/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/rooms/generals", handlers.Repo.GeneralsRoom)
	router.Get("/rooms/majors", handlers.Repo.MajorsRoom)
	router.Get("/contact", handlers.Repo.Contact)

	router.Get("/search-availability", handlers.Repo.Availability)
	router.Post("/search-availability", handlers.Repo.PostAvailability)
	router.Post("/search-availability-json", handlers.Repo.AvailabilityJson)
	
	router.Get("/make-reservation", handlers.Repo.Reservation)
	router.Post("/make-reservation", handlers.Repo.PostReservation)
	router.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
