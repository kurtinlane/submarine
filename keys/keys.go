
package keys

import (
	"github.com/go-martini/martini"
	"sync"
	"net/http"
)

type Key struct {
  Email string //we will actually not want to store their email at all, just the hash
  MD5 string
  App int //Id for which application this key belongs to
  Key string
}

type Keychain struct {
	keys []*Key
	mutex *sync.Mutex
}

func NewKeychain() *Keychain {
	return &Keychain{
		make([]*Key, 0),
		new(sync.Mutex),
	}
}

// GetPath implements webservice.GetPath.
func (k *Keychain) GetPath() string {
	// Associate this service with http://host:port/keys.
	return "/keys"
}

// WebGet implements webservice.WebGet.
func (k *Keychain) WebGet(params martini.Params) (int, string) {

	// Return encoded entry.
	return http.StatusOK, "hello"
}

// WebPost implements webservice.WebPost.
func (k *Keychain) WebPost(params martini.Params,
	req *http.Request) (int, string) {
	// Make sure Body is closed when we are done.
	defer req.Body.Close()

	// Everything is fine.
	return http.StatusOK, "new entry created"
}