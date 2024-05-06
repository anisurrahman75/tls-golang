package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Load client certificate and private key
	//cert, err := tls.LoadX509KeyPair("./certs/client.crt", "./certs/client.key")
	//if err != nil {
	//	panic(err)
	//}

	// Load CA certificate
	caCert, err := ioutil.ReadFile("./certs/ca.crt")
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration with client certificate and CA certificate
	tlsConfig := &tls.Config{
		RootCAs:   caCertPool,
		ClientCAs: caCertPool,
	}

	// Create a HTTPS client with the custom TLS configuration
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make a GET request to the server
	response, err := httpClient.Get("https://backup.local:8443/")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Handle response
	// For example, read response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))

	return
}
