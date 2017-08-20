package highloadcup

type VisitsRepo interface {
	Save(*Visit) error
	Get(int) (*Visit, error)
	Count() int
}

type VisitsRepoImpl struct {
	items map[int]*Visit
}

func NewVisitsRepoImpl() *VisitsRepoImpl {
	return &VisitsRepoImpl{
		items: make(map[int]*Visit),
	}
}

func (r *VisitsRepoImpl) Save(item *Visit) error {
	r.items[item.Id] = item
	return nil
}

func (r *VisitsRepoImpl) Get(id int) (*Visit, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *VisitsRepoImpl) Count() int {
	return len(r.items)
}
