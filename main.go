package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"

	"github.com/NZB3/without_fails_counter-back/controller"
	counterlib "github.com/NZB3/without_fails_counter-back/counter"
)

func main() {
	var initValue int
	flag.IntVar(&initValue, "value", 0, "initial value of counter")
	flag.Parse()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	counter := counterlib.New(int64(initValue))

	ctrl := controller.NewController(&counter)

	ticker := time.NewTicker(24 * time.Hour)

	router := http.NewServeMux()
	router.HandleFunc("/", ctrl.GetDaysCount)
	router.HandleFunc("/fail", ctrl.Reset)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("$PORT must be set")
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: c.Handler(router),
	}

	go func() {
		log.Printf("server listening at %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	for {
		select {
		case <-ticker.C:
			counter.Inc()
		case <-ctx.Done():
			log.Printf("exiting...")
			ticker.Stop()
			return
		}
	}
}
