package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const (
	// environment variables
	envCACertificatePath = `CA_CERTIFICATE` // ca certificate file path
	envCertificatePath   = `CERTIFICATE`    // certificate file path
	envKeyPath           = `KEY`            // key file path
)

// GetCertificates for server TLS from path if env vars are provided
func GetCertificates() (*tls.Config, error) {
	// check out two environment variables
	certificatePath := os.Getenv(envCertificatePath)
	keyPath := os.Getenv(envKeyPath)
	// if they are not blank try load the files
	if certificatePath != "" && keyPath != "" {
		serverCert, err := tls.LoadX509KeyPair(certificatePath, keyPath)
		if err != nil {
			return nil, errors.Wrap(err, "reading certificates")
		}
		config := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS13,
		}
		return config, nil
	}

	// @todo: try get these from vault when we can talk to it

	return nil, nil // nothing to do, run insecure for now
}

// LoadCertificates for client TLS from path if env vars are provided
func LoadCertificates() (*tls.Config, error) {
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
