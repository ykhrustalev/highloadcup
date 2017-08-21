package data_loader

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/ykhrustalev/highloadcup/models"
	"github.com/ykhrustalev/highloadcup/repos"
	"io"
	"strings"
)

type Loader struct {
	repo *repos.Repo
}

func NewLoader(repo *repos.Repo) *Loader {
	return &Loader{
		repo: repo,
	}
}

func (l *Loader) Load(path string) error {
	zf, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, file := range zf.File {
		err = l.loadFile(file)
		if err != nil {
			return err
		}
	}

	fmt.Printf("users: %d\n", l.repo.CountUsers())
	fmt.Printf("location: %d\n", l.repo.CountLocations())
	fmt.Printf("vists: %d\n", l.repo.CountVisits())

	return nil
}

type UsersLoad struct {
	Users []*models.User `json:"users"`
}

type LocationsLoad struct {
	Locations []*models.Location `json:"locations"`
}

type VisitsLoad struct {
	Visits []*models.Visit `json:"visits"`
}

func (l *Loader) loadFile(file *zip.File) error {
	fc, err := file.Open()
	if err != nil {
		return err
	}
	defer fc.Close()

	if strings.HasPrefix(file.Name, "users") {
		l.loadUsers(fc)
		fmt.Printf("load %s\n", file.Name)
	} else if strings.HasPrefix(file.Name, "locations") {
		l.loadLocations(fc)
		fmt.Printf("load %s\n", file.Name)
	} else if strings.HasPrefix(file.Name, "visits") {
		l.loadVisits(fc)
		fmt.Printf("load %s\n", file.Name)
	}

	return nil
}

func (l *Loader) loadUsers(reader io.Reader) error {
	var obj UsersLoad
	d := json.NewDecoder(reader)
	err := d.Decode(&obj)
	if err != nil {
		return err
	}

	for _, item := range obj.Users {
		l.repo.SaveUser(item)
	}

	return nil
}

func (l *Loader) loadLocations(reader io.Reader) error {
	var obj LocationsLoad
	d := json.NewDecoder(reader)
	err := d.Decode(&obj)
	if err != nil {
		return err
	}

	for _, item := range obj.Locations {
		l.repo.SaveLocation(item)
	}

	return nil
}

func (l *Loader) loadVisits(reader io.Reader) error {
	var obj VisitsLoad
	d := json.NewDecoder(reader)
	err := d.Decode(&obj)
	if err != nil {
		return err
	}

	for _, item := range obj.Visits {
		l.repo.SaveVisit(item)
	}

	return nil
}
