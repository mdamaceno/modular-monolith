package usecase

import (
    "fmt"

    "app/components/marketing/dto"
)

func AnnounceSomething(body *dto.RequestBody) string {
    fmt.Println(body)
    s := "Marketing AnnounceSomething"
    fmt.Println(s)

    return s
}
