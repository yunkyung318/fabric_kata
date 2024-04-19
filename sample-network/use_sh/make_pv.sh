sudo rm -rf /data/org*

sudo mkdir /data/org0
sudo mkdir /data/org1
sudo mkdir /data/org2

sudo mkdir /data/org0/node1
sudo mkdir /data/org0/node2
sudo mkdir /data/org0/node3

sudo mkdir /data/org1/peer1
sudo mkdir /data/org1/peer1/statedb

sudo mkdir /data/org1/peer2
sudo mkdir /data/org1/peer2/statedb

sudo mkdir /data/org2/peer1
sudo mkdir /data/org2/peer1/statedb

sudo mkdir /data/org2/peer2
sudo mkdir /data/org2/peer2/statedb

kubectl apply -f pv.yaml
kubectl apply -f peer_pv.yaml

kubectl get pv
