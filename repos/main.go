package repos

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
)

type Repo struct {
	users              map[int]*models.User
	visits             map[int]*models.Visit
	visitsByUser       map[int][]*models.Visit
	locations          map[int]*models.Location
	locationsByCountry map[string]*collections.IntSet
}

func NewRepo() *Repo {
	return &Repo{
		users:              make(map[int]*models.User),
		visits:             make(map[int]*models.Visit),
		visitsByUser:       make(map[int][]*models.Visit),
		locations:          make(map[int]*models.Location),
		locationsByCountry: make(map[string]*collections.IntSet),
	}
}
