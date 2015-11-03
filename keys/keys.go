package keys


import (
	"fmt"
	"sync"
)

type Key struct {
  Id int
  Email string //we will actually not want to store their email at all, just the hash
  MD5 string
  Key string
  App string //Id for which application this key belongs to
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
func (k *Keychain) AddKey(email, app string) *Key {
	// Acquire our lock and make sure it will be released.
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Get an id for this entry.
	newId := len(k.keys)

	// Create new entry with the given data and the computed newId.
	newEntry := &Key{
		newId,
		email,
		"123",
		"345",
		app,
	}

	// Add entry to the Guest Book.
	k.keys = append(k.keys, newEntry)

	// Return the Id for the new entry.
	return newEntry
}

// GetEntry returns the entry identified by the given id or an error if it can
// not find it.
func (k *Keychain) GetEntry(id int) (*Key, error) {
	// Check if we have a valid id.
	if id < 0 || id >= len(k.keys) ||
		k.keys[id] == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// Return the associated entry.
	return k.keys[id], nil
}

// GetAllEntries returns all non-nil entries in the Guest Book.
func (k *Keychain) GetAllEntries() []*Key {
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
func (k *Keychain) RemoveAllEntries() {
	// Acquire our lock and make sure it will be released.
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Reset guestbook to a new empty one.
	k.keys = []*Key{}
}