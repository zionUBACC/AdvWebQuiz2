// Filename: cmd/api/creator.go
package main

import (
	"fmt"
	"net/http"

	"Quiz2Json.zioncastillo.net/internal/data"
	"Quiz2Json.zioncastillo.net/internal/validator"
)

//create entires hander for the POST /v1/entries endpoint
func (app *application) displayAbout(w http.ResponseWriter, r *http.Request){
	//our target decode destination
	var input struct{
		Name      string   `json:"name"`
		Email     string   `json:"email"`
		Education string   `json:"education"`
		Hobby     string   `json:"hobby"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copying the values
	jsondata := &data.About{
		Name:      input.Name,
		Email:     input.Email,
		Education: input.Education,
		Hobby:     input.Hobby,
	}

	//initialize a new validator instance
	v := validator.New()
	
	//check the map to determine if there were any validation errors
	if data.ValidateEntires(v,jsondata); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)
}