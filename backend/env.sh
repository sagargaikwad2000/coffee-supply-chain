#!/bin/bash
export mspId="ProducerMSP"
export certPath="../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/signcerts/User1@producer.example.com-cert.pem"
export keyPath="../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/keystore/"
export tlsCertPath="../fabric-samples/test-network/organizations/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/ca.crt"
export peerEndpoint="localhost:7051"
export gatewayPeer="peer0.producer.example.com"


