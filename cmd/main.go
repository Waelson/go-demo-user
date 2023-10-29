package main

import (
	"github.com/Waelson/go-demo-user/internal/controller"
	"github.com/Waelson/go-demo-user/internal/database"
	"github.com/Waelson/go-demo-user/internal/repository"
	"github.com/Waelson/go-demo-user/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.NewDatabase()

	connection, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	err = db.Init(connection)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	healthController := controller.NewHealthController()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)

	r.Get("/ping", healthController.Get)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userController.Post)
		r.Get("/", userController.Get)
		r.Delete("/", userController.Delete)
		r.Put("/", userController.Update)
	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
