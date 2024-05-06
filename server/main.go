package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

var ctx = context.Background()

const (
	port         = ":8443"
	responseBody = "Hello, TLS!"
)

func main() {
	certFile := "./certs/server.crt"
	keyFile := "./certs/server.key"
	caFile := "./certs/ca.crt"

	cer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Println(err)
		return
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		//ClientAuth:   tls.RequireAndVerifyClientCert, //uncomment for serverside verification also.
		//ClientCAs:    certPool,  //uncomment for serverside verification also.
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	_ = config
	server.TLSConfig = config

	log.Printf("Listening on %s...", port)
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}
