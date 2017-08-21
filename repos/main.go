package repos

import (
	"github.com/ykhrustalev/highloadcup/collections"
	"github.com/ykhrustalev/highloadcup/models"
	"sync"
)

type Repo struct {
	users              map[int]*models.User
	visits             map[int]*models.Visit
	visitsByUser       map[int]*collections.IntSet
	visitsByLocation   map[int]*collections.IntSet
	locations          map[int]*models.Location
	locationsByCountry map[string]*collections.IntSet
}

func NewRepo() *Repo {
	return &Repo{
		users:              make(map[int]*models.User),
		visits:             make(map[int]*models.Visit),
		visitsByUser:       make(map[int]*collections.IntSet),
		visitsByLocation:   make(map[int]*collections.IntSet),
		locations:          make(map[int]*models.Location),
		locationsByCountry: make(map[string]*collections.IntSet),
	}
}
