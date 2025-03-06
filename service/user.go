package service

import "gorm.io/gorm"

type service struct {
	DB *gorm.DB
}

func New(db *gorm.DB) service {
	return service{db}
}

func (s service) AddNewUser() {

}
