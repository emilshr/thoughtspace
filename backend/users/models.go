package users

import (
	"backend/common"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
	Password string `gorm:"column:password"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})
}

// CRUD - Start

func (user *User) CreateUser() {

	db := common.GetDB()

	db.Create(&user)
}

func (user *User) GetUser(id uint64) int {
	db := common.GetDB()

	result := db.Find(&user, id)

	return int(result.RowsAffected)
}

func (user *User) UpdateUser() int {
	db := common.GetDB()

	result := db.Save(&user)

	return int(result.RowsAffected)
}

func (user *User) DeleteUser() int {
	db := common.GetDB()

	result := db.Delete(&user)

	return int(result.RowsAffected)
}

// CRUD - End

func FindUserByUsernameOrEmail(condition interface{}) (User, *gorm.DB) {
	db := common.GetDB()

	var user User

	result := db.Where(condition).Find(&user)

	return user, result
}

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}
