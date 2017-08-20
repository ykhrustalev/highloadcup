package highloadcup

import (
	"sort"
)

type VisitsRepo interface {
	Save(*Visit) error
	Get(int) (*Visit, error)
	Count() int
	Filter(int, *VisitsFilter) []*Visit
}

type VisitsRepoImpl struct {
	visits       map[int]*Visit
	visitsByUser map[int][]*Visit

	locationsRepo LocationsRepo
}

func NewVisitsRepoImpl(locationsRepo LocationsRepo) *VisitsRepoImpl {
	return &VisitsRepoImpl{
		visits:        make(map[int]*Visit),
		visitsByUser:  make(map[int][]*Visit),
		locationsRepo: locationsRepo,
	}
}

func (r *VisitsRepoImpl) Save(item *Visit) error {
	r.visits[item.Id] = item
	r.storeByUser(item)
	return nil
}

func (r *VisitsRepoImpl) storeByUser(item *Visit) {
	arr, ok := r.visitsByUser[item.User]
	if !ok {
		arr = make([]*Visit, 0)
	}

	arr = append(arr, item)
	r.visitsByUser[item.User] = arr
}

func (r *VisitsRepoImpl) Get(id int) (*Visit, error) {
	item, ok := r.visits[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *VisitsRepoImpl) Count() int {
	return len(r.visits)
}

func (r *VisitsRepoImpl) GetLocation(visit *Visit) (*Location, error) {
	return r.locationsRepo.Get(visit.Location)
}

func (r *VisitsRepoImpl) Filter(userId int, filter *VisitsFilter) []*Visit {
	result := make([]*Visit, 0)

	arr, ok := r.visitsByUser[userId]
	if !ok {
		return result
	}

	if filter.FromDate != nil {
		arr = filterVisits(arr, func(item *Visit) bool {
			return filter.FromDate.Before(item.VisitedAt)
		})
	}

	if filter.ToDate != nil {
		arr = filterVisits(arr, func(item *Visit) bool {
			return filter.ToDate.After(item.VisitedAt)
		})
	}

	if filter.ToDistance != nil {
		arr = filterVisits(arr, func(item *Visit) bool {
			location, _ := r.GetLocation(item)
			if location == nil {
				return false
			}

			return location.Distance < *filter.ToDistance
		})
	}

	if filter.Country != nil {
		locationsInCountry := r.locationsRepo.GetLocationIdsForCountry(*filter.Country)

		arr = filterVisits(arr, func(item *Visit) bool {
			return locationsInCountry.Contains(item.Location)
		})
	}

	sort.Sort(VisitsByVisitDate(arr))

	return arr
}

func filterVisits(items []*Visit, predicate func(*Visit) bool) []*Visit {
	filtered := make([]*Visit, 0)
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
