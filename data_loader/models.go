package data_loader

import "github.com/ykhrustalev/highloadcup/models"

type UsersLoad struct {
	Users []*models.UserRaw `json:"users"`
}

type LocationsLoad struct {
	Locations []*models.Location `json:"locations"`
}

type VisitsLoad struct {
	Visits []*models.VisitRaw `json:"visits"`
}
