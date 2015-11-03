package main

import (
  "github.com/go-martini/martini"
  "github.com/kurtinlane/submarine/keys"
  "net/http"
)

func main() {
  m := martini.Classic()
  
  keychain := keys.NewKeychain()
  
  registerWebService(keychain, m)
  
  m.Run()
}

// WebService is the interface that should be implemented by types that want to
// provide web services.
type WebService interface {
	// GetPath returns the path to be associated with the service.
	GetPath() string

	// WebGet is Just as above, but for the GET method. If params is empty,
	// it returns all the entries in the collection. Otherwise it returns
	// the entry with the id as per the "id" key in params.
	WebGet(params martini.Params) (int, string)

	// WebPost wraps the POST method. Again an empty params means that the
	// request should be applied to the collection. A non-empty param will
	// have an "id" key that refers to the entry that should be processed
	// (note this specific case is usually not supported unless each entry
	// is also a collection).
	WebPost(params martini.Params, req *http.Request) (int, string)
}

func registerWebService(webService WebService,
	classicMartini *martini.ClassicMartini) {
	path := webService.GetPath()

	classicMartini.Get(path, webService.WebGet)
	classicMartini.Get(path+"/:id", webService.WebGet)

	classicMartini.Post(path, webService.WebPost)
	classicMartini.Post(path+"/:id", webService.WebPost)

}