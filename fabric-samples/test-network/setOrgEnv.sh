#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0




# default to using Producer
ORG=${1:-Producer}

# Exit on first error, print all commands.
set -e
set -o pipefail

# Where am I?
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

ORDERER_CA=${DIR}/test-network/organizations/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem
PEER0_PRODUCER_CA=${DIR}/test-network/organizations/peerOrganizations/producer.example.com/tlsca/tlsca.producer.example.com-cert.pem
PEER0_PROCESSOR_CA=${DIR}/test-network/organizations/peerOrganizations/processor.example.com/tlsca/tlsca.processor.example.com-cert.pem
PEER0_ORG3_CA=${DIR}/test-network/organizations/peerOrganizations/org3.example.com/tlsca/tlsca.org3.example.com-cert.pem


if [[ ${ORG,,} == "org1" || ${ORG,,} == "digibank" || ${ORG,,} == "producer" ]]; then

   CORE_PEER_LOCALMSPID=ProducerMSP
   CORE_PEER_MSPCONFIGPATH=${DIR}/test-network/organizations/peerOrganizations/producer.example.com/users/Admin@producer.example.com/msp
   CORE_PEER_ADDRESS=localhost:7051
   CORE_PEER_TLS_ROOTCERT_FILE=${DIR}/test-network/organizations/peerOrganizations/producer.example.com/tlsca/tlsca.producer.example.com-cert.pem

elif [[ ${ORG,,} == "org2" || ${ORG,,} == "magnetocorp" || ${ORG,,} == "processor"]]; then

   CORE_PEER_LOCALMSPID=PocessorMSP
   CORE_PEER_MSPCONFIGPATH=${DIR}/test-network/organizations/peerOrganizations/processor.example.com/users/Admin@processor.example.com/msp
   CORE_PEER_ADDRESS=localhost:8051
   CORE_PEER_TLS_ROOTCERT_FILE=${DIR}/test-network/organizations/peerOrganizations/processor.example.com/tlsca/tlsca.processor.example.com-cert.pem

else
   echo "Unknown \"$ORG\", please choose Producer/Digibank or Processor/Magnetocorp"
   echo "For example to get the environment variables to set upa Processor shell environment run:  ./setOrgEnv.sh Processor"
   echo
   echo "This can be automated to set them as well with:"
   echo
   echo 'export $(./setOrgEnv.sh Processor | xargs)'
   exit 1
fi

# output the variables that need to be set
echo "CORE_PEER_TLS_ENABLED=true"
echo "ORDERER_CA=${ORDERER_CA}"
echo "PEER0_PRODUCER_CA=${PEER0_PRODUCER_CA}"
echo "PEER0_PROCESSOR_CA=${PEER0_PROCESSOR_CA}"
echo "PEER0_ORG3_CA=${PEER0_ORG3_CA}"

echo "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}"
echo "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}"
echo "CORE_PEER_TLS_ROOTCERT_FILE=${CORE_PEER_TLS_ROOTCERT_FILE}"

echo "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}"
