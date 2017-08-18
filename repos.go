package highloadcup

type UserRepo interface {
	Add(*User) error
	Get(int) (*User, error)
}

type UserRepoImpl struct {
	users map[int]*User
}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{
		users: make(map[int]*User),
	}
}

func (r *UserRepoImpl) Add(u *User) error {
	r.users[u.Id] = u
	return nil
}

func (r *UserRepoImpl) Get(id int) (*User, error) {
	u, ok := r.users[id]
	if ok {
		return u, nil
	}
	return nil, ErrorNotFound
}
