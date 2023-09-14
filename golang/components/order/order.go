package order

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "app/components/order/use_case"
    "app/components/order/dto"
    "app/helpers"
)

type Component struct {}

var requestBody dto.RequestBody

func Controller(router *gin.RouterGroup) error {
    router.POST("", func(c *gin.Context) {
        result, err := usecase.CreateOrder(&requestBody)

        if err != nil {
            c.JSON(http.StatusBadRequest, helpers.SerializeError(err.Error(), http.StatusBadRequest))
            return
        }

        c.JSON(http.StatusOK, helpers.SerializeData(result))
        return
    })

    return nil
}
