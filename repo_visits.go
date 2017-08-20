package highloadcup

import (
	"github.com/ykhrustalev/highloadcup/models"
	"sort"
)

type VisitsRepo interface {
	Save(*models.Visit) error
	Get(int) (*models.Visit, error)
	Count() int
	Filter(int, *models.VisitsFilter) []*models.Visit
}

type VisitsRepoImpl struct {
	visits       map[int]*models.Visit
	visitsByUser map[int][]*models.Visit

	locationsRepo LocationsRepo
}

func NewVisitsRepoImpl(locationsRepo LocationsRepo) *VisitsRepoImpl {
	return &VisitsRepoImpl{
		visits:        make(map[int]*models.Visit),
		visitsByUser:  make(map[int][]*models.Visit),
		locationsRepo: locationsRepo,
	}
}

func (r *VisitsRepoImpl) Save(item *models.Visit) error {
	r.visits[item.Id] = item
	r.storeByUser(item)
	return nil
}

func (r *VisitsRepoImpl) storeByUser(item *models.Visit) {
	arr, ok := r.visitsByUser[item.User]
	if !ok {
		arr = make([]*models.Visit, 0)
	}

	arr = append(arr, item)
	r.visitsByUser[item.User] = arr
}

func (r *VisitsRepoImpl) Get(id int) (*models.Visit, error) {
	item, ok := r.visits[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *VisitsRepoImpl) Count() int {
	return len(r.visits)
}

func (r *VisitsRepoImpl) GetLocation(visit *models.Visit) (*models.Location, error) {
	return r.locationsRepo.Get(visit.Location)
}

func (r *VisitsRepoImpl) Filter(userId int, filter *models.VisitsFilter) []*models.Visit {
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
			location, _ := r.GetLocation(item)
			if location == nil {
				return false
			}

			return location.Distance < *filter.ToDistance
		})
	}

	if filter.Country != nil {
		locationsInCountry := r.locationsRepo.GetLocationIdsForCountry(*filter.Country)

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
