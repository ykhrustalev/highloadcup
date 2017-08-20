package highloadcup

import (
	"fmt"
	"github.com/ykhrustalev/highloadcup/data_loader"
	"github.com/ykhrustalev/highloadcup/repos"
	"log"
	"net/http"
	"os"
)

func Server() {
	repo := repos.NewRepo()

	usersRepo := NewUsersRepoImpl()
	usersHandler := NewUsersHandler(usersRepo)

	locationsRepo := NewLocationsRepoImpl()
	locationsHandler := NewLocationsHandler(locationsRepo)

	visitsRepo := NewVisitsRepoImpl(locationsRepo)
	visitsHandler := NewVisitsHandler(visitsRepo)

	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "/tmp/data/data.zip"
	}

	loader := data_loader.NewLoader(repo)
	err := loader.Load(path)
	if err != nil {
		log.Fatalf("failed to load data, %v", err)
	}

	router := NewRouter(usersHandler, locationsHandler, visitsHandler)

	http.HandleFunc("/", router.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("listen on %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
