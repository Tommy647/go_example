package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// getCertificatesFromPath if readable
func getCertificatesFromPath() (*tls.Config, error) {
	// check out two environment variables
	certificatePath := os.Getenv(envCertificatePath)
	keyPath := os.Getenv(envKeyPath)
	// if they are not blank try load the files
	if certificatePath != "" && keyPath != "" {
		serverCert, err := tls.LoadX509KeyPair(certificatePath, keyPath)
		if err != nil {
			return nil, errors.Wrap(err, "reading certificates from file")
		}
		return &tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS13,
		}, nil
	}

	return nil, nil // nothing to do, run insecure for now
}

// loadCertificatesFromPath if readable
func loadCertificatesFromPath() (*tls.Config, error) {
	serverCA, err := ioutil.ReadFile(os.Getenv(envCACertificatePath))
	if err != nil {
		return nil, errors.Wrap(err, "reading ca certificate file")
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCA) {
		return nil, errors.Wrap(err, "adding ca certificate to pool")
	}
	return &tls.Config{
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS13,
	}, nil
}
