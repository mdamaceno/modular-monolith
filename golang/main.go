package main

import (
  "github.com/gin-gonic/gin"

  "app/components/marketing"
  "app/components/order"
)

func main() {
    r := gin.Default()

    marketing.Controller(r.Group("/marketing"))
    order.Controller(r.Group("/order"))

    r.Run()
}
