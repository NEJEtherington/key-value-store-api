package main

import (
    "fmt"
    "log"
    "runtime"
    
    "github.com/gin-gonic/gin"
    "kvp-api/internal/db"
    "kvp-api/internal/server"
)

func main() {
    ParallelConfig()
    kvdb := db.NewKeyValueDB(make(map[string]string))
    app := server.InitRoutes(kvdb)
    app.SetTrustedProxies(nil)
    if err := app.Run(":8080"); err != nil {
        log.Panic("Failed to start server", err)
    }
    app.Use(JSONMiddleware(), gin.Recovery())

}

// ParallelConfig sets the number of OS threads
func ParallelConfig() {
    cpus := runtime.NumCPU()
    runtime.GOMAXPROCS(cpus)
    fmt.Printf("Running with %d CPUs", cpus)
}

// JSONMiddleware sets the Content-Type header to applicatiojn/json
func JSONMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Next()
    }
}