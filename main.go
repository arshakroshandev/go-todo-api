package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arshakroshandev/go-todo-api/config"
	"github.com/arshakroshandev/go-todo-api/handlers"
	"github.com/arshakroshandev/go-todo-api/storage"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	if err := storage.InitPostgres(cfg.DBConnectionString()); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Post("/tasks", handlers.CreateTask)
	r.Get("/tasks", handlers.GetAllTasks)
	r.Put("/tasks/{id}", handlers.UpdateTask)
	r.Delete("/tasks/{id}", handlers.DeleteTask)

	fmt.Println("server listening :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
