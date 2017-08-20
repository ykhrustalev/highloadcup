package models

type Location struct {
	Id       int    `json:"id"`       // уникальный внешний id достопримечательности. Устанавливается тестирующей системой. 32-разрядное целое число.
	Place    string `json:"place"`    // описание достопримечательности. Текстовое поле неограниченной длины.
	Country  string `json:"country"`  // название страны расположения. unicode-строка длиной до 50 символов.
	City     string `json:"city"`     // название города расположения. unicode-строка длиной до 50 символов.
	Distance int    `json:"distance"` // расстояние от города по прямой в километрах. 32-разрядное целое число.
}

type LocationPartialRaw struct {
	Id       *int    `json:"id"`
	Place    *string `json:"place"`
	Country  *string `json:"country"`
	City     *string `json:"city"`
	Distance *int    `json:"distance"`
}

func (u *Location) Validate() error {
	return nil
}

func (u *Location) UpdatePartial(source *LocationPartialRaw) error {

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
