package apps

import (
	"sync"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type App struct {
  Id int
  SECRET_API_KEY string
  Name string
}

type Apps struct {
	apps []*App
	mutex *sync.Mutex
}

var a = &Apps{
	make([]*App, 0),
	new(sync.Mutex),
}

func RetreiveAppsList() *Apps {
	return a
}

func (a *Apps) GetApp(id int) (*App, error) {
	if id < 0 || id >= len(a.apps) ||
		a.apps[id] == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// Return the associated entry.
	return a.apps[id], nil
}

func GetAppWithApiKey(apiKey string) (*App, error) {
	
	for i := 0; i < len(a.apps); i++ {
		if a.apps[i].SECRET_API_KEY == apiKey {
			return a.apps[i], nil
		}
	}
	
	return nil, fmt.Errorf("invalid key")
}

func (a *Apps) AddApp(name string) *App {
	// Acquire our lock and make sure it will be released.
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// Get an id for this entry.
	newId := len(a.apps)

	// Create new entry with the given data and the computed newId.
	newEntry := &App{
		newId,
		getUuid(),
		name,
	}

	// Add entry to the Guest Book.
	a.apps = append(a.apps, newEntry)

	// Return the Id for the new entry.
	return newEntry
}

func getUuid() string {
	u, _ := uuid.NewV4()
	
	return u.String()
}