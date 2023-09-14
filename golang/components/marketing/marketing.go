package marketing

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "app/components/marketing/dto"
    "app/components/marketing/use_case"
    "app/helpers"
)

var requestBody dto.RequestBody

func Controller(router *gin.RouterGroup) error {
    router.GET("/announces", func(c *gin.Context) {
        result, err := usecase.AnnounceSomething(&requestBody)

        if err != nil {
            c.JSON(http.StatusBadRequest, helpers.SerializeError(err.Error(), http.StatusBadRequest))
            return
        }

        c.JSON(http.StatusOK, helpers.SerializeData(result))
        return
    })

    router.POST("/announces", func(c *gin.Context) {
        result, err := usecase.AnnounceSomething(&requestBody)

        if err != nil {
            c.JSON(http.StatusBadRequest, helpers.SerializeError(err.Error(), http.StatusBadRequest))
            return
        }

        c.JSON(http.StatusOK, helpers.SerializeData(result))
        return
    })

    return nil
}
