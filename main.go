package main

import (
	"context"
	"flag"
	"fmt"
	"html"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/menucha-de/App.Mqtt/mqtt"

	loglib "github.com/menucha-de/logging"
	utils "github.com/menucha-de/utils"
)

var log *loglib.Logger = loglib.GetLogger("mqtt")

func main() {
	var port = flag.Int("p", 8080, "port")
	flag.Parse()

	mqtt.AddRoutes(loglib.LogRoutes) //must be before router initialization

	router := mqtt.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFound)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	errs := make(chan error)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatal("listen: " + err.Error())
			errs <- err
		}
	}()

	log.Info("Server Started")
	select {
	case err := <-errs:
		log.Error("Could not start serving service due to (error: %s)", err)
	case <-done:
		log.Info("Server Stopped")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		loglib.Close()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed:", err.Error())
	}

	log.Info("Server Exited Properly")

}
func notFound(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET") {
		w.WriteHeader(404)
		return
	}
	file := "./www" + html.EscapeString(r.URL.Path)
	if file == "./www/" {
		file = "./www/index.html"
	}
	if utils.FileExists(file) {
		http.ServeFile(w, r, file)
	} else {
		w.WriteHeader(404)
	}

}
