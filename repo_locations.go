package highloadcup

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
)

type LocationsRepo interface {
	Save(*models.Location) error
	Get(int) (*models.Location, error)
	Count() int
	GetLocationIdsForCountry(string) *collections.IntSet
}

type LocationsRepoImpl struct {
	locations          map[int]*models.Location
	locationsByCountry map[string]*collections.IntSet
}

func NewLocationsRepoImpl() *LocationsRepoImpl {
	return &LocationsRepoImpl{
		locations:          make(map[int]*models.Location),
		locationsByCountry: make(map[string]*collections.IntSet),
	}
}

func (r *LocationsRepoImpl) Save(item *models.Location) error {
	r.locations[item.Id] = item

	countrySet, ok := r.locationsByCountry[item.Country]
	if !ok {
		countrySet = collections.NewIntSet()
		r.locationsByCountry[item.Country] = countrySet
	}

	countrySet.Add(item.Id)

	return nil
}

func (r *LocationsRepoImpl) Get(id int) (*models.Location, error) {
	item, ok := r.locations[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *LocationsRepoImpl) Count() int {
	return len(r.locations)
}

func (r *LocationsRepoImpl) GetLocationIdsForCountry(country string) *collections.IntSet {
	values, ok := r.locationsByCountry[country]
	if !ok {
		return &collections.EmptyIntSet
	}

	return values
}
