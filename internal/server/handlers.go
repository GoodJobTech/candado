package server

import (
	"encoding/json"
	"net/http"

	"github.com/goodjobtech/candado/internal/errors"
	"github.com/gorilla/mux"
)

type Response struct {
	Error string `json:"error"`
	Data  struct {
		ID    string `json:"id"`
		State int    `json:"state"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (s *Server) AcquireHandler(w http.ResponseWriter, r *http.Request) {
	var response Response

	vars := mux.Vars(r)
	response.Data.ID = vars["id"]

	err := s.db.Lock(vars["id"])
	switch err {
	case nil:
		response.Success = true
		response.Data.State = 1
	case errors.ErrAlreadyLocked:
		response.Success = false
		response.Data.State = 1
		response.Error = err.Error()
	}

	js, _ := json.Marshal(response)
	w.Write(js)
}

func (s *Server) ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	var response Response

	vars := mux.Vars(r)
	response.Data.ID = vars["id"]

	err := s.db.Unlock(vars["id"])
	switch err {
	case nil:
		response.Success = true
		response.Data.State = 0
	case errors.ErrAlreadyUnlocked:
		response.Success = false
		response.Data.State = 0
		response.Error = err.Error()
	}

	js, _ := json.Marshal(response)
	w.Write(js)
}

func (s *Server) HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	var response Response

	vars := mux.Vars(r)
	response.Data.ID = vars["id"]

	state, err := s.db.Heartbeat(vars["id"])
	switch err {
	case errors.ErrLockDoesntExists:
		response.Success = false
		response.Data.State = 0
		response.Error = err.Error()
	case nil:
		response.Data.State = int(state)
	}

	js, _ := json.Marshal(response)
	w.Write(js)

}
