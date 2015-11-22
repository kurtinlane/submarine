
package keys

import (
	"github.com/go-martini/martini"
	"net/http"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"github.com/kurtinlane/submarine/models"
	"appengine"
)

func RegisterWebService(server *martini.ClassicMartini) {
	path := "/api/v1/keys"
	
	server.Get(path, Get)
	server.Get(path+"/:id", Get)

	server.Post(path, Post)
	server.Post(path+"/:id", Post)

	server.Delete(path, Delete)
	server.Delete(path+"/:id", Delete)
}

// WebGet implements webservice.WebGet.
func Get(params martini.Params) (int, string) {
	if len(params) == 0 {
		// No params. Return entire collection encoded as JSON.
		encodedEntries, err := json.Marshal(GetAllKeys())
		if err != nil {
			// Failed encoding collection.
			return http.StatusInternalServerError, "internal error"
		}

		// Return encoded entries.
		return http.StatusOK, string(encodedEntries)
	}

	// Convert id to integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Id was not a number.
		return http.StatusBadRequest, "invalid entry id"
	}

	// Get entry identified by id.
	entry, err := GetKey(id)
	if err != nil {
		// Entry not found.
		return http.StatusNotFound, "entry not found"
	}

	// Encode entry in JSON.
	encodedEntry, err := json.Marshal(entry)
	if err != nil {
		// Failed encoding entry.
		return http.StatusInternalServerError, "internal error"
	}

	// Return encoded entry.
	return http.StatusOK, string(encodedEntry)
}

// WebPost implements webservice.WebPost.
func Post(params martini.Params,
	req *http.Request) (int, string) {
		
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
	createdKey, _ := AddKey(key.Email, key.App, context)
	
	encodedKey, err := json.Marshal(createdKey)

	// Everything is fine.
	return http.StatusOK, string(encodedKey)
	
}

func Delete(params martini.Params,
	req *http.Request) (int, string) {
	defer req.Body.Close()
	
	id, _ := strconv.Atoi(params["id"])
	
	RemoveKey(id)
	
	return http.StatusOK, "Done."
}