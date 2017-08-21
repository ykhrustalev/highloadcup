package repos

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
	"sort"
)

func (r *Repo) SaveVisit(item *models.Visit) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.visits[item.Id] = item
	r.storeVisitByUser(item)
	r.storeVisitByLocation(item)
	return nil
}

func (r *Repo) storeVisitByUser(item *models.Visit) {
	arr, ok := r.visitsByUser[item.User]
	if !ok {
		arr = collections.NewIntSet()
		r.visitsByUser[item.User] = arr
	}

	arr.Add(item.Id)
}

func (r *Repo) storeVisitByLocation(item *models.Visit) {
	arr, ok := r.visitsByLocation[item.Location]
	if !ok {
		arr = collections.NewIntSet()
		r.visitsByLocation[item.Location] = arr
	}

	arr.Add(item.Id)
}

func (r *Repo) GetVisit(id int) (*models.Visit, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.getVisitNoLock(id)
}

func (r *Repo) getVisitNoLock(id int) (*models.Visit, bool) {
	item, ok := r.visits[id]
	return item, ok
}

func (r *Repo) CountVisits() int {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return len(r.visits)
}

var emptyVisitsForUser = make([]*models.VisitForUser, 0)

func (r *Repo) FilterVisitsForUser(userId int, filter *models.VisitsFilter) []*models.VisitForUser {
	r.mx.RLock()
	defer r.mx.RUnlock()

	visitsSet, ok := r.visitsByUser[userId]
	if !ok {
		return emptyVisitsForUser
	}

	visits := r.visitsFromIds(visitsSet.Values())

	if filter.FromDate != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			return filter.FromDate.Before(item.VisitedAt)
		})
	}

	if filter.ToDate != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			return filter.ToDate.After(item.VisitedAt)
		})
	}

	if filter.ToDistance != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			location, found := r.getLocationNoLock(item.Location)
			if !found {
				return false
			}

			return location.Distance < *filter.ToDistance
		})
	}

	if filter.Country != nil {
		locationsInCountry := r.getLocationIdsForCountry(*filter.Country)

		visits = filterVisits(visits, func(item *models.Visit) bool {
			return locationsInCountry.Contains(item.Location)
		})
	}

	sort.Sort(models.VisitsByVisitDate(visits))

	result := make([]*models.VisitForUser, 0)
	for _, visit := range visits {
		location, ok := r.getLocationNoLock(visit.Location)
		if !ok {
			// TODO: should filter them?
			continue
		}

		obj := &models.VisitForUser{
			Place:     location.Place,
			VisitedAt: visit.VisitedAt,
			Mark:      visit.Mark,
		}

		result = append(result, obj)
	}

	return result
}

func (r *Repo) AverageLocationMark(locationId int, filter *models.LocationsAvgFilter) float32 {
	r.mx.RLock()
	defer r.mx.RUnlock()

	visitsSet, ok := r.visitsByLocation[locationId]
	if !ok {
		return 0.0
	}

	visits := r.visitsFromIds(visitsSet.Values())

	if filter.FromDate != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			return filter.FromDate.Before(item.VisitedAt)
		})
	}

	if filter.ToDate != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			return filter.ToDate.After(item.VisitedAt)
		})
	}

	if filter.FromAge != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			user, found := r.getUserNoLock(item.User)
			if !found {
				return false
			}

			return filter.FromAge.After(user.BirthDate)
		})
	}

	if filter.ToAge != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			user, found := r.getUserNoLock(item.User)
			if !found {
				return false
			}

			return filter.ToAge.Before(user.BirthDate)
		})
	}

	if filter.Gender != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			user, found := r.getUserNoLock(item.User)
			if !found {
				return false
			}

			return *filter.Gender == user.Gender
		})
	}

	res := float32(0)
	for _, visit := range visits {
		res += float32(visit.Mark)
	}

	visits_len := len(visits)

	if visits_len == 0 {
		return 0.0
	}

	return res / float32(visits_len)
}

func (r *Repo) visitsFromIds(ids []int) []*models.Visit {
	res := make([]*models.Visit, 0)

	for _, id := range ids {
		visit, ok := r.getVisitNoLock(id)
		if ok {
			res = append(res, visit)
		}
	}

	return res
}

func filterVisits(items []*models.Visit, predicate func(*models.Visit) bool) []*models.Visit {
	filtered := make([]*models.Visit, 0)
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
