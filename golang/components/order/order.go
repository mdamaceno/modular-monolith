package order

import (
    "github.com/gin-gonic/gin"
    "app/components/order/use_case"
    "app/components/order/dto"
)

type Component struct {}

var requestBody dto.RequestBody

func (comp *Component) CreateOrder(ctx *gin.Context) (string, error) {
    if err := ctx.BindJSON(&requestBody); err != nil {
        return "", err
    }

    return usecase.CreateOrder(&requestBody), nil
}
