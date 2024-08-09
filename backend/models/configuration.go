package models

type Configuration struct {
	MspID        string
	CertPath     string
	KeyPath      string
	TlsCertPath  string
	PeerEndpoint string
	GatewayPeer  string
}
