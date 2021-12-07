package tls

import (
	"crypto/tls"

	"github.com/pkg/errors"

	"github.com/Tommy647/go_example/internal/vault"
)

// getCertificateFromVault if authorised and available
func getCertificateFromVault() (*tls.Config, error) {
	data, err := vault.GetSecrets("secret", "certificates")
	if err != nil {
		return nil, errors.Wrap(err, "getting certificates")
	}
	_ = data

	cert, err := tls.X509KeyPair([]byte(data["service.pem"]), []byte(data["service.key"]))
	if err != nil {
		return nil, errors.Wrap(err, "reading certificates")
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
		MinVersion:   tls.VersionTLS13,
	}, nil
}
