package models

import (
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	Id        int       `json:"id"`         // уникальный внешний идентификатор пользователя. Устанавливается тестирующей системой и используется затем, для проверки ответов сервера. 32-разрядное целое число.
	Email     string    `json:"email"`      // адрес электронной почты пользователя. Тип - unicode-строка длиной до 100 символов. Гарантируется уникальность.
	FirstName string    `json:"first_name"` // имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	LastName  string    `json:"last_name"`
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

type UserPartialRaw struct {
	Id        *int    `json:"id"`
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *int64  `json:"birth_date"`
}

func (u *User) SetBirthDate(value int64) {
	u.BirthDate = time.Unix(value, 0)
}

func (u *User) Validate() error {
	if u.Gender != "m" && u.Gender != "f" {
		return errors.New("gender unexpected")
	}
	return nil
}

func (u *User) UpdatePartial(source *UserPartialRaw) error {

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

	// TODO: validate?
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&UserRaw{
		u.Id,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.BirthDate.Unix(),
	})
}

func (u *User) UnmarshalJSON(b []byte) error {
	var obj UserRaw
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	u.Id = obj.Id
	u.Email = obj.Email
	u.FirstName = obj.FirstName
	u.LastName = obj.LastName
	u.Gender = obj.Gender
	u.SetBirthDate(obj.BirthDate)

	// TODO: validate?
	return nil
}
