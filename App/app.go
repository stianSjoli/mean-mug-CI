package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "math/rand"
  "strconv"
)

func setupRouter(number string) *gin.Engine {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, number)
  })
  return r
}

func main() {
  r := setupRouter(strconv.Itoa(rand.Intn(10000)))
  r.Run(":8181")
}
