package highloadcup

type UsersRepo interface {
	Save(*User) error
	Get(int) (*User, error)
}

type UsersRepoImpl struct {
	items map[int]*User
}

func NewUsersRepoImpl() *UsersRepoImpl {
	return &UsersRepoImpl{
		items: make(map[int]*User),
	}
}

func (r *UsersRepoImpl) Save(u *User) error {
	r.items[u.Id] = u
	return nil
}

func (r *UsersRepoImpl) Get(id int) (*User, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}
