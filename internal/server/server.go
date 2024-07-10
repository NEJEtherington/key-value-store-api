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
func InitRoutes(kvdb *db.KeyValueDB) *gin.Engine {	
	// kvdb := db.NewKeyValueDB(make(map[string]string, 0)) 

	router := gin.Default()
	router.GET("/", GetKeys(kvdb))
	router.GET("/:key", GetValue(kvdb))
	router.PUT("/:key", UpdateValue(kvdb))
	router.DELETE("/:key", DeleteValue(kvdb))

	return router
}

// GetKeys returns all keys from the database
// Returns an empty list if the db is empty
func GetKeys(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var keys []string
		keys = kv.GetKeys()
	
		c.JSON(http.StatusOK, keys)
	}

	return gin.HandlerFunc(fn)
}


// GetValue returns the value for a key
// Returns a 404 if key does not exist
func GetValue(kv *db.KeyValueDB) gin.HandlerFunc{
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var val string
		var err error

		if val, err = kv.GetValue(key); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		}

		c.String(http.StatusOK, val)
	}

	return gin.HandlerFunc(fn)
}

// UpdateValue updates the value for a key if it exists 
// Returns a 404 if the key does not exist
func UpdateValue(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var val PutRequestBody		
		if bindErr := c.BindJSON(&val); bindErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
			
			return
		}

		var updated map[string]string
		var err error
		if updated, err = kv.UpdateValue(key, val.NewValue); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, updated)
	}

	return gin.HandlerFunc(fn)
}

// DeleteValue deletes a value for a key
// Returns the key deleted or a 404 if the key doesn't exist
func DeleteValue(kv *db.KeyValueDB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		key := c.Param("key")
		var deletedKey string
		var err error
		if deletedKey, err = kv.DeleteValue(key); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, deletedKey)
	}

	return gin.HandlerFunc(fn)
}