package main

import (
	"net/http"

	"github.com/felipeazsantos/bookings/pkg/config"
	"github.com/felipeazsantos/bookings/pkg/handlers"
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
	router.Get("/rooms/generals-quarters", handlers.Repo.GeneralsRoom)
	router.Get("/rooms/majors-suite", handlers.Repo.MajorsRoom)
	router.Get("/contact", handlers.Repo.Contact)
	router.Get("/reservation", handlers.Repo.Reservation)
	router.Get("/make-reservation", handlers.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
