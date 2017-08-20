package repos

import (
	"github.com/ykhrustalev/highloadcup/models"
	"sort"
)

func (r *Repo) SaveVisit(item *models.Visit) error {
	r.visits[item.Id] = item
	r.storeVisitByUser(item)
	return nil
}

func (r *Repo) storeVisitByUser(item *models.Visit) {
	arr, ok := r.visitsByUser[item.User]
	if !ok {
		arr = make([]*models.Visit, 0)
	}

	// TODO: use set?
	arr = append(arr, item)
	r.visitsByUser[item.User] = arr
}

func (r *Repo) GetVisit(id int) *models.Visit {
	item, ok := r.visits[id]
	if ok {
		return item
	}
	return nil
}

func (r *Repo) CountVisits() int {
	return len(r.visits)
}

var emptyVisitsForUser = make([]*models.VisitForUser, 0)

func (r *Repo) FilterVisitsForUser(userId int, filter *models.VisitsFilter) []*models.VisitForUser {
	// TODO: mutex
	visits, ok := r.visitsByUser[userId]
	if !ok {
		return emptyVisitsForUser
	}

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
			location := r.GetLocation(item.Location)
			if location == nil {
				return false
			}

			return location.Distance < *filter.ToDistance
		})
	}

	if filter.Country != nil {
		locationsInCountry := r.GetLocationIdsForCountry(*filter.Country)

		visits = filterVisits(visits, func(item *models.Visit) bool {
			return locationsInCountry.Contains(item.Location)
		})
	}

	sort.Sort(models.VisitsByVisitDate(visits))

	result := make([]*models.VisitForUser, 0)
	for _, visit := range visits {
		location, ok := r.locations[visit.Location]
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

func filterVisits(items []*models.Visit, predicate func(*models.Visit) bool) []*models.Visit {
	filtered := make([]*models.Visit, 0)
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
