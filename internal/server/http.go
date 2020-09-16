package server

import (
	"encoding/json"
	"github.com/camhashemi/doola/internal/user"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Controller user.Controller
}

type CreateUserRequest struct {
	name string
}

type CreateUserResponse struct {
	id int
}

type GetUserRequest struct {
	id int
}

type GetUserResponse struct {
	Id        int
	CreatedAt time.Time
	Name      string
}

func (s Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.CreateUser(w, r)
		case http.MethodGet:
			s.GetUser(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func (s Server) GetUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := s.Controller.GetUser(req.id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	rBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(rBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (s Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := s.Controller.CreateUser(req.name)

	res := CreateUserResponse{
		id: user.Id,
	}

	rBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(rBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
