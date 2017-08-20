package models

import (
	"encoding/json"
	"time"
)

type Visit struct {
	Id        int       // уникальный внешний id посещения. Устанавливается тестирующей системой. 32-разрядное целое число.
	Location  int       // id достопримечательности. 32-разрядное целое число.
	User      int       // id путешественника. 32-разрядное целое число.
	VisitedAt time.Time // дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	Mark      int       // оценка посещения от 0 до 5 включительно. Целое число.
}

type visitRaw struct {
	Id        int   `json:"id"`
	Location  int   `json:"location"`
	User      int   `json:"user"`
	VisitedAt int64 `json:"visited_at"`
	Mark      int   `json:"mark"`
}

func (u *Visit) SetVisitedAt(value int64) {
	u.VisitedAt = time.Unix(value, 0)
}

func (u *Visit) Validate() error {
	return nil
}

func (u *Visit) UpdatePartial(source *VisitPartial) error {

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
	return json.Marshal(&visitRaw{
		u.Id,
		u.Location,
		u.User,
		u.VisitedAt.Unix(),
		u.Mark,
	})
}

func (u *Visit) UnmarshalJSON(b []byte) error {
	var obj visitRaw
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

// Partial

type VisitPartial struct {
	Location  *int
	User      *int
	VisitedAt *int64
	Mark      *int
}

func (u *VisitPartial) UnmarshalJSON(b []byte) error {

	obj := map[string]interface{}{}

	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	value, ok := obj["location"]
	if ok {
		u.Location, err = GetNonNullIntP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["user"]
	if ok {
		u.User, err = GetNonNullIntP(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["visited_at"]
	if ok {
		u.VisitedAt, err = GetNonNullInt64P(value)
		if err != nil {
			return err
		}
	}

	value, ok = obj["mark"]
	if ok {
		u.Mark, err = GetNonNullIntP(value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Sorting

type VisitsByVisitDate []*Visit

func (a VisitsByVisitDate) Len() int           { return len(a) }
func (a VisitsByVisitDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VisitsByVisitDate) Less(i, j int) bool { return a[i].VisitedAt.Before(a[j].VisitedAt) }
