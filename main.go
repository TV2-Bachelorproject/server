package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/controller/users"
	"github.com/TV2-Bachelorproject/server/middleware"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
)

func routes(r *mux.Router) {
	u := mux.NewRouter()
	r.Handle("/users", u)
	r.Handle("/users/{id:[0-9]+}", u)
	u.Use(middleware.Authenticated(user.Admin))
	u.HandleFunc("/users", users.List).Methods("GET")
	u.HandleFunc("/users", users.Create).Methods("POST")
	u.HandleFunc("/users/{id:[0-9]+}", users.Show).Methods("GET")
	u.HandleFunc("/users/{id:[0-9]+}", users.Update).Methods("PUT")
	u.HandleFunc("/users/{id:[0-9]+}", users.Delete).Methods("DELETE")

	p := mux.NewRouter()
	r.Handle("/people", p)
	r.Handle("/people/{id:[0-9]+}", p)
	p.Use(middleware.Authenticated(user.Admin, user.Producer))
	p.HandleFunc("/people", people.List).Methods("GET")
	p.HandleFunc("/people", people.Create).Methods("POST")
	p.HandleFunc("/people/{id:[0-9]+}", people.Show).Methods("GET")
	p.HandleFunc("/people/{id:[0-9]+}", people.Update).Methods("PUT")
	p.HandleFunc("/people/{id:[0-9]+}", people.Delete).Methods("DELETE")

	r.HandleFunc("/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/auth/refresh", auth.Refresh).Methods("POST")
}

func main() {
	model.Migrate()

	u1, err := user.New("admin", "admin@example.com", "123456", user.Admin)

	if err != nil {
		log.Fatal(err)
	}

	u2, err := user.New("producer", "producer@example.com", "123456", user.Producer)

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&u1)
	db.Create(&u2)

	r := mux.NewRouter()
	routes(r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
