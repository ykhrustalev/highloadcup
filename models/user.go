package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`         // уникальный внешний идентификатор пользователя. Устанавливается тестирующей системой и используется затем, для проверки ответов сервера. 32-разрядное целое число.
	Email     string    `json:"email"`      // адрес электронной почты пользователя. Тип - unicode-строка длиной до 100 символов. Гарантируется уникальность.
	FirstName string    `json:"first_name"` // имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	LastName  string    `json:"last_name"`  //
	Gender    string    `json:"gender"`     // unicode-строка "m" означает мужской пол, а "f" - женский.
	BirthDate time.Time `json:"birth_date"` // дата рождения, записанная как число секунд от начала UNIX-эпохи по UTC (другими словами - это timestamp). Ограничено снизу 01.01.1930 и сверху 01.01.1999-ым.
}

type UserRaw struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func (obj *UserRaw) User() *User {
	u := &User{
		Id:        obj.Id,
		Email:     obj.Email,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
		Gender:    obj.Gender,
	}
	u.SetBirthDate(obj.BirthDate)
	return u
}

func (u *User) SetBirthDate(value int64) {
	u.BirthDate = time.Unix(value, 0)
}

func (u *User) UserRaw() *UserRaw {
	return &UserRaw{
		u.Id,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.BirthDate.Unix(),
	}
}

var lowestBirthDate = time.Date(1930, 1, 1, 0, 0, 0, 0, time.UTC)
var highestBirthDate = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)

func (u *User) Validate() error {
	if u.Id == 0 {
		return ErrorInvalidId
	}
	if len(u.FirstName) == 0 || len(u.FirstName) > 50 {
		return ErrorStringOutOfRange
	}
	if len(u.LastName) == 0 || len(u.LastName) > 50 {
		return ErrorStringOutOfRange
	}
	if len(u.Email) == 0 || len(u.Email) > 100 {
		return ErrorStringOutOfRange
	}

	_, err := ValidateGender(u.Gender)
	if err != nil {
		return err
	}

	// TODO: validate email
	// github.com/badoux/checkmail

	if u.BirthDate.Before(lowestBirthDate) {
		return ErrorBirthDayToEarly
	}
	if u.BirthDate.After(highestBirthDate) {
		return ErrorBirthDayToLate
	}

	return nil
}

func (u *User) UpdatePartial(source *UserPartial) error {

	if source.Email != nil {
		u.Email = *source.Email
	}
	if source.FirstName != nil {
		u.FirstName = *source.FirstName
	}
	if source.LastName != nil {
		u.LastName = *source.LastName
	}
	if source.Gender != nil {
		u.Gender = *source.Gender
	}

	if source.BirthDate != nil {
		u.SetBirthDate(*source.BirthDate)
	}

	return nil
}

// Partial

type UserPartial struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *int64  `json:"birth_date"`
}
