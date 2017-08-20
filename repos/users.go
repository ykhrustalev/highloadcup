package repos

import (
	"github.com/ykhrustalev/highloadcup/models"
)

func (r *Repo) SaveUser(item *models.User) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.users[item.Id] = item
	return nil
}

func (r *Repo) GetUser(id int) (*models.User, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	item, ok := r.users[id]
	return item, ok
}

func (r *Repo) CountUsers() int {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return len(r.users)
}
