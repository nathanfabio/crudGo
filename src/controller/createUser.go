package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/crudGo/src/configuration/validation"
	"github.com/nathanfabio/crudGo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error+%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
	}

	fmt.Println(userRequest)
}