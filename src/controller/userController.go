package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaodrts/Go.Crud.Api/src/configuration/validation"
	"github.com/joaodrts/Go.Crud.Api/src/controller/model/request"
)

func GetAll(c *gin.Context)     {}
func GetByEmail(c *gin.Context) {}

func Create(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)

		return
	}

	fmt.Println(userRequest)
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
