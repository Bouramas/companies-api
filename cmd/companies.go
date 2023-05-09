package main

import (
	"companies-api/cmd/resources"
)

// CompanyService defines methods for performing CRUD operations on a Company.
type CompanyService interface {
	// CreateCompany creates a new Company.
	// Returns an error if the Company already exists or if there is an error creating it.
	CreateCompany(c *resources.Company) error

	// PatchCompany updates an existing Company with the given changes.
	// Returns an error if the Company doesn't exist or if there is an error updating it.
	PatchCompany(c *resources.Company) error

	// DeleteCompany deletes the Company with the given ID.
	// Returns an error if the Company doesn't exist or if there is an error deleting it.
	DeleteCompany(id string) error

	// GetCompany retrieves the Company with the given ID.
	// Returns the Company if it exists, or nil and an error if it doesn't.
	GetCompany(id string) (*resources.Company, error)
}
