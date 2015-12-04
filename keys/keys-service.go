
package keys

import (
	"github.com/go-martini/martini"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/kurtinlane/submarine/models"
	"appengine"
)

func RegisterWebService(server *martini.ClassicMartini) {
	path := "/api/v1/keys"
	
	server.Get(path, Get)
	//server.Get(path+"/:email", Get)

	server.Post(path, Post)

	server.Delete(path, Delete)
	server.Delete(path+"/:id", Delete)
}

func Get(params martini.Params, req *http.Request) (int, string) {

	email := params["user"]

	context := appengine.NewContext(req)
	subkey, err := GetKeyForEmail(email, context)
	if err != nil {
		// Entry not found.
		return http.StatusNotFound, "entry not found"
	}

	// Encode entry in JSON.
	encodedSubkey, err := json.Marshal(subkey)
	if err != nil {
		// Failed encoding entry.
		return http.StatusInternalServerError, "internal error"
	}

	// Return encoded entry.
	return http.StatusOK, string(encodedSubkey)
}

// WebPost implements webservice.WebPost.
func Post(params martini.Params,
	req *http.Request, app *models.App) (int, string) {
		
	// Make sure Body is closed when we are done.
	defer req.Body.Close()

	// Read request body.
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	if len(params) != 0 {
		// No keys in params. This is not supported.
		return http.StatusMethodNotAllowed, "method not allowed"
	}

	// Unmarshal entry sent by the user.
	var key models.Key
	err = json.Unmarshal(requestBody, &key)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}

	// Add entry provided by the user.
	context := appengine.NewContext(req) 
	createdKey, _ := AddKey(key.Email, app, context)
	
	encodedKey, err := json.Marshal(createdKey)

	// Everything is fine.
	return http.StatusOK, string(encodedKey)
	
}

func Delete(params martini.Params,
	req *http.Request) (int, string) {
	defer req.Body.Close()
	
	// id, _ := strconv.Atoi(params["id"])
	
	// RemoveKey(id)
	
	return http.StatusOK, "Done."
}