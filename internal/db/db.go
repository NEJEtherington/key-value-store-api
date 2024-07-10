package db

import (
	"errors"
	"sync"
)

// KeyValueDB represents a key-value store
// Uses a mutex to synchronise access to the in-memory map and ensure safe concurrent operations
type KeyValueDB struct {
 data map[string]string
 mu   sync.RWMutex
}

// ErrInexistentKey is a custom error used when a key does not exist in the DB
var ErrInexistentKey = errors.New("Key does not exist")

// NewKeyValueDB creates a new instance of KeyValueDB
func NewKeyValueDB(data map[string]string) *KeyValueDB {
 return &KeyValueDB{
  data: data,
 }
}

// GetKeys returns all the values present in the DB
func (kv *KeyValueDB) GetKeys() []string {
	kv.mu.RLock()
 	defer kv.mu.RUnlock()
	
	// initialise as opposed to just declaring so that an empty list (rather than the nil value of 0) is returned if the database is empty
	keys := make([]string, 0)
	
	for k := range kv.data {
				keys = append(keys, k)
	}

	return keys
}

// UpdateValue updates a key-value pair in the DB
func (kv *KeyValueDB) UpdateValue(key, value string) (map[string]string, error){
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, ok := kv.data[key]; !ok {
		
		return make(map[string]string, 0), ErrInexistentKey
	}

 	kv.data[key] = value

	return map[string]string{key: value}, nil
}


// GetValue retrieves the value associated with a key from the store
func (kv *KeyValueDB) GetValue(key string) (string, error) {
 	kv.mu.RLock()
	defer kv.mu.RUnlock()
	
	var val string
	var ok bool 

 	if val, ok = kv.data[key]; !ok {
		 
		return "", ErrInexistentKey
	 }
 
	return val, nil
}

// DeleteValue deletes a key value pair from the DB
func (kv *KeyValueDB) DeleteValue(key string) (string, error){
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, ok := kv.data[key]; !ok {
		
		return "", ErrInexistentKey
	}

	delete(kv.data, key)

	return key, nil
}