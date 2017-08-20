package repos

import (
	"github.com/ykhrustalev/highloadcup/models"
)

func (r *Repo) SaveUser(item *models.User) error {
	r.users[item.Id] = item
	return nil
}

func (r *Repo) GetUser(id int) *models.User {
	item, ok := r.users[id]
	if ok {
		return item
	}
	return nil
}

func (r *Repo) CountUsers() int {
	return len(r.users)
}
