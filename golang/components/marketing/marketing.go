package marketing

import (
    "github.com/gin-gonic/gin"
    "app/components/marketing/dto"
    "app/components/marketing/use_case"
)

type Component struct {}

var requestBody dto.RequestBody

func (comp *Component) Index(ctx *gin.Context) string {
    return usecase.AnnounceSomething(&requestBody)
}

func (comp *Component) AnnounceSomething(ctx *gin.Context) (string, error) {
    if err := ctx.BindJSON(&requestBody); err != nil {
        return "", err
    }

    return usecase.AnnounceSomething(&requestBody), nil
}
