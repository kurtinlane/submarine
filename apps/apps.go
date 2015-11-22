package apps

import (
	"github.com/nu7hatch/gouuid"
	"appengine"
    "appengine/datastore"
	"github.com/kurtinlane/submarine/models"
	"errors"
)

func GetApp(id int) (*models.App, error) {
	return nil, nil
}

func GetAppWithApiKey(apiKey string, context appengine.Context) (*models.App, error) {
	context.Debugf("Getting app with apikey: " + apiKey)
	
	query := datastore.NewQuery("app").
			Filter("SECRET_API_KEY =", apiKey)

	iter := query.Run(context)
				
	for {
		var app models.App
		key, err := iter.Next(&app)
		if err != nil {
			context.Debugf("Cannot find app")
			err := errors.New("cannot find app")
			return nil, err
		}
		
		app.Id = key.IntID()
		
		return &app, nil
	}
}

func AddApp(name string, context appengine.Context) (*models.App, error) { 
	newKey := datastore.NewIncompleteKey(context, "app", nil)
	newApp := &models.App{
		getUuid(),
		name,
		0,
	}
	completeKey, err := datastore.Put(context, newKey, newApp)
    if err != nil {
        return nil, err
    }
	
	newApp.Id = completeKey.IntID()
	
	// Return the Id for the new entry.
	return newApp, nil
}

func getUuid() string {
	u, _ := uuid.NewV4()
	
	return u.String()
}