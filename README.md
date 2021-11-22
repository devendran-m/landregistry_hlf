# Land Registry on Hyperledger Fabric

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)


## About <a name = "about"></a>

Land Registry using Hyperledger Fabric

## Getting Started <a name = "getting_started"></a>

### Prerequisites

[Hyperledger-Fabric](https://github.com/hyperledger/fabric)
[Golang](https://golang.org/)
[Node.js](https://nodejs.org/en/)
[Docker](https://www.docker.com/)

### Creating a Test-Network, Channel and Installing the chaincode

A step by step series of instructions that tell you how to run this in a development env.

Step 1: Clone the repo to a folder in your machine
Step 2: Follow the below steps to ensure the required Chaincode dependencies are available

```bash
cd chaincode/
```

```bash
go mod init landregistry-application-chaincode 
```

```bash
go get -u github.com/hyperledger/fabric-contract-api-go
```

```bash
go mod tidy
```

```bash
go mod vendor
```

Step 3: Build the Chaincode - i.e. compile the Chaincode(Smart Contract) and ensure a successful compilation

```bash
go build
```
Step 4: Deploy the chaincode;
1. Change into test-network directory
   
```bash
cd landregistry/test-network
```

2. Bring the Network Down

```bash
./network.sh down
```

3. Bring the network Up
   
```bash
./network.sh up
```

4. Create channel using createChannel
   
```bash
./network.sh createChannel
```

5. Check docker containers
   
```bash
docker ps -a
```

6. Deploy Chaincode
   
```bash
./network.sh deployCC -ccn landreg -ccl go -ccp ../chaincode -cci InitLedger
```

7. Prepare to use command line arguments
   
```bash
export PATH=${PWD}/../bin:$PATH
```
	
```bash
export FABRIC_CFG_PATH=$PWD/../config/
```

8. Set the Endorsing peers information
   
```bash
source ./scripts/setPeerConnectionParam.sh 1 2
```

9. Set the context for Org1
    
```bash
source ./scripts/setOrgPeerContext.sh 1
```

### Creating a Test-Network, Channel and Installing the chaincode

10. Quert the Property - using QueryProp function - one example below;


```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n landreg $PEER_CONN_PARAMS -c '{"function":"QueryProp","Args":["PROP1"]}'
```

### Output will look like below;

```JPEG

2021-11-22 14:07:46.156 IST [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 001 Chaincode invoke successful. result: status:200 payload:"{\"proptype\":\"Ind House\",\"propcity\":\"Bengaluru\",\"propstate\":\"KA\",\"propsqft\":\"3200\",\"propowner\":\"Abraham\"}" 

```

