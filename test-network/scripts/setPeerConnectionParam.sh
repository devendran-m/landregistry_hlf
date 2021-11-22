#!/bin/bash

source scripts/envVar.sh

parsePeerConnectionParameters $@

echo ${PEER_CONN_PARMS[@]}

export PEER_CONN_PARAMS=${PEER_CONN_PARMS[@]}