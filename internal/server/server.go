package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kvp-api/internal/db"
)

// PutRequestBody represents the value included as the request body of an http PUT request
type PutRequestBody struct {
	NewValue string `json:"value"`
}


// InitRoutes sets up the router with endpoints and controllers
func InitRoutes() *gin.Engine {	
	kvdb := db.NewKeyValueDB(make(map[string]string, 0)) 

	router := gin.Default()
	router.GET("/", getKeys(kvdb))
	router.GET("/:key", getValue(kvdb))
	router.PUT("/:key", updateValue(kvdb))
	router.DELETE("/:key", deleteValue(kvdb))

	return router
}

// Get all keys from the database
func getKeys(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var keys []string
		keys = kv.GetKeys()
	
		c.JSON(http.StatusOK, keys)
	}

	return gin.HandlerFunc(fn)
}


// Get the value for a key
// Return a 404 if key does not exist
func getValue(kv *db.KeyValueDB) gin.HandlerFunc{
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var val string
		var err error

		if val, err = kv.GetValue(key); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		c.String(http.StatusOK, val)
	}

	return gin.HandlerFunc(fn)
}

// Update the value for a key
// If the key exists, update the value
func updateValue(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var val PutRequestBody		
		if err := c.BindJSON(&val); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var updated map[string]string
		var err error
		if updated, err = kv.UpdateValue(key, val.NewValue); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, updated)
	}

	return gin.HandlerFunc(fn)
}

// Delete a value for a key
// Return a 404 if the key doesn't exist
func deleteValue(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var deletedKey string
		var err error
		if deletedKey, err = kv.DeleteValue(key); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, deletedKey)
	}

	return gin.HandlerFunc(fn)
}