package apps

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"appengine"
    "appengine/datastore"
	"github.com/kurtinlane/submarine/models"
)

func GetApp(id int) (*models.App, error) {
	return nil, nil
}

func GetAppWithApiKey(apiKey string) (*models.App, error) {
	
	return nil, fmt.Errorf("invalid key")
}

func AddApp(name string, context appengine.Context) (*models.App, error) { 
	newKey := datastore.NewIncompleteKey(context, "app", nil)
	newApp := &models.App{
		getUuid(),
		name,
	}
	_, err := datastore.Put(context, newKey, newApp)
    if err != nil {
        return nil, err
    }
	
	// Return the Id for the new entry.
	return newApp, nil
}

func getUuid() string {
	u, _ := uuid.NewV4()
	
	return u.String()
}