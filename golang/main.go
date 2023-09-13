package main

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "app/components/marketing"
  "app/components/order"
  "app/helpers"
)

func main() {
    r := gin.Default()
    marketingRoute := r.Group("/marketing")

    marketingRoute.GET("/announces", func(c *gin.Context) {
        marketingComponent := marketing.Component{}
        c.JSON(http.StatusOK, helpers.SerializeData(marketingComponent.Index(c)))
    })

    marketingRoute.POST("/announces", func(c *gin.Context) {
        marketingComponent := marketing.Component{}
        result, err := marketingComponent.AnnounceSomething(c)

        if err != nil {
            c.JSON(http.StatusBadRequest, helpers.SerializeError(err.Error(), http.StatusBadRequest))
            return
        }

        c.JSON(http.StatusOK, helpers.SerializeData(result))
    })

    r.POST("/orders", func(c *gin.Context) {
        orderComponent := order.Component{}
        result, err := orderComponent.CreateOrder(c)

        if err != nil {
            c.JSON(http.StatusBadRequest, helpers.SerializeError(err.Error(), http.StatusBadRequest))
            return
        }

        c.JSON(http.StatusOK, helpers.SerializeData(result))
    })

    r.Run()
}
