package controller

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	ac ApiController
}

func NewRouter(ac ApiController) Router {
	return &router{ac}
}

func (ro *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.ac.GetDoing(w, r)
	case "POST":
		ro.ac.PostDoing(w, r)
	case "PUT":
		log.Info().Msg("a")
	case "DELETE":
		log.Info().Msg("b")
	default:
		w.WriteHeader(405)
	}
}
