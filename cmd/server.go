package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"mime"
	"net/http"
	"strings"

	"companies-api/cmd/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Log            *logrus.Logger
	CompanyService *store.CompanyService

	srv *http.Server
}

func NewServer(cfg *config, logger *logrus.Logger) (*Server, error) {
	if cfg == nil {
		return nil, errors.New("config cannot be nil")
	}

	s := new(Server)
	r := s.NewRouter()

	s.CompanyService = store.NewCompanyService(cfg.DB)
	s.Log = logger
	w := s.Log.Writer()
	defer w.Close()

	s.srv = &http.Server{
		Addr:     ":8080",
		ErrorLog: log.New(w, "", 0),
		Handler:  r,
	}
	return s, nil
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) hasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	return compareContentTypes(contentType, mimetype)
}

func compareContentTypes(contentType, mimetype string) bool {
	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}

	return false
}

func (s *Server) writeJSONError(w http.ResponseWriter, code int, msg string) {
	errObject := map[string]interface{}{"error": true, "code": code, "message": msg}
	res, _ := json.Marshal(errObject)

	w.WriteHeader(code)
	_, _ = w.Write(res)
}
