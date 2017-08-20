package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id        int       // уникальный внешний идентификатор пользователя. Устанавливается тестирующей системой и используется затем, для проверки ответов сервера. 32-разрядное целое число.
	Email     string    // адрес электронной почты пользователя. Тип - unicode-строка длиной до 100 символов. Гарантируется уникальность.
	FirstName string    // имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	LastName  string    //
	Gender    string    // unicode-строка "m" означает мужской пол, а "f" - женский.
	BirthDate time.Time // дата рождения, записанная как число секунд от начала UNIX-эпохи по UTC (другими словами - это timestamp). Ограничено снизу 01.01.1930 и сверху 01.01.1999-ым.
}

type userRaw struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func (u *User) SetBirthDate(value int64) {
	u.BirthDate = time.Unix(value, 0)
}

var lowestBirthDate = parseDate("1930-01-01")
var highestBirthDate = parseDate("1999-01-01")

func parseDate(value string) time.Time {
	res, err := time.Parse("2006-01-02", value)
	if err != nil {
		panic(err)
	}
	return res
}

func (u *User) Validate() error {
	if u.Id == 0 {
		return ErrorInvalidId
	}
	if len(u.FirstName) > 50 {
		return ErrorStringTooLong
	}
	if len(u.LastName) > 50 {
		return ErrorStringTooLong
	}
	if len(u.Email) > 100 {
		return ErrorStringTooLong
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

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&userRaw{
		u.Id,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.BirthDate.Unix(),
	})
}

func (u *User) UnmarshalJSON(b []byte) error {
	var obj userRaw
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

	return nil
}

// Partial

type UserPartial struct {
	Email     *string
	FirstName *string
	LastName  *string
	Gender    *string
	BirthDate *int64
}

func (u *UserPartial) UnmarshalJSON(b []byte) error {

	obj := map[string]interface{}{}

	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	value, ok := obj["email"]
	if ok {
		u.Email, err = GetNonNullStringP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["first_name"]
	if ok {
		u.FirstName, err = GetNonNullStringP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["last_name"]
	if ok {
		u.LastName, err = GetNonNullStringP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["gender"]
	if ok {
		u.Gender, err = GetNonNullStringP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["birth_date"]
	if ok {
		u.BirthDate, err = GetNonNullInt64P(value)
		if err != nil {
			return err
		}
	}

	return nil
}
