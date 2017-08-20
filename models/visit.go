package models

import (
	"encoding/json"
	"net/url"
	"strconv"
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

func (u *Visit) ToVisitForUser() *VisitForUser {
	return &VisitForUser{
		Location:  u.Location,
		VisitedAt: u.VisitedAt,
		Mark:      u.Mark,
	}
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

type VisitsFilter struct {
	FromDate   *time.Time // посещения с visited_at > fromDate
	ToDate     *time.Time // посещения с visited_at < toDate
	Country    *string    // название страны, в которой находятся интересующие достопримечательности
	ToDistance *int       // возвращать только те места, у которых расстояние от города меньше этого параметра
}

func VisitsFilterFromValues(values *url.Values) (*VisitsFilter, error) {

	filter := &VisitsFilter{}

	fromDate := values.Get("fromDate")
	if fromDate != "" {
		fromDateInt, err := strconv.ParseInt(fromDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetFromDate(fromDateInt)
	}

	toDate := values.Get("toDate")
	if toDate != "" {
		toDateInt, err := strconv.ParseInt(toDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetToDate(toDateInt)
	}

	country := values.Get("country")
	if country != "" {
		filter.Country = &country
	}

	toDistance := values.Get("toDistance")
	if toDistance != "" {
		toDistanceInt, err := strconv.Atoi(toDistance)
		if err != nil {
			return nil, err
		}
		filter.ToDistance = &toDistanceInt
	}

	return filter, nil
}

func (o *VisitsFilter) SetFromDate(value int64) {
	tm := time.Unix(value, 0)
	o.FromDate = &tm
}

func (o *VisitsFilter) SetToDate(value int64) {
	tm := time.Unix(value, 0)
	o.ToDate = &tm
}

func (o *VisitsFilter) Validate() error {
	return nil
}

type VisitsByVisitDate []*Visit

func (a VisitsByVisitDate) Len() int           { return len(a) }
func (a VisitsByVisitDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VisitsByVisitDate) Less(i, j int) bool { return a[i].VisitedAt.Before(a[j].VisitedAt) }

type VisitForUser struct {
	Location  int
	VisitedAt time.Time
	Mark      int
}

type VisitForUserRaw struct {
	Location  int   `json:"location"`
	VisitedAt int64 `json:"visited_at"`
	Mark      int   `json:"mark"`
}

func (u *VisitForUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&VisitForUserRaw{
		u.Location,
		u.VisitedAt.Unix(),
		u.Mark,
	})
}
