package highloadcup

import (
	"context"
	"flag"
	"fmt"
	"github.com/ykhrustalev/highloadcup/data_loader"
	"github.com/ykhrustalev/highloadcup/handlers"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"github.com/ykhrustalev/highloadcup/repos"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Server() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	t0 := time.Now()

	repo := repos.NewRepo()

	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "/tmp/data/data.zip"
	}

	loader := data_loader.NewLoader(repo)
	err := loader.Load(path)
	if err != nil {
		log.Fatalf("failed to load data, %v", err)
	}

	router := NewRouter(
		crud.NewHandler(repo),
		handlers.NewListVisitsHandler(repo),
		handlers.NewLocationsAvgHandler(repo),
	)

	http.HandleFunc("/", router.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	t1 := time.Now()
	fmt.Printf("booted in %d seconds\n", t1.Unix()-t0.Unix())

	fmt.Printf("listen on %s\n", port)
	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: nil}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, os.Interrupt:
			fmt.Println("exiting")
			server.Shutdown(context.Background())
		}
	}()

	server.ListenAndServe()
}
