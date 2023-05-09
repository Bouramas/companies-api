package resources

import (
	"errors"
	"fmt"
	"strings"
)

type Company struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Employees   int    `json:"employees" validate:"required"`
	Registered  bool   `json:"registered" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=Corporation NonProfit Cooperative SoleProprietorship"`
}

// Validate performs basic sanity checks on the request payload.
func (c *Company) Validate() error {
	var required []string

	if c.Name == "" {
		required = append(required, "name")
	}
	if len(c.Name) > 15 {
		return errors.New("field name should be 15 chars max")
	}
	if len(c.Description) > 3000 {
		return errors.New("field name should be 3000 chars max")
	}
	if c.Type == "" {
		required = append(required, "type")
	}
	if err := isCompanyTypeValid(c.Type); err != nil {
		return err
	}
	if len(required) > 0 {
		return errors.New(strings.Join(required, ", ") + " required.")
	}

	return nil
}

func (c *Company) ApplyPatch(patchData map[string]interface{}) error {
	// Loop over the patch data and update the company fields
	for key, value := range patchData {
		switch key {
		case "name":
			name, ok := value.(string)
			if !ok {
				return fmt.Errorf("name must be a string")
			}
			c.Name = name
		case "description":
			if value == nil {
				c.Description = ""
			} else {
				description, ok := value.(string)
				if !ok {
					return fmt.Errorf("description must be a string")
				}
				c.Description = description
			}
		case "employees":
			employees, ok := value.(float64)
			if !ok {
				return fmt.Errorf("employees must be a number")
			}
			c.Employees = int(employees)
		case "registered":
			registered, ok := value.(bool)
			if !ok {
				return fmt.Errorf("registered must be a boolean")
			}
			c.Registered = registered
		case "type":
			cType, ok := value.(string)
			if !ok {
				return fmt.Errorf("type must be a string")
			}
			if err := isCompanyTypeValid(cType); err != nil {
				return err
			}
			c.Type = cType
		default:
			return fmt.Errorf("unsupported field: %s", key)
		}
	}
	return nil
}

func isCompanyTypeValid(companyType string) error {
	validTypes := []string{"Corporation", "NonProfit", "Cooperative", "SoleProprietorship"}

	for _, t := range validTypes {
		if t == companyType {
			return nil
		}
	}
	return fmt.Errorf("invalid field type (valid types: \"Corporation\", \"NonProfit\", \"Cooperative\", \"SoleProprietorship\")")
}
