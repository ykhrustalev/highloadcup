package highloadcup

import (
	"encoding/json"
	"time"
)

type VisitsRepo interface {
	Save(*Visit) error
	Get(int) (*Visit, error)
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

func (u *Visit) SetVisitedAt(value int64) {
	u.VisitedAt = time.Unix(value, 0)
}

func (u *Visit) Validate() error {

	return nil
}

func (u *Visit) UpdatePartial(source *VisitPartialRaw) error {

	if source.Location != nil {
		u.Location = *source.Location
	}
	if source.User != nil {
		u.User = *source.User
	}
	if source.VisitedAt != nil {
		u.SetVisitedAt(*source.VisitedAt)
	}
	if source.Mark != nil {
		u.Mark = *source.Mark
	}

	return nil
}

func (u *Visit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&VisitRaw{
		u.Id,
		u.Location,
		u.User,
		u.VisitedAt.Unix(),
		u.Mark,
	})
}

func (u *Visit) UnmarshalJSON(b []byte) error {
	var obj VisitRaw
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	u.Id = obj.Id
	u.Location = obj.Location
	u.User = obj.User
	u.SetVisitedAt(obj.VisitedAt)
	u.Mark = obj.Mark

	return nil
}
