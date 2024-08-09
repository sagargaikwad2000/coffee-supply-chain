package configuration

import (
	"os"

	"github.com/Ashu23042000/coffee-supply-chain/backend/models"
)

func SetNetworkConfig() models.Configuration {
	return models.Configuration{
		MspID:        os.Getenv("mspId"),
		CertPath:     os.Getenv("certPath"),
		KeyPath:      os.Getenv("keyPath"),
		TlsCertPath:  os.Getenv("tlsCertPath"),
		PeerEndpoint: os.Getenv("peerEndpoint"),
		GatewayPeer:  os.Getenv("gatewayPeer"),
	}
}
