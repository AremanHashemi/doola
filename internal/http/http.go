package http

import (
	"doola/internal/user"
	"encoding/json"
	"log"
	"net/http"
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

func (s Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.CreateUser(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
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
