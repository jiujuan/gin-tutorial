package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

//POST http://localhost:5000/name?ids[a]=123&ids[b]=789&names[first]=jimmy

//response:
//{
//   "ids": {
//       "a": "123",
//       "b": "789",
//    },
//"names": {
//       "first": "jimmy"
//    },
// }
func main() {
    r := gin.Default()
    r.POST("/name", getName)

    _ = r.Run()
}

func getName(c *gin.Context) {
    ids := c.QueryMap("ids")
    names := c.QueryMap("names")

    c.JSON(http.StatusOK, gin.H{
        "ids": ids,
        "names": names,
    })
    c.String(http.StatusOK, "\nids: %v, names; %v", ids, names)
}
