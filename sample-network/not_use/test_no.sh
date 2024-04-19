# 로컬 레지스트리 실행
LOCAL_REGISTRY_NAME="local-registry"
LOCAL_REGISTRY_PORT=5000
LOCAL_REGISTRY_INTERFACE=127.0.0.1

running="$(docker inspect -f '{{.State.Running}}' "${LOCAL_REGISTRY_NAME}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
    docker run  \
      --detach  \
      --restart always \
      --name    "${LOCAL_REGISTRY_NAME}" \
      --publish "${LOCAL_REGISTRY_INTERFACE}:${LOCAL_REGISTRY_PORT}:5000" \
      registry:2
fi

# 레지스트리를 네트워크에 연결
docker network connect "bridge" "${LOCAL_REGISTRY_NAME}" || true

# 레지스트리 설정을 클러스터에 적용
cat <<EOF | kubectl apply -f -
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${LOCAL_REGISTRY_PORT}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF

