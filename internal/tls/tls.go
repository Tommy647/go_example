package tls

import (
	"crypto/tls"
	"log"
)

const (
	// environment variables
	envCACertificatePath = `CA_CERTIFICATE` // ca certificate file path
	envCertificatePath   = `CERTIFICATE`    // certificate file path
	envKeyPath           = `KEY`            // key file path
)

// GetCertificates for server TLS from path if env vars are provided
func GetCertificates() (*tls.Config, error) {
	// @todo: try get these from vault when we can talk to it
	config, err := getCertificateFromVault()
	if err != nil {
		log.Println("vault:", err.Error())

		return getCertificatesFromPath()
	}
	return config, nil
}

// LoadCertificates for client TLS from path if env vars are provided
func LoadCertificates() (*tls.Config, error) {
	// @todo: try get these from vault when we can talk to it

	return loadCertificatesFromPath()
}
