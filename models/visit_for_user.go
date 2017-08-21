package models

import (
	"encoding/json"
	"time"
)

type VisitForUser struct {
	Place     string
	VisitedAt time.Time
	Mark      int
}

type visitForUserRaw struct {
	Place     string `json:"place"`
	VisitedAt int64  `json:"visited_at"`
	Mark      int    `json:"mark"`
}

func (u *VisitForUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&visitForUserRaw{
		u.Place,
		u.VisitedAt.Unix(),
		u.Mark,
	})
}
