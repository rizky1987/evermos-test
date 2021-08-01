package entity

import (
	"evermos-test/helper"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	FullName     string        `bson:"fullname" json:"fullname"`
	NIM          string        `bson:"nim" json:"nim"`
	Role         string        `bson:"role" json:"role"`
	Username     string        `bson:"username" json:"username"`
	Password     string        `bson:"password" json:"password"`
	Address      string        `bson:"address" json:"address"`
	PhoneNumber  string        `bson:"phone_number" json:"phone_number"`
	Email        string        `bson:"email" json:"email"`
	CreatedAtUTC time.Time     `bson:"created_at_utc,omitempty" json:"created_at_utc"`
	UpdatedAtUTC *time.Time    `bson:"updated_at_utc,omitempty" json:"updated_at_utc"`
	CreatedAt    time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt    *time.Time    `bson:"updated_at,omitempty" json:"updated_at"`
}

func (u *User) ValidateBeforeCreate() error {

	u.CreatedAt = helper.GetCurrentTimeAsiaJakarta()
	u.CreatedAtUTC = helper.GetCurrentTimeUTC()

	return nil
}

func (u *User) ValidateBeforeUpdate() error {

	currentTimeAsiaJakarta := helper.GetCurrentTimeAsiaJakarta()
	currentTimeUTC := helper.GetCurrentTimeUTC()

	u.UpdatedAt = &currentTimeAsiaJakarta
	u.UpdatedAtUTC = &currentTimeUTC

	return nil
}

// Please make sure all data is a correct data before we save it to DB
func (u *User) ValidatebeforeSaveToDB() error {

	// Add you Validation Here
	if u.FullName == "" {

	}

	return nil
}

// this function used for clear all input data. Like Remove unused white space of string and etc
func (u *User) CleaningUnUsedData() {

	// elementOfStruct := reflect.ValueOf(&u).Elem()
	// typeOfStruct := elementOfStruct.Type()

	// for i := 0; i < elementOfStruct.NumField(); i++ {
	// 	f := elementOfStruct.Field(i)
	// 	u.f

	// }

}
