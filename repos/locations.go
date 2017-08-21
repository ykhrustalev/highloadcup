package repos

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
)

func (r *Repo) UpdateLocation(target *models.Location, source *models.LocationPartial) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	oldCountry, newCountry := target.Country, *source.Country

	target.UpdatePartial(source)
	err := target.Validate()
	if err != nil {
		return err
	}

	if oldCountry != newCountry {
		removeKeyIn2(r.locationsByCountry, oldCountry, target.Id)
		addKeyTo2(r.locationsByCountry, newCountry, target.Id)
	}

	return nil
}

func (r *Repo) SaveLocation(item *models.Location) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.locations[item.Id] = item

	addKeyTo2(r.locationsByCountry, item.Country, item.Id)

	return nil
}

func (r *Repo) GetLocation(id int) (*models.Location, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.getLocationNoLock(id)
}

func (r *Repo) getLocationNoLock(id int) (*models.Location, bool) {
	item, ok := r.locations[id]
	return item, ok
}

func (r *Repo) CountLocations() int {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return len(r.locations)
}

func (r *Repo) getLocationIdsForCountry(country string) *collections.IntSet {
	//r.mx.RLock()
	//defer r.mx.RUnlock()

	values, ok := r.locationsByCountry[country]
	if !ok {
		return &collections.EmptyIntSet
	}

	return values.Copy()
}
