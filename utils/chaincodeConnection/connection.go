package chaincodeConnection

import (
	"crypto/x509"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func newGrpcConnection(tlsCertPath string, peerAddress string) (*grpc.ClientConn, error) {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	cred := credentials.NewClientTLSFromCert(
		certPool,
		"",
	)
	conn, err := grpc.Dial(peerAddress, grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createConnection(tlsCertPath string, peerAddress string, certPath string, mspID string, keystorePath string) (*client.Gateway, error) {
	grpcClient, err := newGrpcConnection(tlsCertPath, peerAddress)
	if err != nil {
		return nil, err
	}
	id, err := newIdentity(certPath, mspID)
	if err != nil {
		return nil, err
	}
	sign, err := newSign(keystorePath)
	if err != nil {
		return nil, err
	}
	gw, err := client.Connect(
		id,
		client.WithClientConnection(grpcClient),
		client.WithSign(sign),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return nil, err
	}
	return gw, nil
}
