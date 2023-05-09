package main

import (
	"net/http"

	"companies-api/cmd/middleware"
	"github.com/gorilla/mux"
)

func (s *Server) NewRouter() *mux.Router {

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.AuthHandler())

	companiesRouter := apiRouter.PathPrefix("/companies").Subrouter()
	companiesRouter.HandleFunc("/", s.HandlePostCompanies()).Methods(http.MethodPost, http.MethodOptions)
	companiesRouter.HandleFunc("/{id}/", s.HandlePatchCompanies()).Methods(http.MethodPatch, http.MethodOptions)
	companiesRouter.HandleFunc("/{id}/", s.HandleDeleteCompanies()).Methods(http.MethodDelete, http.MethodOptions)
	companiesRouter.HandleFunc("/{id}/", s.HandleGetCompanies()).Methods(http.MethodGet, http.MethodOptions)

	r.NotFoundHandler = r.NewRoute().BuildOnly().HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			middleware.SetSecureHeaders(w)
			s.writeJSONError(w, http.StatusNotFound, "Not found.")
		},
	).GetHandler()
	r.MethodNotAllowedHandler = r.NewRoute().BuildOnly().HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			middleware.SetSecureHeaders(w)
			s.writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed.")
		},
	).GetHandler()

	r.Use(
		middleware.CORSHandler(),
	)
	return r
}
