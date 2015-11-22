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
)

// AddEntry adds a new GuestBookEntry with the provided data.
func AddKey(email string, app int, context appengine.Context) (*models.Key, error) {
	newKey := datastore.NewIncompleteKey(context, "submarinekey", nil)
	newSubmarineKey := &models.Key{
		0,
		email,
		GetSha256Hash(email), 
		getRandomString(32), // need to create random string to act as key
		app,
	}
	_, err := datastore.Put(context, newKey, newSubmarineKey)
    if err != nil {
        return nil, err
    }
	
	return newSubmarineKey, nil
}

// GetEntry returns the entry identified by the given id or an error if it can
// not find it.
func GetKey(id int) (*models.Key, error) {

	return nil, nil
}

// GetAllEntries returns all non-nil entries in the Guest Book.
func GetAllKeys() []*models.Key {
	return nil
}

// RemoveAllEntries removes all entries from the Guest Book.
func RemoveKey(id int) {
	key, _ := GetKey(id)
	
	key.DO_NOT_STORE_DO_NOT_LOG = ""
}

// RemoveAllEntries removes all entries from the Guest Book.
func RemoveAllKeys() {
	
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