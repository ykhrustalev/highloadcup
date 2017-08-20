package repos

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
)

func (r *Repo) SaveLocation(item *models.Location) error {
	r.locations[item.Id] = item

	countrySet, ok := r.locationsByCountry[item.Country]
	if !ok {
		countrySet = collections.NewIntSet()
		r.locationsByCountry[item.Country] = countrySet
	}

	countrySet.Add(item.Id)

	return nil
}

func (r *Repo) GetLocation(id int) (*models.Location, bool) {
	item, ok := r.locations[id]
	return item, ok
}

func (r *Repo) CountLocations() int {
	return len(r.locations)
}

func (r *Repo) GetLocationIdsForCountry(country string) *collections.IntSet {
	values, ok := r.locationsByCountry[country]
	if !ok {
		return &collections.EmptyIntSet
	}

	return values
}
