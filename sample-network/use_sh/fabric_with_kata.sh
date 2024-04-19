./network up
source apply_kata.sh

sleep 5
kubectl get pods -A
