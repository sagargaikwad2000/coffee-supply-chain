package gatewaycli

// import (
// 	"crypto/x509"
// 	"fmt"
// 	"os"
// 	"path"
// 	"time"

// 	"github.com/hyperledger/fabric-gateway/pkg/client"
// 	"github.com/hyperledger/fabric-gateway/pkg/identity"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"
// )

// var now = time.Now()
// var assetId = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

// type GatewayClient struct {
// 	Contract *client.Contract
// }

// const (
// 	mspId        = "ProducerMSP"
// 	certPath     = "../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/signcerts/User1@producer.example.com-cert.pem"
// 	keyPath      = "../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/keystore/"
// 	tlsCertPath  = "../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/ca.crt"
// 	peerEndpoint = "localhost:7051"
// 	gatewayPeer  = "peer0.producer.example.com"
// )

// // func New(contract *client.Contract) GatewayClient {
// func New() GatewayClient {

// 	// config := configuration.SetNetworkConfig()

// 	// fmt.Printf("%v\n", config)

// 	clientConnection := newGrpcConnection()
// 	defer clientConnection.Close()

// 	id := newIdentity()
// 	sign := newSign()

// 	// Create a Gateway connection for a specific client identity
// 	gw, err := client.Connect(
// 		id,
// 		client.WithSign(sign),
// 		client.WithClientConnection(clientConnection),
// 		// Default timeouts for different gRPC calls
// 		client.WithEvaluateTimeout(5*time.Second),
// 		client.WithEndorseTimeout(15*time.Second),
// 		client.WithSubmitTimeout(5*time.Second),
// 		client.WithCommitStatusTimeout(1*time.Minute),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer gw.Close()

// 	// Override default values for chaincode and channel name as they may differ in testing contexts.
// 	chaincodeName := "basic"
// 	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
// 		chaincodeName = ccname
// 	}

// 	channelName := "mychannel"
// 	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
// 		channelName = cname
// 	}

// 	network := gw.GetNetwork(channelName)
// 	contract := network.GetContract(chaincodeName)

// 	return GatewayClient{
// 		Contract: contract,
// 	}
// }

// // newGrpcConnection creates a gRPC connection to the Gateway server.
// func newGrpcConnection() *grpc.ClientConn {
// 	certificate, err := loadCertificate(tlsCertPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	certPool := x509.NewCertPool()
// 	certPool.AddCert(certificate)
// 	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

// 	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
// 	if err != nil {
// 		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
// 	}

// 	return connection
// }

// // newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
// func newIdentity() *identity.X509Identity {
// 	certificate, err := loadCertificate(certPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	id, err := identity.NewX509Identity(mspId, certificate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return id
// }

// func loadCertificate(filename string) (*x509.Certificate, error) {
// 	certificatePEM, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read certificate file: %w", err)
// 	}
// 	return identity.CertificateFromPEM(certificatePEM)
// }

// // newSign creates a function that generates a digital signature from a message digest using a private key.
// func newSign() identity.Sign {
// 	files, err := os.ReadDir(keyPath)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to read private key directory: %w", err))
// 	}
// 	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))

// 	if err != nil {
// 		panic(fmt.Errorf("failed to read private key file: %w", err))
// 	}

// 	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
// 	if err != nil {
// 		panic(err)
// 	}

// 	sign, err := identity.NewPrivateKeySign(privateKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return sign
// }
