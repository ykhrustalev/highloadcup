package highloadcup

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

type VisitPartialRaw struct {
	Id        *int   `json:"id"`
	Location  *int   `json:"location"`
	User      *int   `json:"user"`
	VisitedAt *int64 `json:"visited_at"`
	Mark      *int   `json:"mark"`
}
