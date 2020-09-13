package server

import (
	"bytes"
	"encoding/json"
	"github.com/camhashemi/doola/internal/user"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServer_CreateUser(t *testing.T) {
	r := user.FakeRepo{}
	c := user.NewController(&r)
	s := Server{
		Controller: c,
	}

	go s.Start()

	url := "http://localhost:8080/users"

	t.Run("test get returns 404", func(t *testing.T) {
		httpResponse, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, httpResponse.StatusCode)
	})

	t.Run("test post returns id", func(t *testing.T) {
		req := CreateUserRequest{
			name: "Cam",
		}
		reqBytes, err := json.Marshal(req)
		httpResponse, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, httpResponse.StatusCode)

		var createUserResponse CreateUserResponse
		err = json.NewDecoder(httpResponse.Body).Decode(&createUserResponse)
		assert.NoError(t, err)

		assert.Equal(t, 0, createUserResponse.id)
	})
}

func TestServer_GetUser(t *testing.T) {
	r := user.FakeRepo{}
	c := user.NewController(&r)
	s := Server{
		Controller: c,
	}

	go s.Start()

	url := "http://localhost:8080/users"

	t.Run("test get returns 404", func(t *testing.T) {
		httpResponse, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, httpResponse.StatusCode)
	})

	t.Run("test post returns id", func(t *testing.T) {
		req := CreateUserRequest{
			name: "Cam",
		}
		reqBytes, err := json.Marshal(req)
		httpResponse, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, httpResponse.StatusCode)

		var createUserResponse CreateUserResponse
		err = json.NewDecoder(httpResponse.Body).Decode(&createUserResponse)
		assert.NoError(t, err)

		assert.Equal(t, 0, createUserResponse.id)
	})

}
