package apps

import (
	"github.com/go-martini/martini"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// GetPath implements webservice.GetPath.
func (a *Apps) GetPath() string {
	// Associate this service with http://host:port/keys.
	return "/apps"
}

// WebGet implements webservice.WebGet.
func (a *Apps) WebGet(params martini.Params) (int, string) {
	if len(params) == 0 {
		return http.StatusBadRequest, "must provide api key"
	}

	// Convert id to integer.
	apiKey := params["API-KEY"]

	app, err := a.GetApp(apiKey)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}
	
	// Encode entry in JSON.
	encodedEntry, err := json.Marshal(app)
	if err != nil {
		// Failed encoding entry.
		return http.StatusInternalServerError, "internal error"
	}
	
	return http.StatusOK, string(encodedEntry)
}

// WebPost implements webservice.WebPost.
func (a *Apps) WebPost(params martini.Params,
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
	var app App
	err = json.Unmarshal(requestBody, &app)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}

	// Add entry provided by the user.
	createdApp := a.AddApp(app.Name)
	
	encodedApp, err := json.Marshal(createdApp)

	// Everything is fine.
	return http.StatusOK, string(encodedApp)
	
}