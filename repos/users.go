package repos

import (
	"github.com/ykhrustalev/highloadcup/models"
)

func (r *Repo) UpdateUser(target *models.User,source *models.UserPartial) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	err := target.UpdatePartial(source)
	if err != nil {
		return err
	}
	err = target.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) SaveUser(item *models.User) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.users[item.Id] = item
	return nil
}

func (r *Repo) GetUser(id int) (*models.User, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.getUserNoLock(id)
}

func (r *Repo) getUserNoLock(id int) (*models.User, bool) {
	item, ok := r.users[id]
	return item, ok
}

func (r *Repo) CountUsers() int {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return len(r.users)
}
