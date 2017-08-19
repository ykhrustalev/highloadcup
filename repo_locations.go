package highloadcup

type LocationsRepo interface {
	Save(*User) error
	Get(int) (*User, error)
}

type LocationsRepoImpl struct {
	items map[int]*Location
}

func NewLocationsRepoImpl() *LocationsRepoImpl {
	return &LocationsRepoImpl{
		items: make(map[int]*Location),
	}
}

func (r *LocationsRepoImpl) Save(u *Location) error {
	r.items[u.Id] = u
	return nil
}

func (r *LocationsRepoImpl) Get(id int) (*Location, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}
