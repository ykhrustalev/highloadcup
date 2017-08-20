package highloadcup

type UsersRepo interface {
	Save(*User) error
	Get(int) (*User, error)
	Count() int
}

type UsersRepoImpl struct {
	items map[int]*User
}

func NewUsersRepoImpl() *UsersRepoImpl {
	return &UsersRepoImpl{
		items: make(map[int]*User),
	}
}

func (r *UsersRepoImpl) Save(item *User) error {
	r.items[item.Id] = item
	return nil
}

func (r *UsersRepoImpl) Get(id int) (*User, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *UsersRepoImpl) Count() int {
	return len(r.items)
}
