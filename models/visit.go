package models

import (
	"time"
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

func (obj *VisitRaw) Visit() *Visit {
	v := &Visit{
		Id:       obj.Id,
		Location: obj.Location,
		User:     obj.User,
		Mark:     obj.Mark,
	}
	v.SetVisitedAt(obj.VisitedAt)

	return v
}

func (u *Visit) VisitRaw() *VisitRaw {
	return &VisitRaw{
		u.Id,
		u.Location,
		u.User,
		u.VisitedAt.Unix(),
		u.Mark,
	}
}

func (u *Visit) SetVisitedAt(value int64) {
	u.VisitedAt = time.Unix(value, 0)
}

func (u *Visit) Validate() error {
	if u.Id == 0 {
		return ErrorInvalidId
	}
	if u.Location == 0 {
		return ErrorInvalidId
	}
	if u.User == 0 {
		return ErrorInvalidId
	}
	if u.Mark < 0 || u.Mark > 5 {
		return ErrorMark
	}
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

// Partial

type VisitPartial struct {
	Location  *int   `json:"location"`
	User      *int   `json:"user"`
	VisitedAt *int64 `json:"visited_at"`
	Mark      *int   `json:"mark"`
}

// Sorting

type VisitsByVisitDate []*Visit

func (a VisitsByVisitDate) Len() int           { return len(a) }
func (a VisitsByVisitDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VisitsByVisitDate) Less(i, j int) bool { return a[i].VisitedAt.Before(a[j].VisitedAt) }
