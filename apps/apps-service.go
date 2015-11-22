package apps

import (
	//"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"encoding/json"
	//"strconv"
	"io/ioutil"
	"appengine"
	"github.com/kurtinlane/submarine/models"
)

func RegisterWebService(server *martini.ClassicMartini) {
	path := "/apps"
	
	server.Get(path, Get)
	server.Get(path+"/:id", Get)

	server.Post(path, Post)
	server.Post(path+"/:id", Post)

	server.Delete(path, Delete)
	server.Delete(path+"/:id", Delete)
}

func Get(params martini.Params) (int, string) {
	
	return http.StatusOK, string("")
}

func Post(params martini.Params, req *http.Request) (int, string) {
	// Make sure Body is closed when we are done.
	defer req.Body.Close()
	
	// Read request body.
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}
	
	// No keys in params. This is not supported.
	if len(params) != 0 {
		return http.StatusMethodNotAllowed, "method not allowed"
	}
	
	// Unmarshal new app sent in request.
	var app models.App
	err = json.Unmarshal(requestBody, &app)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}

	// Add entry provided by the user.
	context := appengine.NewContext(req) 
	createdApp, err := AddApp(app.Name, context)
	
	encodedApp, err := json.Marshal(createdApp)

	// Everything is fine.
	return http.StatusOK, string(encodedApp)
}

func Delete(params martini.Params, req *http.Request) (int, string) {
	defer req.Body.Close()
	
	return http.StatusMethodNotAllowed, "method not allowed"
}
