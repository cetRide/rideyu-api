package models

import (
	"fmt"
	"time"

	u "github.com/cetRide/rideyu-api/helpers"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type Base struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := u.GenerateUID()
	return scope.SetColumn("ID", uuid)
}

func InitDB(dbUrI string) {
	connection, err := gorm.Open("postgres", dbUrI)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	connection.AutoMigrate(&User{})
	database = connection
}

func GetDB() *gorm.DB {
	return database
}
