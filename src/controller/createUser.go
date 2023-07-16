package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/crudGo/src/configuration/rest_err"
	"github.com/nathanfabio/crudGo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Incorrect fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}