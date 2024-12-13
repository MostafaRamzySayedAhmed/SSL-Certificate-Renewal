package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
)

func main() {
	// Initialize the ACME client
	email := "youremail@example.com"
	privKey, err := certcrypto.GeneratePrivateKey(certcrypto.RSA2048)
	if err != nil {
		log.Fatal(err)
	}

	client, err := lego.NewClient(lego.NewConfig(privKey))
	if err != nil {
		log.Fatal(err)
	}

	// Register the client
	err = client.Register(lego.RegisterOptions{Email: email})
	if err != nil {
		log.Fatal(err)
	}

	// Create a certificate for a domain
	cert, err := client.ObtainCertificate([]string{"example.com"})
	if err != nil {
		log.Fatal(err)
	}

	// Save the certificate and private key
	err = os.WriteFile("cert.pem", cert.Certificate, 0600)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("privkey.pem", cert.PrivateKey, 0600)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SSL Certificate renewed and saved successfully!")
}
