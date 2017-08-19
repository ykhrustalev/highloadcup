package highloadcup

type VisitsRepo interface {
	Save(*User) error
	Get(int) (*User, error)
}

type VisitsRepoImpl struct {
	items map[int]*Visit
}

func NewVisitsRepoImpl() *VisitsRepoImpl {
	return &VisitsRepoImpl{
		items: make(map[int]*Visit),
	}
}

func (r *VisitsRepoImpl) Save(u *Visit) error {
	r.items[u.Id] = u
	return nil
}

func (r *VisitsRepoImpl) Get(id int) (*Visit, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}
