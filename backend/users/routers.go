package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.RouterGroup) {

	router.POST("/auth/register", RegisterUser)

	router.GET("/:id", Get)
	router.PUT("/:id", Update)
	router.DELETE("/:id", Delete)
}

func RegisterUser(c *gin.Context) {
	var userInput struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, userNameRes := FindUserByUsernameOrEmail(&User{Username: userInput.Username})
	if userNameRes.Error != nil {
		if userNameRes.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": userNameRes.Error})
			return
		}
	}

	if userNameRes.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken!"})
		return
	}

	_, emailRes := FindUserByUsernameOrEmail(&User{Email: userInput.Email})
	if emailRes.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": emailRes.Error})
		return
	}

	if emailRes.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered!"})
		return
	}

	var user User
	user.Email = userInput.Email
	user.Username = userInput.Username
	err := user.setPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user.CreateUser()

	c.JSON(http.StatusCreated, gin.H{"user": user})
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
