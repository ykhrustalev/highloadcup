package models

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type VisitsFilter struct {
	FromDate   *time.Time // посещения с visited_at > fromDate
	ToDate     *time.Time // посещения с visited_at < toDate
	Country    *string    // название страны, в которой находятся интересующие достопримечательности
	ToDistance *int       // возвращать только те места, у которых расстояние от города меньше этого параметра
}

func VisitsFilterFromValues(values *fasthttp.Args) (*VisitsFilter, error) {
	filter := &VisitsFilter{}

	fromDate := string(values.Peek("fromDate"))
	if fromDate != "" {
		values, err := strconv.ParseInt(fromDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetFromDate(values)
	}

	toDate := string(values.Peek("toDate"))
	if toDate != "" {
		value, err := strconv.ParseInt(toDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetToDate(value)
	}

	country := string(values.Peek("country"))
	if country != "" {
		filter.Country = &country
	}

	toDistance := string(values.Peek("toDistance"))
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
