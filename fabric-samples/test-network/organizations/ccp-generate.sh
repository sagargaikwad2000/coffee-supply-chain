#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=1
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/producer.example.com/tlsca/tlsca.producer.example.com-cert.pem
CAPEM=organizations/peerOrganizations/producer.example.com/ca/ca.producer.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/producer.example.com/connection-producer.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/producer.example.com/connection-producer.yaml

ORG=2
P0PORT=8051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/processor.example.com/tlsca/tlsca.processor.example.com-cert.pem
CAPEM=organizations/peerOrganizations/processor.example.com/ca/ca.processor.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/processor.example.com/connection-processor.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/processor.example.com/connection-processor.yaml


ORG=3
P0PORT=9051
CAPORT=9054
PEERPEM=organizations/peerOrganizations/checker.example.com/tlsca/tlsca.checker.example.com-cert.pem
CAPEM=organizations/peerOrganizations/checker.example.com/ca/ca.checker.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/checker.example.com/connection-checker.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/checker.example.com/connection-checker.yaml


ORG=4
P0PORT=10051
CAPORT=10054
PEERPEM=organizations/peerOrganizations/exporter.example.com/tlsca/tlsca.exporter.example.com-cert.pem
CAPEM=organizations/peerOrganizations/exporter.example.com/ca/ca.exporter.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/exporter.example.com/connection-exporter.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/exporter.example.com/connection-exporter.yaml


ORG=5
P0PORT=11051
CAPORT=11054
PEERPEM=organizations/peerOrganizations/importer.example.com/tlsca/tlsca.importer.example.com-cert.pem
CAPEM=organizations/peerOrganizations/importer.example.com/ca/ca.importer.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/importer.example.com/connection-importer.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/importer.example.com/connection-importer.yaml
