package models

import (
	"net/url"
	"strconv"
	"time"
)

type LocationsAvgFilter struct {
	FromDate *time.Time // учитывать оценки только с visited_at > fromDate
	ToDate   *time.Time // учитывать оценки только с visited_at < toDate
	FromAge  *time.Time // учитывать только путешественников, у которых возраст (считается от текущего timestamp) больше этого параметра
	ToAge    *time.Time // как предыдущее, но наоборот
	Gender   *string    // учитывать оценки только мужчин или женщин
}

func LocationsAvgFilterFromValues(values *url.Values) (*LocationsAvgFilter, error) {
	filter := &LocationsAvgFilter{}

	fromDate := values.Get("fromDate")
	if fromDate != "" {
		value, err := strconv.ParseInt(fromDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetFromDate(value)
	}

	toDate := values.Get("toDate")
	if toDate != "" {
		value, err := strconv.ParseInt(toDate, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetToDate(value)
	}

	fromAge := values.Get("fromAge")
	if fromAge != "" {
		value, err := strconv.ParseInt(fromAge, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetFromAge(value)
	}

	toAge := values.Get("toAge")
	if toAge != "" {
		value, err := strconv.ParseInt(toAge, 10, 64)
		if err != nil {
			return nil, err
		}
		filter.SetToAge(value)
	}

	gender := values.Get("gender")
	if gender != "" {
		gender, err := ValidateGender(gender)
		if err != nil {
			return nil, err
		}
		filter.Gender = &gender
	}

	return filter, nil
}

func (o *LocationsAvgFilter) SetFromDate(value int64) {
	tm := time.Unix(value, 0)
	o.FromDate = &tm
}

func (o *LocationsAvgFilter) SetToDate(value int64) {
	tm := time.Unix(value, 0)
	o.ToDate = &tm
}

func nowSubYear(years int) time.Time {
	n := time.Now()
	return time.Date(n.Year()-years, n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second(), n.Nanosecond(), n.Location())
}

func (o *LocationsAvgFilter) SetFromAge(value int64) {
	tm := nowSubYear(int(value))
	o.FromAge = &tm
}

func (o *LocationsAvgFilter) SetToAge(value int64) {
	tm := nowSubYear(int(value))
	o.ToAge = &tm
}
