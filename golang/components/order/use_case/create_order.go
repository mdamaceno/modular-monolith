package usecase

import (
    "fmt"

    "app/components/order/dto"
)

func CreateOrder(body *dto.RequestBody) (string, error) {
    s := "Order CreateOrder"
    fmt.Println(s)

    return s, nil
}
