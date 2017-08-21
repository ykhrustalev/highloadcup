package repos

import (
	"fmt"
	"github.com/ykhrustalev/highloadcup/models"
	"sort"
)

func (r *Repo) UpdateVisit(target *models.Visit, source *models.VisitPartial) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	var oldLocation, newLocation int
	var oldUser, newUser int

	updateLocationStore := source.Location != nil
	if updateLocationStore {
		oldLocation, newLocation := target.Location, *source.Location
		updateLocationStore = oldLocation != newLocation
	}

	updateUserStore := source.User != nil
	if updateUserStore {
		oldUser, newUser := target.User, *source.User
		updateUserStore = oldUser != newUser
	}

	err := target.UpdatePartial(source)
	if err != nil {
		return err
	}
	err = target.Validate()
	if err != nil {
		return err
	}

	if updateLocationStore {
		removeKeyIn(r.visitsByLocation, oldLocation, target.Id)
		addKeyTo(r.visitsByLocation, newLocation, target.Id)
	}

	if updateUserStore {
		removeKeyIn(r.visitsByUser, oldUser, target.Id)
		addKeyTo(r.visitsByUser, newUser, target.Id)
	}

	return nil
}

func (r *Repo) SaveVisit(item *models.Visit) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.visits[item.Id] = item
	addKeyTo(r.visitsByUser, item.User, item.Id)
	addKeyTo(r.visitsByLocation, item.Location, item.Id)

	return nil
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
				fmt.Println("noloc", item)
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
			fmt.Println("noloc2", visit)
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
				fmt.Println("nouser", item)
				return false
			}

			return filter.FromAge.After(user.BirthDate)
		})
	}

	if filter.ToAge != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			user, found := r.getUserNoLock(item.User)
			if !found {
				fmt.Println("nouser2", item)
				return false
			}

			return filter.ToAge.Before(user.BirthDate)
		})
	}

	if filter.Gender != nil {
		visits = filterVisits(visits, func(item *models.Visit) bool {
			user, found := r.getUserNoLock(item.User)
			if !found {
				fmt.Println("nouser3", item)
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
		} else {
			fmt.Println("No visit for id", id)
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
