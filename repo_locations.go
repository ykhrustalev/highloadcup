package highloadcup

type LocationsRepo interface {
	Save(*Location) error
	Get(int) (*Location, error)
	Count() int
}

type LocationsRepoImpl struct {
	items map[int]Location
}

func NewLocationsRepoImpl() *LocationsRepoImpl {
	return &LocationsRepoImpl{
		items: make(map[int]Location),
	}
}

func (r *LocationsRepoImpl) Save(item *Location) error {
	r.items[item.Id] = *item
	return nil
}

func (r *LocationsRepoImpl) Get(id int) (*Location, error) {
	item, ok := r.items[id]
	if ok {
		return &item, nil
	}
	return nil, ErrorNotFound
}

func (r *LocationsRepoImpl) Count() int {
	return len(r.items)
}
