package usecase

import (
    "fmt"

    "app/components/order/dto"
)

func CreateOrder(body *dto.RequestBody) string {
    s := "Order CreateOrder"
    fmt.Println(s)

    return s
}
