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

func (r *Repo) GetVisit(id int) (*models.Visit, error) {
	item, ok := r.visits[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *Repo) CountVisits() int {
	return len(r.visits)
}

func (r *Repo) FilterVisits(userId int, filter *models.VisitsFilter) []*models.Visit {
	result := make([]*models.Visit, 0)

	arr, ok := r.visitsByUser[userId]
	if !ok {
		return result
	}

	if filter.FromDate != nil {
		arr = filterVisits(arr, func(item *models.Visit) bool {
			return filter.FromDate.Before(item.VisitedAt)
		})
	}

	if filter.ToDate != nil {
		arr = filterVisits(arr, func(item *models.Visit) bool {
			return filter.ToDate.After(item.VisitedAt)
		})
	}

	if filter.ToDistance != nil {
		arr = filterVisits(arr, func(item *models.Visit) bool {
			location, _ := r.GetLocation(item.Location)
			if location == nil {
				return false
			}

			return location.Distance < *filter.ToDistance
		})
	}

	if filter.Country != nil {
		locationsInCountry := r.GetLocationIdsForCountry(*filter.Country)

		arr = filterVisits(arr, func(item *models.Visit) bool {
			return locationsInCountry.Contains(item.Location)
		})
	}

	sort.Sort(models.VisitsByVisitDate(arr))

	return arr
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
