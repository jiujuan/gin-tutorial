package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/name", getName)
    _ = r.Run()
}

// Query string parameters are parsed using the existing underlying request object.
// The request responds to a url matching:  /name?firstname=Jane&lastname=Doe
func getName(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastName := c.Query("lastname") //shortcut for: c.Request.URL.Query().Get("lastname")
    c.String(http.StatusOK, "hello %s %s", firstname, lastName)
}
