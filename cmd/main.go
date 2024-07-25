package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lohanguedes/first-api/internal/store"
)

type userRepo interface {
	FindAll() []store.User
	FindById(uuid string) (store.User, error)
	Insert(firstName, lastName, bio string) (store.User, error)
	Update(uuid string, u store.User) (store.User, error)
	Delete(uuid string) (store.User, error)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	app := application{
		db:     store.UserRepo{},
		router: r,
	}

	http.ListenAndServe("localhost:3000", app.router)
}
