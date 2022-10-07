// Filename: internal/data/creator.go
package data

import "Quiz2Json.zioncastillo.net/internal/validator"

type About struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Education string   `json:"education"`
	Hobby     string   `json:"hobby"`
}

func ValidateEntires(v *validator.Validator, creator *About)  {
	//use the check method to execute our validation checks
	v.Check(creator.Name != "", "name", "must be provided")
	v.Check(len(creator.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(creator.Email != "", "Email", "must be provided")
	v.Check(len(creator.Email) <= 200, "Email", "must not be more than 200 bytes long")

	v.Check(creator.Education != "", "Education", "must be provided")
	v.Check(len(creator.Education) <= 200, "Education", "must not be more than 200 bytes long")
	
}