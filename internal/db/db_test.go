package db

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

var kvdb *KeyValueDB
func TestMain(m *testing.M) {
	// set up the database
	kvdb = NewKeyValueDB(map[string]string{"b": "B"})

	// run the tests
	code := m.Run()

	os.Exit(code)
}

func TestGetKeys(t *testing.T) {
	emptykvdb := NewKeyValueDB(make(map[string]string, 0))

	keys := emptykvdb.GetKeys()
	assert.Equal(t, keys, make([]string, 0))
}

func TestUpdateValue(t *testing.T) {
	kv, err := kvdb.UpdateValue("a", "A")
	assert.Equal(t, kv, map[string]string{})
	assert.Equal(t, err, ErrInexistentKey)

	kv, err = kvdb.UpdateValue("b", "b")
	assert.Equal(t, kv, map[string]string{"b": "b"})
	assert.Equal(t, err, nil)
}

func TestGetValue(t *testing.T) {
	val, err := kvdb.GetValue("a")
	assert.Equal(t, val, "")
	assert.Equal(t, err, ErrInexistentKey)

	val, err = kvdb.GetValue("b")
	assert.Equal(t, val, "b")
	assert.Nil(t, err)
}

func TestDeleteValue(t *testing.T) {
	val, err := kvdb.DeleteValue("a")
	assert.Equal(t, val, "")
	assert.Equal(t, err, ErrInexistentKey)

	val, err = kvdb.DeleteValue("b")
	assert.Equal(t, val, "b")
	assert.Nil(t, err)
}