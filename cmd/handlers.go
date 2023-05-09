package main

import (
	"encoding/json"
	"io"
	"net/http"

	"companies-api/cmd/resources"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) HandlePostCompanies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !s.hasContentType(r, "application/json") {
			s.writeJSONError(w, http.StatusBadRequest, "Content-Type is not application/json.")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		company := &resources.Company{}
		if err = json.Unmarshal(body, &company); err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		company.ID = uuid.New().String()

		err = company.Validate()
		if err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		exists, err := s.CompanyService.CompanyExists(company.Name)
		if err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if exists {
			s.writeJSONError(w, http.StatusBadRequest, "A company with the same name already exists")
			return
		}

		err = s.CompanyService.CreateCompany(company)
		if err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _ := json.Marshal(company)
		_, _ = w.Write(res)
	}
}

func (s *Server) HandlePatchCompanies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !s.hasContentType(r, "application/json") {
			s.writeJSONError(w, http.StatusBadRequest, "Content-Type is not application/json.")
			return
		}

		// parse id from route variable
		id := mux.Vars(r)["id"]

		company, exists, err := s.CompanyService.GetCompany(id)
		if err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if !exists {
			s.writeJSONError(w, http.StatusBadRequest, "No company exists with the specified ID.")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parse the patch request body
		var patchData map[string]interface{}
		if err = json.Unmarshal(body, &patchData); err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Apply Patch and Validate
		err = company.ApplyPatch(patchData)
		if err != nil {
			s.writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = s.CompanyService.PatchCompany(company)
		if err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _ := json.Marshal(company)
		_, _ = w.Write(res)
	}
}

func (s *Server) HandleGetCompanies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parse id from route variable
		id := mux.Vars(r)["id"]

		company, exists, err := s.CompanyService.GetCompany(id)
		if err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if !exists {
			s.writeJSONError(w, http.StatusBadRequest, "No company exists with the specified ID.")
			return
		}

		res, _ := json.Marshal(company)
		_, _ = w.Write(res)
	}
}

func (s *Server) HandleDeleteCompanies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parse id from route variable
		id := mux.Vars(r)["id"]

		// delete company
		if err := s.CompanyService.DeleteCompany(id); err != nil {
			s.writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
