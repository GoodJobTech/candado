package server

import (
	"encoding/json"
	"net/http"

	"github.com/goodjobtech/candado/internal/errors"
	"github.com/gorilla/mux"
)

type Response struct {
	Error   string `json:"error"`
	State   int    `json:"state"`
	Success bool   `json:"success"`
}

func (s *Server) AcquireHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response Response

	err := s.db.Lock(vars["id"])
	switch err {
	case nil:
		response.Success = true
		response.State = 1
	case errors.ErrAlreadyLocked:
		response.Success = false
		response.State = 1
		response.Error = err.Error()
	}

	js, _ := json.Marshal(response)
	w.Write(js)
}

func (s *Server) ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response Response

	err := s.db.Unlock(vars["id"])
	switch err {
	case nil:
		response.Success = true
		response.State = 0
	case errors.ErrAlreadyUnlocked:
		response.Success = false
		response.State = 0
		response.Error = err.Error()
	}

	js, _ := json.Marshal(response)
	w.Write(js)
}

func (s *Server) HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response Response

	state, err := s.db.Heartbeat(vars["id"])
	switch err {
	case errors.ErrLockDoesntExists:
		response.Success = false
		response.State = 2
		response.Error = err.Error()
	case nil:
		response.State = int(state)
	}

	js, _ := json.Marshal(response)
	w.Write(js)

}
