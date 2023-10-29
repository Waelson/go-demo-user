package controller

import "net/http"

type HealthController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type healthController struct {
}

func (h *healthController) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func NewHealthController() HealthController {
	return &healthController{}
}
