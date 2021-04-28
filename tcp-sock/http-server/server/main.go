package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "ok")
}

func main() {
	http.HandleFunc("/", helloHandler)
	srv := &http.Server{
		//ReadHeaderTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
		//WriteTimeout: 60 * time.Second,
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("failed to listen server", err)
			os.Exit(1)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("err", err)
	}
}
