package models

import (
	"time"
)

type VisitsResponse struct {
	Visits []*VisitForUserRaw `json:"visits"`
}

type VisitForUser struct {
	Place     string
	VisitedAt time.Time
	Mark      int
}

func (obj *VisitForUser) Raw() *VisitForUserRaw {
	return &VisitForUserRaw{
		obj.Place,
		obj.VisitedAt.Unix(),
		obj.Mark,
	}
}

type VisitForUserRaw struct {
	Place     string `json:"place"`
	VisitedAt int64  `json:"visited_at"`
	Mark      int    `json:"mark"`
}
