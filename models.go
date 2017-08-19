package highloadcup

import (
	"time"
)

type Location struct {
	Id       int    `json:"id"`       // уникальный внешний id достопримечательности. Устанавливается тестирующей системой. 32-разрядное целое число.
	Place    string `json:"place"`    // описание достопримечательности. Текстовое поле неограниченной длины.
	Country  string `json:"country"`  // название страны расположения. unicode-строка длиной до 50 символов.
	City     string `json:"city"`     // название города расположения. unicode-строка длиной до 50 символов.
	Distance int    `json:"distance"` // расстояние от города по прямой в километрах. 32-разрядное целое число.
}

type Visit struct {
	Id        int       `json:"id"`         // уникальный внешний id посещения. Устанавливается тестирующей системой. 32-разрядное целое число.
	Location  int       `json:"location"`   // id достопримечательности. 32-разрядное целое число.
	User      int       `json:"user"`       // id путешественника. 32-разрядное целое число.
	VisitedAt time.Time `json:"visited_at"` // дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	Mark      int       `json:"mark"`       // оценка посещения от 0 до 5 включительно. Целое число.
}
