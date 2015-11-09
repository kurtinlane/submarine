package keys


import (
	"fmt"
	"sync"
	"crypto/sha256"
	"encoding/hex"
	"crypto/rand"
	"encoding/base64"
)

type Key struct {
  Id int
  Email string //we will actually not want to store their email at all, just the hash
  Sha256 string
  DO_NOT_STORE_DO_NOT_LOG string
  App int //Id for which application this key belongs to
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

// AddEntry adds a new GuestBookEntry with the provided data.
func (k *Keychain) AddKey(email string, app int) *Key {
	// Acquire our lock and make sure it will be released.
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Get an id for this entry.
	newId := len(k.keys)

	// Create new entry with the given data and the computed newId.
	newEntry := &Key{
		newId,
		email,
		GetSha256Hash(email), 
		getRandomString(32), // need to create random string to act as key
		app,
	}

	// Add entry to the Guest Book.
	k.keys = append(k.keys, newEntry)

	// Return the Id for the new entry.
	return newEntry
}

// GetEntry returns the entry identified by the given id or an error if it can
// not find it.
func (k *Keychain) GetKey(id int) (*Key, error) {
	// Check if we have a valid id.
	if id < 0 || id >= len(k.keys) ||
		k.keys[id] == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// Return the associated entry.
	return k.keys[id], nil
}

// GetAllEntries returns all non-nil entries in the Guest Book.
func (k *Keychain) GetAllKeys() []*Key {
	// Placeholder for the entries we will be returning.
	entries := make([]*Key, 0)

	// Iterate through all existig entries.
	for _, entry := range k.keys {
		if entry != nil {
			// Entry is not nil, so we want to return it.
			entries = append(entries, entry)
		}
	}

	return entries
}

// RemoveAllEntries removes all entries from the Guest Book.
func (k *Keychain) RemoveKey(id int) {
	key, _ := k.GetKey(id)
	
	key.DO_NOT_STORE_DO_NOT_LOG = ""
}

// RemoveAllEntries removes all entries from the Guest Book.
func (k *Keychain) RemoveAllKeys() {
	// Acquire our lock and make sure it will be released.
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Reset guestbook to a new empty one.
	k.keys = []*Key{}
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