package resources

import "testing"

func TestCompany_Validate(t *testing.T) {

	tests := []struct {
		name    string
		company *Company
		wantErr bool
	}{
		{"success", &Company{"1", "A Corporation", "A company that makes everything", 100, true, "Corporation"}, false},
		{"name required", &Company{"1", "", "A company that makes everything", 100, true, "Corporation"}, true},
		{"invalid type", &Company{"1", "B Corporation", "A company that makes everything", 100, true, "Invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Company{
				ID:          tt.company.ID,
				Name:        tt.company.Name,
				Description: tt.company.Description,
				Employees:   tt.company.Employees,
				Registered:  tt.company.Registered,
				Type:        tt.company.Type,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
