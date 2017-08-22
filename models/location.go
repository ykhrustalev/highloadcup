package models

type Location struct {
	Id       int    `json:"id"`       // уникальный внешний id достопримечательности. Устанавливается тестирующей системой. 32-разрядное целое число.
	Place    string `json:"place"`    // описание достопримечательности. Текстовое поле неограниченной длины.
	Country  string `json:"country"`  // название страны расположения. unicode-строка длиной до 50 символов.
	City     string `json:"city"`     // название города расположения. unicode-строка длиной до 50 символов.
	Distance int    `json:"distance"` // расстояние от города по прямой в километрах. 32-разрядное целое число.
}

func (u *Location) Validate() error {
	if u.Id == 0 {
		return ErrorInvalidId
	}
	if len(u.Place) == 0 {
		return ErrorStringOutOfRange
	}
	if len(u.Country) == 0 || len(u.Country) > 50 {
		return ErrorStringOutOfRange
	}
	if len(u.City) == 0 || len(u.City) > 50 {
		return ErrorStringOutOfRange
	}
	return nil
}

func (u *Location) UpdatePartial(source *LocationPartial) error {

	if source.Place != nil {
		u.Place = *source.Place
	}
	if source.Country != nil {
		u.Country = *source.Country
	}
	if source.City != nil {
		u.City = *source.City
	}
	if source.Distance != nil {
		u.Distance = *source.Distance
	}

	return nil
}

// Partial

type LocationPartial struct {
	Place    *string `json:"place"`
	Country  *string `json:"country"`
	City     *string `json:"city"`
	Distance *int    `json:"distance"`
}
