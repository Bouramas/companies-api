package store

import (
	"database/sql"
	"fmt"

	"companies-api/cmd/resources"
)

// CompanyService represents a MySQL implementation of companies.CompanyService
type CompanyService struct {
	DB *sql.DB
}

func NewCompanyService(db *sql.DB) *CompanyService {
	cs := new(CompanyService)
	cs.DB = db
	return cs
}

// CreateCompany creates a new Company.
// Returns an error if the Company already exists or if there is an error creating it.
func (cs *CompanyService) CreateCompany(c *resources.Company) error {

	query := "INSERT INTO companies (id, name, description, employees, registered, type) VALUES (UNHEX(REPLACE(?, '-', '')), ?, ?, ?, ?, ?)"
	uuidBytes := []byte(c.ID)
	_, err := cs.DB.Exec(query, uuidBytes, c.Name, c.Description, c.Employees, c.Registered, c.Type)
	if err != nil {
		return fmt.Errorf("error creating new Company: %v", err)
	}

	return nil
}

// PatchCompany updates an existing Company with the given changes.
// Returns an error if the Company doesn't exist or if there is an error updating it.
func (cs *CompanyService) PatchCompany(c *resources.Company) error {
	query := "UPDATE companies SET name = ?, description = ?, employees = ?, registered = ?, type = ? WHERE id = UNHEX(REPLACE(?, '-', ''))"
	uuidBytes := []byte(c.ID)
	_, err := cs.DB.Exec(query, c.Name, c.Description, c.Employees, c.Registered, c.Type, uuidBytes)
	if err != nil {
		return fmt.Errorf("error updating Company: %v", err)
	}

	return nil
}

// DeleteCompany deletes the Company with the given ID.
// Returns an error if the Company doesn't exist or if there is an error deleting it.
func (cs *CompanyService) DeleteCompany(id string) error {
	query := "DELETE FROM companies WHERE id = UNHEX(REPLACE(?, '-', ''))"
	uuidBytes := []byte(id)
	_, err := cs.DB.Exec(query, uuidBytes)
	return err
}

// GetCompany retrieves the Company with the given ID.
// Returns the Company if it exists, or nil and an error if it doesn't.
func (cs *CompanyService) GetCompany(id string) (*resources.Company, bool, error) {
	stmt, err := cs.DB.Prepare("SELECT HEX(id), name, description, employees, registered, type FROM Companies WHERE id = UNHEX(REPLACE(?, '-', ''))")
	if err != nil {
		return nil, false, err
	}
	defer stmt.Close()

	// Execute the query and retrieve the result
	var c resources.Company
	uuidBytes := []byte(id)

	err = stmt.QueryRow(uuidBytes).Scan(&c.ID, &c.Name, &c.Description, &c.Employees, &c.Registered, &c.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &c, true, nil
}

// CompanyExists checks if the company exists
// Returns whether or not the company exists and an error if there is an error.
func (cs *CompanyService) CompanyExists(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM companies WHERE name = ?"
	row := cs.DB.QueryRow(query, name)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, fmt.Errorf("error checking for Company existence: %v", err)
	}
	return count > 0, nil
}
