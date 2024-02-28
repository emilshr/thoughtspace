package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.POST("/", Create)
	router.GET("/:id", Get)
	router.PUT("/:id", Update)
	router.DELETE("/:id", Delete)
}

func Create(c *gin.Context) {

	var userInput struct {
		Username string
		Email    string
	}

	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	user.Username = userInput.Username
	user.Email = userInput.Email

	user.CreateUser()

	c.JSON(http.StatusCreated, user)

}

func Get(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var user User

	user.GetUser(id)

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		fmt.Println("users/Update Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var userUpdateInput struct {
		Email string
	}

	var user User

	fetchResult := user.GetUser(id)
	if fetchResult == 0 {
		c.JSON(http.StatusOK, "User not found")
		return
	}

	c.Bind(&userUpdateInput)

	user.Email = userUpdateInput.Email

	user.UpdateUser()

	c.JSON(http.StatusOK, user)
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		fmt.Println("users/Delete Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var user User

	user.GetUser(id)

	user.DeleteUser()

	c.JSON(http.StatusOK, "success")
}
