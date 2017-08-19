package highloadcup

import (
	"time"
	"encoding/json"
)

type Visit struct {
	Id        int       `json:"id"`         // уникальный внешний id посещения. Устанавливается тестирующей системой. 32-разрядное целое число.
	Location  int       `json:"location"`   // id достопримечательности. 32-разрядное целое число.
	User      int       `json:"user"`       // id путешественника. 32-разрядное целое число.
	VisitedAt time.Time `json:"visited_at"` // дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	Mark      int       `json:"mark"`       // оценка посещения от 0 до 5 включительно. Целое число.
}

type VisitRaw struct {
	Id        int   `json:"id"`
	Location  int   `json:"location"`
	User      int   `json:"user"`
	VisitedAt int64 `json:"visited_at"`
	Mark      int   `json:"mark"`
}

type VisitPartialRaw struct {
	Id        *int   `json:"id"`
	Location  *int   `json:"location"`
	User      *int   `json:"user"`
	VisitedAt *int64 `json:"visited_at"`
	Mark      *int   `json:"mark"`
}

func (u *Visit) SetVisitedAt(value int64) {
	u.VisitedAt = time.Unix(value, 0)
}

func (u *Visit) Validate() error {

	return nil
}

func (u *Visit) UpdatePartial(source *VisitPartialRaw) error {

	if source.Location != nil {
		u.Location = *source.Location
	}
	if source.User != nil {
		u.User = *source.User
	}
	if source.VisitedAt != nil {
		u.SetVisitedAt(*source.VisitedAt)
	}
	if source.Mark != nil {
		u.Mark = *source.Mark
	}

	return nil
}

func (u *Visit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&VisitRaw{
		u.Id,
		u.Location,
		u.User,
		u.VisitedAt.Unix(),
		u.Mark,
	})
}

func (u *Visit) UnmarshalJSON(b []byte) error {
	var obj VisitRaw
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	u.Id = obj.Id
	u.Location = obj.Location
	u.User = obj.User
	u.SetVisitedAt(obj.VisitedAt)
	u.Mark = obj.Mark

	return nil
}
