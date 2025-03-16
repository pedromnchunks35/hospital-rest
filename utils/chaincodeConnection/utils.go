package chaincodeConnection

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"os"
	"path"
)

func loadCertificate(path string) (*x509.Certificate, error) {
	certificatePem, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error loading the certificate")
	}
	return identity.CertificateFromPEM(certificatePem)
}

func newIdentity(certPath string, mspID string) (*identity.X509Identity, error) {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		return nil, err
	}
	id, err := identity.NewX509Identity(
		mspID,
		certificate,
	)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func newSign(keystorePath string) (identity.Sign, error) {
	files, err := os.ReadDir(keystorePath)
	if err != nil {
		return nil, err
	}
	privateKeyFile, err := os.ReadFile(path.Join(keystorePath, files[0].Name()))
	if err != nil {
		return nil, err
	}
	privateKey, err := identity.PrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		return nil, err
	}
	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, err
	}
	return sign, nil
}
