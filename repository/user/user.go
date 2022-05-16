package user

import (
	"errors"
	"fmt"
	"todolist/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewRepoUser(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) Login(email string, password string) (entities.User, error) {
	userID := entities.User{}

	if err := u.Db.Where("email = ? AND password = ?", email, password).First(&userID).Error; err != nil {
		log.Warn(err)
		fmt.Println(err.Error())
		return entities.User{}, errors.New("Email or Password Incorrect")
	}

	return userID, nil
}

func (u *UserRepo) InsertUser(newUser entities.User) (entities.User, error) {
	if err := u.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}

	log.Info()
	return newUser, nil
}

func (u *UserRepo) GetUserID(ID int) (entities.User, error) {
	arrUser := []entities.User{}

	if err := u.Db.Where("id = ?", ID).First(&arrUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("Data Not Found")
	}

	log.Info()
	return arrUser[0], nil
}

func (u *UserRepo) UpdateUser(ID int, update entities.User) (entities.User, error) {
	var res entities.User
	if err := u.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("Email or Phone Must Be Unique")
	}

	log.Info()
	return res, nil
}

func (u *UserRepo) DeleteUser(ID int) error {
	var user entities.User
	if err := u.Db.Delete(&user, "id = ?", ID).Error; err != nil {
		log.Warn(err)
		return errors.New("Data Not Found")
	}
	log.Info()
	return nil

}
