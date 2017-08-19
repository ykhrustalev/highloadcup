package highloadcup

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Loader struct {
	users     UsersRepo
	locations LocationsRepo
	visits    VisitsRepo
}

func NewLoader(users UsersRepo, locations LocationsRepo, visits VisitsRepo) *Loader {
	return &Loader{
		users:     users,
		locations: locations,
		visits:    visits,
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

	return nil
}

type UsersLoad struct {
	Users []User `json:"users"`
}

type LocationsLoad struct {
	Locations []Location `json:"locations"`
}

type VisitsLoad struct {
	Visits []Visit `json:"visit"`
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
		l.users.Save(&item)
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
		l.locations.Save(&item)
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
		l.visits.Save(&item)
	}

	return nil
}
