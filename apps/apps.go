package apps

import (
	"sync"
	"fmt"
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

func NewAppsList() *Apps {
	return &Apps{
		make([]*App, 0),
		new(sync.Mutex),
	}
}

func (a *Apps) GetApp(apiKey string) (*App, error) {

	// Return the associated entry.
	//return nil, nil
	return nil, fmt.Errorf("not implemented")
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
		"123", //GetApiKey(),
		name,
	}

	// Add entry to the Guest Book.
	a.apps = append(a.apps, newEntry)

	// Return the Id for the new entry.
	return newEntry
}
