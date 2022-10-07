// Filename: cmd/api/randomize.go

package main
import (
	"crypto/rand"
	"net/http"

	"Quiz2Json.zioncastillo.net/internal/data"
)
type Tools struct{}
const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func GenerateRandomString(length int64) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

func (app *application) outputString(w http.ResponseWriter, r *http.Request){
	
	id,err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w,r)
		return
	}

	randString := GenerateRandomString(id)

	data := data.RandomString{
		Data: randString,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w,r,err)
	}
}
