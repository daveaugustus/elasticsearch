package models

import (
	"elasticsearch/random"
	"time"
)

type User struct {
	FullName string    `json:"full_name"`
	DOB      time.Time `json:"date_of_birth"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

func (u *User) CreateARandomUser() *User {
	u.FullName = random.NewName()
	u.DOB = random.NewTime()
	u.Email = random.NewEmail(u.FullName)
	u.Created = time.Now()
	u.Updated = time.Now()

	return u
}

func CreateRandomUsers(noOfUsers int32) []User {
	var users []User

	for i := int32(0); i < noOfUsers; i++ {
		var user User
		user.FullName = random.NewName()
		user.DOB = random.NewTime()
		user.Email = random.NewEmail(user.FullName)
		user.Created = time.Now()
		user.Updated = time.Now()

		users = append(users, user)
	}

	return users
}
