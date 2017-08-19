package highloadcup

import (
	"encoding/json"
	"time"
	"errors"
)

type User struct {
	Id        *int       `json:"id"`         // уникальный внешний идентификатор пользователя. Устанавливается тестирующей системой и используется затем, для проверки ответов сервера. 32-разрядное целое число.
	Email     *string    `json:"email"`      // адрес электронной почты пользователя. Тип - unicode-строка длиной до 100 символов. Гарантируется уникальность.
	FirstName *string    `json:"first_name"` // имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	LastName  *string    `json:"last_name"`
	Gender    *string    `json:"gender"`     // unicode-строка "m" означает мужской пол, а "f" - женский.
	BirthDate *time.Time `json:"birth_date"` // дата рождения, записанная как число секунд от начала UNIX-эпохи по UTC (другими словами - это timestamp). Ограничено снизу 01.01.1930 и сверху 01.01.1999-ым.
}

func (u *User) MarshalJSON() ([]byte, error) {
	type aux struct {
		Id        *int    `json:"id"`
		Email     *string `json:"email"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Gender    *string `json:"gender"`
		BirthDate int64  `json:"birth_date"`
	}

	obj := &aux{
		u.Id,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.BirthDate.Unix(),
	}

	return json.Marshal(obj)
}

func (u *User) UnmarshalJSON(b []byte) error {
	type aux struct {
		Id        *int    `json:"id"`
		Email     *string `json:"email"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Gender    *string `json:"gender"`
		BirthDate int64   `json:"birth_date"`
	}

	var obj aux
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	u.Id = obj.Id
	u.Email = obj.Email
	u.FirstName = obj.FirstName
	u.LastName = obj.LastName

	if obj.Gender != nil {
		if *obj.Gender != "m" && *obj.Gender != "f" {
			return errors.New("gender unexpected")
		}
	}
	u.Gender = obj.Gender

	tm := time.Unix(obj.BirthDate, 0)
	u.BirthDate = &tm

	return nil
}

type Location struct {
	Id       int    `json:"id"`       // уникальный внешний id достопримечательности. Устанавливается тестирующей системой. 32-разрядное целое число.
	Place    string `json:"place"`    // описание достопримечательности. Текстовое поле неограниченной длины.
	Country  string `json:"country"`  // название страны расположения. unicode-строка длиной до 50 символов.
	City     string `json:"city"`     // название города расположения. unicode-строка длиной до 50 символов.
	Distance int    `json:"distance"` // расстояние от города по прямой в километрах. 32-разрядное целое число.
}

type Visit struct {
	Id        int       `json:"id"`         // уникальный внешний id посещения. Устанавливается тестирующей системой. 32-разрядное целое число.
	Location  int       `json:"location"`   // id достопримечательности. 32-разрядное целое число.
	User      int       `json:"user"`       // id путешественника. 32-разрядное целое число.
	VisitedAt time.Time `json:"visited_at"` // дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	Mark      int       `json:"mark"`       // оценка посещения от 0 до 5 включительно. Целое число.
}
