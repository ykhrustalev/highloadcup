package highloadcup

import "github.com/ykhrustalev/highloadcup/models"

type UsersRepo interface {
	Save(*models.User) error
	Get(int) (*models.User, error)
	Count() int
}

type UsersRepoImpl struct {
	items map[int]*models.User
}

func NewUsersRepoImpl() *UsersRepoImpl {
	return &UsersRepoImpl{
		items: make(map[int]*models.User),
	}
}

func (r *UsersRepoImpl) Save(item *models.User) error {
	r.items[item.Id] = item
	return nil
}

func (r *UsersRepoImpl) Get(id int) (*models.User, error) {
	item, ok := r.items[id]
	if ok {
		return item, nil
	}
	return nil, ErrorNotFound
}

func (r *UsersRepoImpl) Count() int {
	return len(r.items)
}
