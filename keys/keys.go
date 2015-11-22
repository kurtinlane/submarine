package keys


import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"crypto/rand"
	"encoding/base64"
	"appengine"
    "appengine/datastore"
	"github.com/kurtinlane/submarine/models"
	"errors"
)

func AddKey(email string, app *models.App, context appengine.Context) (*models.Key, error) {
	appKey := datastore.NewKey(context, "app", "", app.Id, nil)
	newKey := datastore.NewKey(context, "submarinekey", GetSha256Hash(email), 0, appKey)
	
	newSubmarineKey := &models.Key{
		email,
		GetSha256Hash(email), 
		getRandomString(32), // need to create random string to act as key
		app.Id,
	}
	_, err := datastore.Put(context, newKey, newSubmarineKey)
    if err != nil {
        return nil, err
    }
	
	return newSubmarineKey, nil
}

// GetEntry returns the entry identified by the given id or an error if it can
// not find it.
func GetKey(id int, context appengine.Context) (*models.Key, error) {

	return nil, nil
}

func GetKeyForEmail(email string, context appengine.Context) (*models.Key, error) {
	context.Debugf("Getting subKey for email: " + email)
	emailHash := GetSha256Hash(email)
	
	query := datastore.NewQuery("submarinekey").
			Filter("Sha256 =", emailHash)

	iter := query.Run(context)
				
	for {
		var subKey models.Key
		_, err := iter.Next(&subKey)
		if err != nil {
			context.Debugf("Cannot find subKey")
			err := errors.New("cannot find subKey")
			return nil, err
		}
		
		return &subKey, nil
	}
}

// RemoveAllEntries removes all entries from the Guest Book.
func RemoveKey(id int, context appengine.Context) {
	key, _ := GetKey(id, context)
	
	key.DO_NOT_STORE_DO_NOT_LOG = ""
}


func GetSha256Hash(text string) string {
    hasher := sha256.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func getRandomString(numBytes int) string {
	key := make([]byte, numBytes)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error getting random bytes", err.Error())
	}
	
	return base64.StdEncoding.EncodeToString(key)
	
}