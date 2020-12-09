#!/bin/bash

if [ $# -ne 2 ]; then
	echo "Arguments are missing. ex) ./cc_like.sh instantiate 1.0.0"
	exit 1
fi

instruction=$1
version=$2

set -ev

#chaincode install
docker exec cli peer chaincode install -n bloc -v $version -p github.com/bloc
#chaincode instatiate
docker exec cli peer chaincode $instruction -n bloc -v $version -C mychannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member")'
sleep 3
#chaincode invoke user1
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["addStudent","20164045","2"]}'
sleep 3
#chaincode query user1
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["attand","20164045","algorithm"]}'
sleep 3
#chaincode invoke add rating
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["sit","20164045","algorithm","5"]}'
sleep 3
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["download_material","20164045","algorithm"]}'
sleep 3
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["query_answer","20164045","algorithm","Help me","Very Easy"]}'
sleep 3
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["query_answer","20164045","algorithm","Help me1","Very Easy1"]}'
sleep 3
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["save_note","20164045","algorithm","Root is very hard"]}'
sleep 3
docker exec cli peer chaincode invoke -n bloc -C mychannel -c '{"Args":["exit","20164045","algorithm"]}'
echo '-------------------------------------END-------------------------------------'