package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func (app *App) RunServer() {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         app.Addr,
		Handler:      app.Routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server listening on %s", app.Addr)
	err := srv.ListenAndServeTLS(app.TLSCert, app.TLSKey)
	log.Fatal(err)
}
