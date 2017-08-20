package models

import (
	"time"
	"net/url"
	"strconv"
)

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
		values, err := strconv.ParseInt(fromDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetFromDate(values)
	}

	toDate := values.Get("toDate")
	if toDate != "" {
		value, err := strconv.ParseInt(toDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetToDate(value)
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
